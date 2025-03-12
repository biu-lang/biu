package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func emitExpr(expr ast.Expr) {
	switch e := expr.(type) {
	case *ast.CallExpr:
		arg0 := e.Args[0]
		fun := e.Fun
		fmt.Printf("# funcall=%T\n", fun)
		switch fn := fun.(type) {
		case *ast.SelectorExpr:
			emitExpr(arg0)
			fmt.Printf("  popq %%rax\n")
			fmt.Printf("  pushq %%rax\n")
			symbol := fmt.Sprintf("%s.%s", fn.X, fn.Sel)
			fmt.Printf("  call %s\n", symbol)
		}
	case *ast.ParenExpr:
		emitExpr(e.X)
	case *ast.BasicLit:
		fmt.Printf("# start %T\n", e)
		val := e.Value
		ival, _ := strconv.Atoi(val)
		fmt.Printf("  movq $%d, %%rax\n", ival)
		fmt.Printf("  pushq %%rax\n")
		fmt.Printf("# end %T\n", e)
	case *ast.BinaryExpr:
		fmt.Printf("# start %T\n", e)
		emitExpr(e.X) // lhs
		emitExpr(e.Y) // rhs
		if e.Op.String() == "+" {
			fmt.Printf("  popq %%rdi # rhs\n")
			fmt.Printf("  popq %%rax # lhs\n")
			fmt.Printf("  addq %%rdi, %%rax\n")
			fmt.Printf("  pushq %%rax\n")
		} else if e.Op.String() == "-" {
			fmt.Printf("  popq %%rdi # rhs\n")
			fmt.Printf("  popq %%rax # lsh\n")
			fmt.Printf("  subq %%rdi, %%rax\n")
			fmt.Printf("  pushq %%rax\n")
		} else if e.Op.String() == "*" {
			fmt.Printf("  popq %%rdi # rhs\n")
			fmt.Printf("  popq %%rax # rhs\n")
			fmt.Printf("  imulq %%rdi, %%rax\n")
			fmt.Printf("  pushq %%rax\n")
		} else {
			panic(fmt.Sprintf("Unexpected binary operator %s", e.Op))
		}
		fmt.Printf("# end %T\n", e)

	default:
		panic(fmt.Sprintf("Unexpected expr type %T", e))
	}
}

func emitFuncDecl(pkgPrefix string, funcDecl *ast.FuncDecl) {
	fmt.Printf(".text\n")
	fmt.Printf("%s.%s:\n", pkgPrefix, funcDecl.Name)

	for _, stmt := range funcDecl.Body.List {
		switch stmt.(type) {
		case *ast.ExprStmt:
			expr := stmt.(*ast.ExprStmt).X
			emitExpr(expr)
		default:
			panic("Unexpected stmt type")
		}
	}

	fmt.Printf("  ret\n")
}

func generateCode(f *ast.File) {
	for _, decl := range f.Decls {
		switch decl.(type) {
		case *ast.GenDecl:
			continue
		case *ast.FuncDecl:
			funcDecl := decl.(*ast.FuncDecl)
			fmt.Printf("# funcDecl %s\n", funcDecl.Name)
			emitFuncDecl("main", funcDecl)
		default:
			panic("unexpected decl type")
		}
	}
}

func main() {
	fset := &token.FileSet{}
	f, err := parser.ParseFile(fset, "./t/source.go", nil, 0)
	if err != nil {
		panic(err)
	}

	generateCode(f)
}
