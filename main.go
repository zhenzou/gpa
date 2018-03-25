package main

import (
	"flag"
	"os"

	"github.com/zhenzou/gpa/gen"
	"github.com/zhenzou/gpa/log"
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
	// 支持go generate
	if fp := os.Getenv("GOFILE"); fp != "" {
		gpa.Process(fp)
	}
	for _, filename := range args {
		log.Debug(filename)
		gpa.Process(filename)
	}
}
