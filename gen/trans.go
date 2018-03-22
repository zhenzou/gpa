package gen

type Transformer interface {
	TransformCreate(create *CreateFunc) string
	TransformUpdate(update *UpdateFunc) string
	TransformFind(find *FindFunc) string
	TransformDelete(delete *DeleteFunc) string
}
