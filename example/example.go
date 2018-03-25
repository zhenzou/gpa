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

// result can come without a name
func (m *Model) save(model *Model) (err error) {

}

// result can come without a name
func (m *Model) Update(model *Model, Id int64) (err error) {

}

// result can come without a name
func (m *Model) UpdateById(model *Model, Id int64) (err error) {

}

// must not use *[]*Model
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
