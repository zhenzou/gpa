package gen

import (
	"gpa/common"
	"encoding/json"
)

type GormTransformer struct {
}

func (g *GormTransformer) TransformCreate(create *common.Create) string {
	s, _ := json.Marshal(create)
	str := string(s)
	common.Debug(str)
	return str
}

func (g *GormTransformer) TransformUpdate(update *common.Update) string {
	s, _ := json.Marshal(update)
	str := string(s)
	common.Debug(str)
	return str
}

func (g *GormTransformer) TransformFind(find *common.Find) string {
	s, _ := json.Marshal(find)
	str := string(s)
	common.Debug(str)
	return str
}

func (g *GormTransformer) TransformDelete(delete *common.Delete) string {
	s, _ := json.Marshal(delete)
	str := string(s)
	common.Debug(str)
	return str
}
