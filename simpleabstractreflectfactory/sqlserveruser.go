package abstractfactory

import "fmt"

type SqlServerUser struct {
}

func (*SqlServerUser) Insert(user *User) bool {
	fmt.Println("sqlServer成功插入一条用户记录")
	return true
}

func (*SqlServerUser) GetUser(id int) *User {
	fmt.Println("sqlServer成功查询一条用户记录")
	return nil
}
