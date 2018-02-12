package abstractfactory

type SqlServerUserFactory struct {
}

func (*SqlServerUserFactory) CreateUser() IUser {
	return new(SqlServerUser)
}

func (*SqlServerUserFactory) CreateDepartment() IDepartment {
	return new(SqlServerDepartment)
}