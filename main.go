package main

import (
	"flag"
	"go/parser"
	"go/token"

	"github.com/zhenzou/gpa/common"
	"github.com/zhenzou/gpa/gen"
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
	gpa.Process("/media/Media/Projects/Go/GOPATH/src/github.com/zhenzou/gpa/example/example.go")
}
