package gen

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/zhenzou/gpa/util"
)

const (
	GetDb          = "GetDb()"
	ErrorDecl      = "var err error"
	ModelDecl      = `%s%s%s{}`
	ModelSliceDecl = `%s%s[]%s{}`
	CreateTemplate = `err=%s.Table("%s").Create(%s).Error`
	UpdateTemplate = `err=%s.Table("%s").Where(%s).Update(%s).Error`
	DeleteTemplate = `err=%s.Table("%s").Delete(%s).Error`
	FindTemplate   = `err=%s.Table("%s")%s.Find(&%s,%s).Error`
	PageTemplate   = `.Offset(offset).Limit(limit)`
	ReturnErr      = "return err"
	ReturnModel    = "return %s,err"
	StmtTemplate   = "%s\n%s\n%s"
)

type GormTransformer struct {
}

func (g *GormTransformer) TransformCreate(create *CreateFunc) string {
	tb := util.TableName(create.Table)
	expr := fmt.Sprintf(CreateTemplate, GetDb, tb, create.Func.Params[0])
	_, decls := g.transformResult(create.Func.Receiver.Typ, create.Func.Results)
	return fmt.Sprintf(StmtTemplate, strings.Join(decls, "\n"), expr, ReturnErr)
}

func (g *GormTransformer) TransformUpdate(update *UpdateFunc) string {
	tb := util.TableName(update.Table)
	// 第一个参数是更新的值，暂时不支持按字段更新
	where := g.transformPredicates(update.Predicates, update.Func.Params[1:])
	expr := fmt.Sprintf(UpdateTemplate, GetDb, tb, where, update.Func.Params[0])
	_, decls := g.transformResult(update.Func.Receiver.Typ, update.Func.Results)
	return fmt.Sprintf(StmtTemplate, strings.Join(decls, "\n"), expr, ReturnErr)
}

func (g *GormTransformer) writeBuf(buf *bytes.Buffer, str string) *bytes.Buffer {
	buf.WriteString(str + " ")
	return buf
}

func (g *GormTransformer) transformPredicates(predicates []*Predicate, params []*Field) string {
	buf := bytes.NewBufferString(`"`)
	for i, p := range predicates {
		col := util.TableName(p.Field)
		g.writeBuf(buf, strings.ToUpper(p.Logic))
		g.writeBuf(buf, col)
		switch p.OpCode {
		case OpBetween:
			g.writeBuf(buf, "BETWEEN ? AND ?")
		case OpNotNull:
			g.writeBuf(buf, "IS NOT NULL")
		case OpNull:
			g.writeBuf(buf, "IS NULL")
		case OpLessThan:
			g.writeBuf(buf, "<?")
		case OpLessThanEqual:
			g.writeBuf(buf, "<=?")
		case OpGreaterThan:
			g.writeBuf(buf, ">?")
		case OpGreaterThanEqual:
			g.writeBuf(buf, ">= ?")
		case OpLike:
			g.writeBuf(buf, "LIKE ?")
			params[i].Name = `"%"+` + params[i].Name + `+"%"`
		case OpNotLike:
			g.writeBuf(buf, "NOT LIKE ?")
			params[i].Name = `"%"+` + params[i].Name + `+"%"`
		case OpStartWith:
			g.writeBuf(buf, "LIKE ?")
			params[i].Name = params[i].Name + `+"%"`
		case OpEndWith:
			g.writeBuf(buf, "LIKE ?")
			params[i].Name = `"%"+` + params[i].Name
		case OpNotEmpty:
			g.writeBuf(buf, "<> ''")
		case OpEmpty:
			// TODO
		case OpIn:
			g.writeBuf(buf, "IN ?")
		case OpRegex:
			g.writeBuf(buf, "REGEXP ?")
		case OpNotIn:
			g.writeBuf(buf, "NOT IN ?")
		case OpTrue:
			g.writeBuf(buf, "=true")
		case OpFalse:
			g.writeBuf(buf, "=false")
		case OpNot:
			g.writeBuf(buf, "<>?")
		case OpEqual:
			g.writeBuf(buf, "=?")
		}
	}
	buf.WriteString(`",`)

	length := len(params) - 1
	for i, p := range params {
		if i < length {
			buf.WriteString(p.Name + ",")
		} else {
			buf.WriteString(p.Name)
		}
	}
	return buf.String()
}

var (
	// 如果有这两个参数，应该在最后面，而且应该同时出现
	Limit = Field{
		Name: "limit",
		Typ: Type{
			TypeName:  "int",
			IsPointer: false,
		},
		IsPointer: false,
		IsSlice:   false,
	}
	Offset = Field{
		Name: "offset",
		Typ: Type{
			TypeName:  "int",
			IsPointer: false,
		},
		IsPointer: false,
		IsSlice:   false,
	}
	Err = Field{
		Name: "err",
		Typ: Type{
			TypeName:  "error",
			IsPointer: false,
		},
		IsPointer: false,
		IsSlice:   false,
	}
)

func (g *GormTransformer) TransformFind(find *FindFunc) string {
	tb := util.TableName(find.Table)
	end := len(find.Func.Params)
	var page string
	if len(find.Func.Params)-find.ParamCount == 2 {
		if g.checkPage(find.Func.Params) {
			page = PageTemplate
			end = end - 2
		}
	}
	where := g.transformPredicates(find.Predicates, find.Func.Params[:end])
	name, decls := g.transformResult(find.Func.Receiver.Typ, find.Func.Results)
	expr := fmt.Sprintf(FindTemplate, GetDb, tb, page, name, where)
	return fmt.Sprintf(StmtTemplate, strings.Join(decls, "\n"), expr, fmt.Sprintf(ReturnModel, name))
}

// 检查参数中有没有分页相关参数
func (g *GormTransformer) checkPage(fields []*Field) bool {
	length := len(fields) - 2
	if (*fields[length] == Limit && *fields[length+1] == Offset) || (*fields[length] == Offset && *fields[length+1] == Limit) {
		return true
	}
	return false
}

// 转换返回值
// 暂时只支持（model，error）或者（error）这种吧，其他的没啥必要了，这样也简单
// return modelName,声明列表
// TODO 重构
func (g *GormTransformer) transformResult(recv Type, result []*Field) (string, []string) {
	var decls []string
	var modelName string
	if len(result) == 0 {
		decls = append(decls, ErrorDecl)
	} else if len(result) == 1 && result[0].Typ.TypeName == "error" {
		decls = append(decls, g.transformError(result[0]))
	} else if len(result) == 2 {
		field1 := result[0]
		if field1.Typ.TypeName == recv.TypeName {
			var decl string
			modelName, decl = g.transformModel(field1)
			decls = append(decls, decl)
		}
		decls = append(decls, g.transformError(result[1]))
	}
	return modelName, decls
}

func (g *GormTransformer) transformError(field *Field) string {
	if field.Typ.TypeName == "error" {
		if field.Name == "" {
			return ErrorDecl
		}
	}
	return ""
}

func (g *GormTransformer) transformModel(field *Field) (string, string) {
	name := field.Name
	typ := field.Typ.TypeName
	equal := "="
	template := ModelSliceDecl
	if field.IsSlice {
		if field.Name == "" {
			equal = ":="
			name = util.VarName(field.Typ.TypeName, true)
		}
		if field.Typ.IsPointer {
			typ = "*" + typ
		}
	} else {
		template = ModelDecl
		if field.Name == "" {
			equal = ":="
			name = util.VarName(field.Typ.TypeName, false)
		}
		if field.IsPointer {
			typ = "&" + typ
		}
	}
	decl := fmt.Sprintf(template, name, equal, typ)

	return name, decl
}

func (g *GormTransformer) TransformDelete(delete *DeleteFunc) string {
	tb := util.TableName(delete.Table)
	where := g.transformPredicates(delete.Predicates, delete.Func.Params)
	_, decls := g.transformResult(delete.Func.Receiver.Typ, delete.Func.Results)
	expr := fmt.Sprintf(DeleteTemplate, GetDb, tb, where)
	return fmt.Sprintf(StmtTemplate, strings.Join(decls, "\n"), expr, ReturnErr)
}
