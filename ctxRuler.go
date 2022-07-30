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

	config := &types.Config{
		Importer: importer.Default(),
	}
	p, err := config.Importer.Import("context")
	if err != nil {
		log.Fatal("context package import error:", err)
	}
	it, ok := p.Scope().Lookup("Context").Type().Underlying().(*types.Interface)
	if !ok {
		log.Fatal("should be found `Context` interface")
	}
	fmt.Printf("it: %+v\n", it)
	fmt.Printf("pass: %+v\n", pass.TypesInfo.Defs)
	fmt.Println("-------------------------------------------")

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}

		for i, p := range f.Type.Params.List {
			varIdent := p.Names[0]

			obj := pass.TypesInfo.ObjectOf(varIdent)
			fmt.Printf("i: %+v, obj: %+v\n", i, obj)
			if obj == nil {
				continue
			}
			i, ok := obj.Type().Underlying().(*types.Interface)
			if !ok {
				continue
			}
			fmt.Println(i, it)
			fmt.Println(i == it)
			fmt.Println(types.Implements(obj.Type().Underlying(), it))

			// tp := pass.TypesInfo.TypeOf(p.Type)
			// fmt.Printf("f: %+v, i: %+v, p: %+v, var: %+v\n", f.Name, i, p.Type, p.Names[0].Name)
			// fmt.Printf("tp: %+v, underlying: %+v\n", tp, tp.Underlying())

			// t, ok := tp.Underlying().(*types.Interface)
			// if !ok {
			// 	continue
			// }
			// fmt.Printf("t : %+v\n", t)
			// fmt.Printf("it: %+v\n", it)

			// if types.Implements(pass.TypesInfo.TypeOf(p.Type).Underlying(), it) {
			// 	// fmt.Printf("i: %d, tp name: %s\n", i, p.Names[0].Name)
			// 	pass.Reportf(f.Pos(), "variable name of `context.Context` is invalid")
			// }
			fmt.Println("--------------------")
		}
	})

	return nil, nil
}
