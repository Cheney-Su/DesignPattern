package abstractfactory

type IFactory interface {
	CreateUser() IUser
	CreateDepartment() IDepartment
}
