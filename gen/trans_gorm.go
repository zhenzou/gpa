package gen

import (
	"bytes"
	"fmt"

	"gpa/common"
)

const (
	GetDb          = "db"
	CreateTemplate = `%s.Table("%s").Create(%s)`
	UpdateTemplate = `%s.Table("%s").Where(%s).Update(%s)`
	DeleteTemplate = `%s.Table("%s").Delete(%s)`
	FindTemplate   = `%s.Table("%s").Find(%s,%s)`
)

type GormTransformer struct {
}

func (g *GormTransformer) TransformCreate(create *common.Create) string {
	tb := common.TableName(create.Table)
	expr := fmt.Sprintf(CreateTemplate, GetDb, tb, create.Func.Params[0])
	return expr
}

func (g *GormTransformer) TransformUpdate(update *common.Update) string {
	tb := common.TableName(update.Table)
	where := g.transformPredicates(update.Predicates, update.Func.Params[1:])
	expr := fmt.Sprintf(UpdateTemplate, GetDb, tb, where, update.Func.Params[0])
	return expr
}

func (g *GormTransformer) writeBuf(buf *bytes.Buffer, str string) *bytes.Buffer {
	buf.WriteString(str + " ")
	return buf
}

func (g *GormTransformer) transformPredicates(predicates []*common.Predicate, params []*common.Field) string {
	buf := bytes.NewBufferString(`"`)
	for i, p := range predicates {
		col := common.TableName(p.Field)
		g.writeBuf(buf, p.Logic)
		g.writeBuf(buf, col)
		switch p.OpCode {
		case common.OpBetween:
			g.writeBuf(buf, "BETWEEN ? AND ?")
		case common.OpNotNull:
			g.writeBuf(buf, "IS NOT NULL")
		case common.OpNull:
			g.writeBuf(buf, "IS NULL")
		case common.OpLessThan:
			g.writeBuf(buf, "<?")
		case common.OpLessThanEqual:
			g.writeBuf(buf, "<=?")
		case common.OpGreaterThan:
			g.writeBuf(buf, ">?")
		case common.OpGreaterThanEqual:
			g.writeBuf(buf, ">= ?")
		case common.OpLike:
			g.writeBuf(buf, "LIKE ?")
			params[i].Name = `"%"+` + params[i].Name + `+"%"`
		case common.OpNotLike:
			g.writeBuf(buf, "NOT LIKE ?")
			params[i].Name = `"%"+` + params[i].Name + `+"%"`
		case common.OpStartWith:
			g.writeBuf(buf, "LIKE ?")
			params[i].Name = params[i].Name + `+"%"`
		case common.OpEndWith:
			g.writeBuf(buf, "LIKE ?")
			params[i].Name = `"%"+` + params[i].Name
		case common.OpNotEmpty:
			g.writeBuf(buf, "<> ''")
		case common.OpEmpty:
			//TODO
		case common.OpIn:
			g.writeBuf(buf, "IN ?")
		case common.OpRegex:
			g.writeBuf(buf, "REGEXP ?")
		case common.OpNotIn:
			g.writeBuf(buf, "NOT IN ?")
		case common.OpTrue:
			g.writeBuf(buf, "=true")
		case common.OpFalse:
			g.writeBuf(buf, "=false")
		case common.OpNot:
			g.writeBuf(buf, "<>?")
		case common.OpEqual:
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

func (g *GormTransformer) TransformFind(find *common.Find) string {
	tb := common.TableName(find.Table)
	where := g.transformPredicates(find.Predicates, find.Func.Params)
	//find.Func.Results[0]
	expr := fmt.Sprintf(FindTemplate, GetDb, tb, find.Func.Results[0].Name, where)
	return expr
}

func (g *GormTransformer) TransformDelete(delete *common.Delete) string {
	tb := common.TableName(delete.Table)
	where := g.transformPredicates(delete.Predicates, delete.Func.Params)
	expr := fmt.Sprintf(DeleteTemplate, GetDb, tb, where)
	return expr
}
