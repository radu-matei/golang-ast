package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"net/http"
)

func main() {

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "main.go", nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("failed!, %v", err)
	}

	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if _, err := conf.Check("main.go", fs, []*ast.File{f}, info); err != nil {
		log.Fatal(err) // type error
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			for _, p := range x.Type.Params.List {

				tv, ok := info.Types[p.Type]
				if !ok {
					fmt.Printf("nil...\n")
					return false
				}
				fmt.Printf("%v %v \n", p.Names, tv.Type)
			}
		}
		return true
	})
}

func test(anotherint, n int, abc string, req http.Request) {

}
