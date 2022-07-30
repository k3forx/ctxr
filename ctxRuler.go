package ctxRuler

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/types"
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	doc = "ctxRuler is lint tool to check whether context.Context is used in appropriate way..."
)

var Analyzer = &analysis.Analyzer{
	Name: "ctxRuler",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
	}

	config := &types.Config{
		Importer: importer.Default(),
	}
	p, err := config.Importer.Import("context")
	if err != nil {
		log.Fatal("context package import error:", err)
	}
	ctxType := p.Scope().Lookup("Context").Type()
	it, ok := ctxType.Underlying().(*types.Interface)
	if !ok {
		log.Fatal("should be found `Context` interface")
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}
		params := f.Type.Params.List
		for _, p := range params {
			// ast.Print(fset, p)
			varIdent := p.Names[0]

			obj := info.ObjectOf(varIdent)
			if obj == nil {
				return
			}

			if types.Implements(obj.Type(), it) && (varIdent.Name != "ctx") {
				fmt.Printf("%v: variable name of context.Context is invalid\n", varIdent.Pos())
			}
		}
	})

	return nil, nil
}
