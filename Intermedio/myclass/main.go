package myclass

type Employee struct {
	Id       int
	Name     string
	Vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		Id:       id,
		Name:     name,
		Vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.Id = id
}

func (e *Employee) SetName(name string) {

}

func (e *Employee) GetId() int {
	return e.Id
}

func (e *Employee) GetName() string {
	return e.Name
}
