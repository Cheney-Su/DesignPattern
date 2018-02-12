package abstractfactory

import "fmt"

type SqlServerDepartment struct {
}

func (*SqlServerDepartment) Insert(user *Department) bool {
	fmt.Println("sqlServer成功插入一条部门记录")
	return true
}

func (*SqlServerDepartment) GetDepartment(id int) *Department {
	fmt.Println("sqlServer成功查询一条部门记录")
	return nil
}
