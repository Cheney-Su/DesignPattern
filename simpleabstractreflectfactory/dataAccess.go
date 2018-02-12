package abstractfactory

type DataAccess struct {
}

var (
	db = "sqlServer"
)

func (*DataAccess) CreateUser() IUser {
	var iUser IUser
	switch db {
	case "mysql":
		iUser = new(MysqlUser)
	case "sqlServer":
		iUser = new(SqlServerUser)
	}
	return iUser
}

func (*DataAccess) CreateDepartment() IDepartment {
	var IDepartment IDepartment
	switch db {
	case "mysql":
		IDepartment = new(MysqlDepartment)
	case "sqlServer":
		IDepartment = new(SqlServerDepartment)
	}
	return IDepartment
}
