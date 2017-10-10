package main

import (
	"flag"

	"github.com/zhenzou/gpa/common"
	"github.com/zhenzou/gpa/gen"
)

var (
	filename string
	debug    bool
)

func init() {
	flag.StringVar(&filename, "file", "", "path to process,file or dir")
	flag.BoolVar(&debug, "debug", true, "true to display the result,false to replace the file with the generated code")

	flag.Parse()
}

func main() {
	var gpa *gen.Gpa
	if debug {
		gpa = gen.NewDebugGpa(gen.NewGenerator(&gen.GormTransformer{}, &common.GpaParser{}))
	} else {
		gpa = gen.NewRewriteGpa(gen.NewGenerator(&gen.GormTransformer{}, &common.GpaParser{}))
	}
	gpa.Process(filename)
}
