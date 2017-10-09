package gen

import "gpa/common"

type Transformer interface {
	TransformCreate(create *common.Create) string
	TransformUpdate(update *common.Update) string
	TransformFind(find *common.Find) string
	TransformDelete(delete *common.Delete) string
}
