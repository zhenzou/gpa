package example

type Model struct {
	Id       string
	Name     string
	LastName string
}

func (m *Model) findById() (model *Model) {

}

func (m *Model) findByName() (models []*Model) {

}

func (m *Model) findByNameAndLastname() (models []*Model) {

}

func (m *Model) deleteByName() (*Model) {

}

func (m *Model) deleteByIdAndLastName(id, lastName string) (*Model) {

}
