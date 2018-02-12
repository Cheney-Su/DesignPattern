package abstractfactory

type IDepartment interface {
	Insert(department *Department) bool
	GetDepartment(id int) *Department
}
