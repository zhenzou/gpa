## intro
gpa is a experimental,spring-data-jpa like code generator for GO.
For now,it is only impl little feature.

<strong>NOTE: it is experimental,and the code generated may be not usable.

<strong>NOTE: it is just a toy,so you should not use it :).

## example
```go
package example

import "github.com/jinzhu/gorm"

// example to guide impl :)
type Model struct {
	Id       string
	Name     string
	LastName string
}

func GetDb() *gorm.DB {
	return &gorm.DB{}
}

// result can come without a name
func (m *Model) findById(id string) (model *Model, err error) {

}

// must no use *[]*Model
func (m *Model) findByName(name string) (models []*Model, err error) {

}

func (m *Model) findByNameAndLastname(name, lastName string) (models []*Model, err error) {

}

func (m *Model) deleteByName(name string) (error) {

}

func (m *Model) deleteByIdAndLastName(id, lastName string) (err error) {

}

func (m *Model) findByLastName(lastName string, limit int, offset int) (models []*Model, err error) {

}

func (m *Model) findByAgeBetween(min, max, limit, offset int) (models []*Model, err error) {

}
```

$ go run main.go -file ./example/example.go

After：

```go
package example

import "github.com/jinzhu/gorm"

// example to guide impl :)
type Model struct {
	Id       string
	Name     string
	LastName string
}

func GetDb() *gorm.DB {
	return &gorm.DB{}
}

// result can come without a name
func (m *Model) findById(id string) (model *Model, err error) {
	model = &Model{}

	err = GetDb().Table("model").Find(&model, " id =? ", id).Error
	return model, err

}

// must no use *[]*Model
func (m *Model) findByName(name string) (models []*Model, err error) {
	models = []*Model{}

	err = GetDb().Table("model").Find(&models, " name =? ", name).Error
	return models, err

}

func (m *Model) findByNameAndLastname(name, lastName string) (models []*Model, err error) {
	models = []*Model{}

	err = GetDb().Table("model").Find(&models, " name =? And lastname =? ", name, lastName).Error
	return models, err

}

func (m *Model) deleteByName(name string) (error) {
	var err error
	err = GetDb().Table("model").Delete(" name =? ", ).Error
	return err

}

func (m *Model) deleteByIdAndLastName(id, lastName string) (err error) {

	err = GetDb().Table("model").Delete(" id =? And last_name =? ", lastName).Error
	return err

}

func (m *Model) findByLastName(lastName string, limit int, offset int) (models []*Model, err error) {
	models = []*Model{}

	err = GetDb().Table("model").Offset(offset).Limit(limit).Find(&models, " last_name =? ", lastName).Error
	return models, err

}

func (m *Model) findByAgeBetween(min, max, limit, offset int) (models []*Model, err error) {
	models = []*Model{}

	err = GetDb().Table("model").Offset(offset).Limit(limit).Find(&models, " age BETWEEN ? AND ? ", min, max).Error
	return models, err

}

```


## License

gpa is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text


