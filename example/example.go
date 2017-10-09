package example

type Model struct {
	Id       string
	Name     string
	LastName string
}

func (m *Model) findModelById() (*Model) {

}

func (m *Model) findModelByName() ([]*Model) {

}

func (m *Model) deleteByName() (*Model) {

}

func (m *Model) deleteByIdAndLastName(id, lastName string) (*Model) {

}
