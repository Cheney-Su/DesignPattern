package abstractfactory

import "fmt"

type MysqlUser struct {
}

func (*MysqlUser) Insert(user *User) bool {
	fmt.Println("mysql成功插入一条用户记录")
	return true
}

func (*MysqlUser) GetUser(id int) *User {
	fmt.Println("mysql成功查询一条用户记录")
	return nil
}
