package ctxRuler

import (
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "ctxRuler is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "ctxRuler",
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

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}

		for i, p := range f.Type.Params.List {
			for j := range p.Names {
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
							positionText(i+1), f.Name,
						)
					}
					if i != 0 {
						pass.Reportf(
							obj.Pos(),
							"%s args of func '%s' is context.Context, and it should be first arg",
							positionText(i+1), f.Name,
						)
					}
				}
			}
		}
	})

	return nil, nil
}

func positionText(index int) string {
	indexStr := strconv.Itoa(index)
	var first string
	if len(indexStr) == 1 {
		first = indexStr
	} else {
		first = strings.Split(indexStr, "")[0]
	}
	switch first {
	case "1":
		return indexStr + "st"
	case "2":
		return indexStr + "nd"
	case "3":
		return indexStr + "rd"
	default:
		return indexStr + "th"
	}
}
