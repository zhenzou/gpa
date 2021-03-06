package gen

import (
	"errors"
)

func NewGenerator(trans Transformer, parser Parser) *Generator {
	return &Generator{
		trans:  trans,
		parser: parser,
	}
}

type Generator struct {
	trans  Transformer
	parser Parser
}

func (g *Generator) Generate(f *Func) (code string, err error) {
	if c, err := g.parser.ParseCreate(f); err == nil {
		code = g.trans.TransformCreate(c)
	} else if d, err := g.parser.ParseDelete(f); err == nil {
		code = g.trans.TransformDelete(d)
	} else if u, err := g.parser.ParseUpdate(f); err == nil {
		code = g.trans.TransformUpdate(u)
	} else if f, err := g.parser.ParseFind(f); err == nil {
		code = g.trans.TransformFind(f)
	} else {
		err = errors.New("not a legal gpa func ")
	}
	return
}
