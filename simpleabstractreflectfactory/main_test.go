package abstractfactory

import (
	"testing"
)

//简单工厂整合抽象工厂实现(反射的方式未实现)
//1、优点：很好的解决了抽象工厂每次新增产品都需要添加工厂类的繁琐，用一个简单工厂管理所有的工厂类，可以用反射去解决switch的问题。
//2、使用场景：现在很多框架都是用这套去实现具体的业务逻辑。
func Test_main(t *testing.T) {
	user := User{Id: 1, Name: "test"}
	department := Department{Id: 2, Name: "test2"}
	dataAccess := DataAccess{}
	iUser := dataAccess.CreateUser()
	iDepartment := dataAccess.CreateDepartment()
	iUser.Insert(&user)
	iUser.GetUser(1)
	iDepartment.Insert(&department)
	iDepartment.GetDepartment(2)

}
