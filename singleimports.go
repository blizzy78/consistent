package consistent

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const (
	singleImportsBare   = "bare"
	singleImportsParens = "parens"
)

var singleImportsFlagAllowedValues = []string{flagIgnore, singleImportsBare, singleImportsParens}

func checkSingleImports(pass *analysis.Pass, file *ast.File, mode string) {
	if mode == flagIgnore {
		return
	}

	decls := importGenDecls(file)

	if len(decls) != 1 {
		return
	}

	decl := decls[0]

	if len(decl.Specs) != 1 {
		return
	}

	switch parens := decl.Lparen.IsValid() && decl.Rparen.IsValid(); {
	case mode == singleImportsBare && parens:
		reportf(pass, decl.Pos(), "remove parens around single import declaration")

	case mode == singleImportsParens && !parens:
		reportf(pass, decl.Pos(), "add parens around single import declaration")
	}
}

func importGenDecls(file *ast.File) []*ast.GenDecl {
	decls := []*ast.GenDecl{}

	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if decl.Tok != token.IMPORT {
			continue
		}

		decls = append(decls, decl)
	}

	return decls
}
