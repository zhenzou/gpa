package main

import (
	"flag"

	"github.com/zhenzou/gpa/gen"
)

var (
	debug bool
)

func init() {
	flag.BoolVar(&debug, "debug", true, "true to show the code,false to replace the file with the generated code")

	flag.Parse()
}

func main() {
	var gpa *gen.Gpa
	if debug {
		gpa = gen.NewDebugGpa(gen.NewGenerator(&gen.GormTransformer{}, &gen.GpaParser{}))
	} else {
		gpa = gen.NewRewriteGpa(gen.NewGenerator(&gen.GormTransformer{}, &gen.GpaParser{}))
	}
	args := flag.Args()
	for _, filename := range args {
		gpa.Process(filename)
	}
}
