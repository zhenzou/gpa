package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"gpa/common"
)

var (
	mode    = parser.ParseComments | parser.AllErrors
	fileSet = token.NewFileSet()
)

type PostHandler func(fp, origin, result string) error

func NewDebugGpa(gen *Generator) *Gpa {
	return NewGpa(gen, DebugHandler)
}

func NewRewriteGpa(gen *Generator) *Gpa {
	return NewGpa(gen, RewriteHandler)
}

func NewGpa(gen *Generator, handler PostHandler) *Gpa {
	return &Gpa{
		gen:     gen,
		handler: handler,
	}
}

func DebugHandler(fp, origin, result string) error {
	fmt.Println(fp + ":")
	fmt.Println(result)
	return nil
}

func RewriteHandler(fp, origin, result string) error {
	fmt.Println(fp + ":")
	fmt.Println(result)
	return nil
}

type Gpa struct {
	filter  string //reg
	root    string
	handler PostHandler
	gen     *Generator
}

func (g *Gpa) Process(fp string) {
	file, err := os.Stat(fp)
	if err != nil {
		panic(err)
	}
	if common.IsGoFile(file) {
		g.processFile(fp)
	} else {
		fmt.Println("dir not support now")
	}
}

func (g *Gpa) processFile(fp string) string {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}
	file, err := parser.ParseFile(fileSet, fp, bytes.NewReader(data), mode)
	if err != nil {
		panic(err)
	}
	decls := file.Decls
	for _, decl := range decls {
		if fd, ok := decl.(*ast.FuncDecl); ok && len(fd.Body.List) == 0 {
			fun := g.extractFunc(data, fd)
			fun.FileName = fp
			s, err := g.gen.Generate(fun)
			if err != nil {
				println(err.Error())
			} else {
				println(s)
			}
			marshal, _ := json.Marshal(fun)
			println(string(marshal))
		}
	}
	return ""
}

func (g *Gpa) extractFunc(data []byte, decl *ast.FuncDecl) *common.Func {
	fn := decl.Name.Name
	params := []*common.Field{}
	results := []*common.Field{}
	for _, f := range decl.Type.Params.List {
		params = append(params, g.extractField(data, f)...)
	}
	for _, f := range decl.Type.Results.List {
		results = append(results, g.extractField(data, f)...)
	}

	f := &common.Func{
		FullName: fn,
		Params:   params,
		Results:  results,
	}

	if decl.Recv != nil {
		f.Receiver = g.extractField(data, decl.Recv.List[0])[0]
	}

	return f
}

func (g *Gpa) extractField(data []byte, field *ast.Field) []*common.Field {

	fields := []*common.Field{}
	typ := strings.TrimSpace(string(data[field.Type.Pos()-1: field.Type.End()-1]))
	isPointer := false
	isSlice := false
	if strings.HasPrefix(typ, "*") {
		isPointer = true
		typ = strings.TrimPrefix(typ, "*")
	}
	//TODO
	if strings.HasPrefix(typ, "[") {
		isSlice = true
		//strings.Trim
	}
	if len(field.Names) > 0 {
		for _, name := range field.Names {
			fields = append(fields, &common.Field{
				Typ:       common.Type{IsPointer: isPointer, Typ: typ},
				IsPointer: isPointer,
				IsSlice:   isSlice,
				Name:      strings.TrimSpace(name.Name),
			})
		}
	} else {
		fields = append(fields, &common.Field{
			Typ:       common.Type{IsPointer: isPointer, Typ: typ},
			IsPointer: isPointer,
			Name:      "",
		})
	}

	return fields
}
