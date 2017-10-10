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
