package main

import (
	"flag"
	"go/parser"
	"go/token"
	"gpa/gen"
	"gpa/common"
)

var (
	mode     = parser.ParseComments | parser.AllErrors
	fileSet  = token.NewFileSet()
	filename string
)

func init() {
	flag.StringVar(&filename, "file", "", "filename path to parse")

	flag.Parse()
}

func main() {

	gpa := gen.NewDebugGpa(gen.NewGenerator(&gen.GormTransformer{}, &common.GpaParser{}))
	gpa.Process("/media/Media/Projects/Go/GOPATH/src/gpa/example/example.go")
}
