package flagparse

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "flagparse",
	Doc:      "Detecting calls of flag.Parse() during init()",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if !hasImport(pass.Pkg, "flag") {
		return nil, nil
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		funcdecl := n.(*ast.FuncDecl)
		if funcdecl.Name.Name != "init" {
			return
		}
		for _, stmt := range funcdecl.Body.List {
			// Don't report if found declared variable name "flag" before calling flag.Parse().
			if declareFlagVarName(stmt) {
				break
			}
			if node := detectFlagParseCall(stmt); node != nil {
				pass.Report(analysis.Diagnostic{
					Pos:     node.Pos(),
					End:     node.End(),
					Message: "Detect calls of flag.Parse() during init()",
				})
			}
		}
	})
	return nil, nil
}

func detectFlagParseCall(stmt ast.Stmt) ast.Node {
	expr, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return nil
	}
	callexpr, ok := expr.X.(*ast.CallExpr)
	if !ok {
		return nil
	}
	sexpr, ok := callexpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil
	}

	// check flag.Parse().
	if sexpr.X.(*ast.Ident).Name == "flag" && sexpr.Sel.Name == "Parse" && len(callexpr.Args) == 0 {
		return sexpr
	}

	return nil
}

func declareFlagVarName(stmt ast.Stmt) bool {
	assignstmt, ok := stmt.(*ast.AssignStmt)
	if !ok {
		return false
	}
	if assignstmt.Lhs[0].(*ast.Ident).Name == "flag" {
		return true
	}
	return false
}

// copied from analysisutil package in https://github.com/golang/tools/blob/master/go/analysis/passes/internal/analysisutil/util.go#L109.
func hasImport(pkg *types.Package, pkgName string) bool {
	for _, imp := range pkg.Imports() {
		if imp.Path() == pkgName {
			return true
		}
	}
	return false
}
