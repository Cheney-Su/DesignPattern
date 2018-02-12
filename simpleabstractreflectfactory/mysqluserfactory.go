package abstractfactory

type MysqlUserFactory struct {
}

func (*MysqlUserFactory) CreateUser() IUser {
	return new(MysqlUser)
}

func (*MysqlUserFactory) CreateDepartment() IDepartment {
	return new(MysqlDepartment)
}
