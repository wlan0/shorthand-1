package gen

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"github.com/golang/glog"
	"github.com/kr/pretty"
)

// DeserializeAndPrintFileAST parse the contents of a Go source file and print the AST.
func DeserializeAndPrintFileAST(fileContents string) {
	fset, f := DeserializeFileAST(fileContents)
	PrintAST(fset, f)
}

// DeserializeFileAST parse the contents of a Go source file and return FileSet and AST
func DeserializeFileAST(fileContents string) (*token.FileSet, *ast.File) {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", fileContents, 0)
	if err != nil {
		glog.Fatal(err)
	}

	return fset, f
}

// PrintAST print the AST
func PrintAST(fset *token.FileSet, f *ast.File) {
	ast.Print(fset, f)
	pretty.Println(f)
}

// SerializeFileAST use go/format to write a Go source file to a buffer.
func SerializeFileAST(file *ast.File) *bytes.Buffer {
	// Create a FileSet for node. Since the node does not come
	// from a real source file, fset will be empty.
	fset := token.NewFileSet()

	var buf bytes.Buffer
	err := format.Node(&buf, fset, file)
	if err != nil {
		glog.Fatal(err)
	}

	return &buf
}

// PrintAllTokens print all tokens.
func PrintAllTokens() {
	for i := 0; i <= int(token.VAR); i++ {
		_, _ = pretty.Printf("%d: %s\n", i, token.Token(i).String())
	}
}
