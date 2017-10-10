package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/zhenzou/gpa/common"
	"github.com/zhenzou/gpa/log"
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
	fmt.Println()
	fmt.Println(fp + ":")
	fmt.Println(result)
	return nil
}

func RewriteHandler(fp, origin, result string) error {
	return ioutil.WriteFile(fp, []byte(result), os.ModePerm)
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
	buf := common.NewOutputBuffer(data)
	var off int
	for _, decl := range file.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok && len(fd.Body.List) == 0 {
			fun := g.extractFunc(data, fd)
			fun.FileName = fp
			s, err := g.gen.Generate(fun)
			if err != nil {
				log.Error(err.Error())
			} else {
				s = "\n" + s
				buf.WriteStringAt(s, off+int(fd.Body.Lbrace))
				off += len(s)
			}
		}
	}
	if err = g.handler(fp, string(data), buf.String()); err != nil {
		log.Error(err.Error())
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

// NOTE 一些错误没有处理
// 假设，方法除了没有方法体以外，没有其他语法错误
func (g *Gpa) extractField(data []byte, field *ast.Field) []*common.Field {

	fields := []*common.Field{}
	typ := strings.TrimSpace(string(data[field.Type.Pos()-1: field.Type.End()-1]))
	isPointer := false
	isSlice := false

	// NOTE 暂时不支持 *[]*Type这种
	// TODO Array，暂时只支持slice吧，反正Array用的少
	if strings.HasPrefix(typ, "[") {
		isSlice = true
		buf := bytes.NewBuffer(nil)
		for _, r := range typ {
			buf.WriteRune(r)
			if r == ']' {
				break
			}
		}
		typ = strings.TrimPrefix(typ, buf.String())
	}

	if strings.HasPrefix(typ, "*") {
		isPointer = true
		typ = strings.TrimSpace(strings.TrimPrefix(typ, "*"))
	}
	if len(field.Names) > 0 {
		for _, name := range field.Names {
			fields = append(fields, &common.Field{
				Typ:       common.Type{IsPointer: isPointer, TypeName: typ},
				IsPointer: isPointer,
				IsSlice:   isSlice,
				Name:      strings.TrimSpace(name.Name),
			})
		}
	} else {
		fields = append(fields, &common.Field{
			Typ:       common.Type{IsPointer: isPointer, TypeName: typ},
			IsPointer: isPointer,
			IsSlice:   isSlice,
		})
	}
	return fields
}
