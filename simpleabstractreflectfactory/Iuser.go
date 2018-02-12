package abstractfactory

type IUser interface {
	Insert(user *User) bool
	GetUser(id int) *User
}
