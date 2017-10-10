package example

type Model struct {
	Id       string
	Name     string
	LastName string
}

func (m *Model) findById(id string) (model *Model) {

}

func (m *Model) findByName(name string) (models []*Model) {

}

func (m *Model) findByNameAndLastname(name, lastName string) (models []*Model) {

}

func (m *Model) deleteByName(name string) (*Model) {

}

func (m *Model) deleteByIdAndLastName(id, lastName string) (*Model) {

}
