package abstractfactory

import "fmt"

type MysqlDepartment struct {
}

func (*MysqlDepartment) Insert(user *Department) bool {
	fmt.Println("mysql成功插入一条部门记录")
	return true
}

func (*MysqlDepartment) GetDepartment(id int) *Department {
	fmt.Println("mysql成功查询一条部门记录")
	return nil
}
