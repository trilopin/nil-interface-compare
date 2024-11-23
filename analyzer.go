package nilinterfacecompare

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const errMessage = "\"%s\" is an interface and must not be compared to nil. Check your interface usage or compare against reflect.Valueof"

var Analyzer = &analysis.Analyzer{
	Name: "nilinterfacecompare",
	Doc:  "Check that there are no comparison between interfaces and nil",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {

		// only interested on if
		ifStmt, ok := node.(*ast.IfStmt)
		if !ok {
			return true
		}

		// focus on the binary expression of the if condition
		binaryExpr, ok := ifStmt.Cond.(*ast.BinaryExpr)
		if !ok {
			return true
		}

		// extract x and y from "x op y"
		xIdent, ok := binaryExpr.X.(*ast.Ident)
		if !ok {
			return true
		}

		yIdent, ok := binaryExpr.Y.(*ast.Ident)
		if !ok {
			return true
		}

		// check "if a == nil"
		_, isXInterface := pass.TypesInfo.Types[xIdent].Type.(*types.Interface)
		if yIdent.Name == "nil" && isXInterface {
			pass.Reportf(node.Pos(), errMessage, xIdent.Name)
		}

		// check "if nil == a"
		_, isYInterface := pass.TypesInfo.Types[yIdent].Type.(*types.Interface)
		if isYInterface && xIdent.Name == "nil" {
			pass.Reportf(node.Pos(), errMessage, yIdent.Name)
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
