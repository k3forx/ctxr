package ctxr

import (
	"go/ast"
	"go/types"
	"strconv"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "ctxr is lint to check whether context.Context used in a func is passed as first arg with name of `ctx`"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "ctxr",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	ctxObj := analysisutil.LookupFromImports(pass.Pkg.Imports(), "context", "Context")
	if ctxObj == nil {
		return nil, nil
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}

		var index int
		for i, p := range f.Type.Params.List {
			for j := range p.Names {
				index++
				ident := p.Names[j]

				obj := pass.TypesInfo.ObjectOf(ident)
				if obj == nil {
					continue
				}

				if types.Identical(obj.Type(), ctxObj.Type()) {
					if ident.Name != "ctx" {
						pass.Reportf(
							obj.Pos(),
							"%s args of func '%s' is context.Context, and its name should be 'ctx'",
							ordinal(index), f.Name,
						)
					}
					if i != 0 {
						pass.Reportf(
							obj.Pos(),
							"%s args of func '%s' is context.Context, and it should be first arg",
							ordinal(index), f.Name,
						)
					}
				}
			}
		}
	})

	return nil, nil
}

func ordinal(n int) string {
	i := n % 100
	if i == 11 || i == 12 || i == 13 {
		return strconv.Itoa(n) + "th"
	}

	switch n % 10 {
	case 1:
		return strconv.Itoa(n) + "st"
	case 2:
		return strconv.Itoa(n) + "nd"
	case 3:
		return strconv.Itoa(n) + "rd"
	default:
		return strconv.Itoa(n) + "th"
	}
}
