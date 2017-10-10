package common

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/zhenzou/gpa/log"
)

const (
	FindPrefix   = "find"
	CreatePrefix = "save"
	UpdatePrefix = "update"
	DeletePrefix = "delete"

	By = "By"
)

// 对方法的参数，返回值，接收者的封装
// 如果IsSlice==true ,Type表示的slice元素的类型
type Field struct {
	Name      string
	Typ       Type
	IsPointer bool
	IsSlice   bool
}

type Type struct {
	TypeName  string
	IsPointer bool
}

type Func struct {
	FileName string
	FullName string
	Params   []*Field
	Results  []*Field
	Receiver *Field
}

type Parser interface {
	ParseCreate(fd *Func) (*Create, error)
	ParseDelete(fd *Func) (*Delete, error)
	ParseUpdate(fd *Func) (*Update, error)
	ParseFind(fd *Func) (*Find, error)
}

//TODO 类型检查，参数检查
type GpaParser struct {
}

// 方法前缀忽略大小写
// 忽略大小写，去掉fullName中的prefix，如果没有返回错误
func (g *GpaParser) trimPrefix(fullName, prefix string) (string, error) {
	fn := strings.ToLower(fullName)
	prefix = strings.ToLower(prefix)
	if strings.HasPrefix(fn, prefix) {
		return fullName[len(prefix):], nil
	} else {
		return "", fmt.Errorf("%s does not has prefix %s", fullName, prefix)
	}
}

// NOTE 暂时只支持使用Receiver作为Table的这种格式吧
// 参考：example
func (g *GpaParser) ParseCreate(fd *Func) (create *Create, err error) {
	create = &Create{Func: fd}
	if _, err = g.trimPrefix(fd.FullName, CreatePrefix); err != nil {
		return
	}
	if fd.Receiver == nil {
		err = errors.New(fd.FullName + " must have a receiver")
	}
	return
}

func (g *GpaParser) ParseDelete(fd *Func) (delete *Delete, err error) {
	var fullName string
	if fullName, err = g.trimPrefix(fd.FullName, DeletePrefix); err != nil {
		return
	}
	if fullName == "" {
		log.Warnf("delete without predicate in %s:%s", fd.FileName, fd.FullName)
	} else {
		if strings.HasPrefix(fullName, By) {
			fn := strings.TrimPrefix(fullName, By)
			predicates, paramCount, err := g.extractPredicate(fn)
			if err != nil {
				return nil, err
			}
			if len(fd.Params) < paramCount {
				err = fmt.Errorf("%s:%s require %d param but found %d", fd.FileName, fd.FullName, paramCount, len(fd.Params))
			}
			delete = &Delete{Func: fd, Table: fd.Receiver.Typ.TypeName, Predicates: predicates}
		} else {
			err = errors.New(fullName + " should start with By")
		}
	}
	return
}

func (g *GpaParser) extractPredicate(str string) (predicates []*Predicate, paramCount int, err error) {
	predicates = []*Predicate{}

	var hasField bool
	var field string
	var op string
	var logic string
	for {
		if str == "" {
			predicate := NewPredicate(field, op, logic)
			predicates = append(predicates, predicate)
			paramCount += predicate.ParamCount
			break
		}
		if !hasField {
			var title string
			if title, err = g.extractTitle(str); err != nil {
				return
			}
			field = title
			hasField = true
			str = strings.TrimPrefix(str, title)
		} else {
			var prefix string
			var ok bool
			if prefix, ok = AnyPrefix(str, AllLogic); ok {
				predicate := NewPredicate(field, op, logic)
				predicates = append(predicates, predicate)
				paramCount += predicate.ParamCount
				logic = prefix
				op = ""
				field = ""
				hasField = false
			} else if prefix, ok = AnyPrefix(str, AllOp); ok {
				op = prefix
			} else {
				if prefix, err = g.extractTitle(str); err != nil {
					return
				}
				field += prefix
			}
			str = strings.TrimPrefix(str, prefix)
		}
	}
	return
}

// 获取第一个单词
// example: SomeOne->Some
func (g *GpaParser) extractTitle(str string) (title string, err error) {
	buf := bytes.NewBuffer(nil)
	for i, r := range str {
		if i == 0 {
			if !unicode.IsUpper(r) {
				err = errors.New(str + " must start with a upper char")
				return
			}
			buf.WriteRune(r)
		} else {
			if unicode.IsUpper(r) {
				break
			}
			buf.WriteRune(r)
		}
	}
	title = buf.String()
	return
}

// TODO 重构
func (g *GpaParser) ParseUpdate(fd *Func) (update *Update, err error) {
	var fullName string
	if fullName, err = g.trimPrefix(fd.FullName, UpdatePrefix); err != nil {
		return
	}
	if fullName == "" {
		log.Warnf("update without predicate in %s:%s", fd.FileName, fd.FullName)
	} else {
		if strings.HasPrefix(fullName, By) {
			fn := strings.TrimPrefix(fullName, By)
			predicates, paramCount, err := g.extractPredicate(fn)
			if err != nil {
				return nil, err
			}
			if len(fd.Params) < paramCount+1 {
				err = fmt.Errorf("%s:%s require %d param but found %d", fd.FileName, fd.FullName, paramCount, len(fd.Params))
			}
			update = &Update{Func: fd, Table: fd.Receiver.Typ.TypeName, Predicates: predicates}
		} else {
			err = errors.New(fullName + " should start with By")
		}
	}
	return
}

//TODO SortBy GroupBy
func (g *GpaParser) ParseFind(fd *Func) (find *Find, err error) {
	var fullName string

	if fullName, err = g.trimPrefix(fd.FullName, FindPrefix); err != nil {
		return
	}
	if fullName == "" || fullName == "All" {
		find = &Find{Func: fd, Table: fd.Receiver.Typ.TypeName}
		return
	}
	if strings.HasPrefix(fullName, By) {
		fn := strings.TrimPrefix(fullName, By)
		predicates, paramCount, err := g.extractPredicate(fn)
		if err != nil {
			return nil, err
		}
		if len(fd.Params) < paramCount {
			err = fmt.Errorf("%s:%s require %d param but found %d", fd.FileName, fd.FullName, paramCount, len(fd.Params))
		}
		find = &Find{Func: fd, Table: fd.Receiver.Typ.TypeName, Predicates: predicates}
	} else {
		err = errors.New(fullName + " should start with By")
	}
	return
}
