package abstractfactory

import (
	"testing"
)

//抽象工厂模式,提供一个创建一系列相关或相互依赖对象的接口，而无需指定他们具体的类。
//1、优点：当一个产品族的多个对象被设计到一起工作时，它能保证客户端始终只使用同一个产品族的对象；
//		  它让具体的创建实例与客户端分离，客户端通过他们的抽象接口操纵实例，产品的具体类名也没具体工厂的实现分离。
//2、缺点：增加新的产品结构较为复杂，需要修改抽象工厂接口和新增具体实现类和工厂类，对开闭原则的支持程倾斜性。
//3、使用场景：当需要创建的对象是一系列相互关联或者相互依赖的产品族时，便可以使用抽象工厂模式。
func Test_main(t *testing.T) {
	user := User{Id: 1, Name: "test"}
	department := Department{Id: 2, Name: "test2"}
	iMysqlUserFactory := new(MysqlUserFactory)
	iMysqlUser := iMysqlUserFactory.CreateUser()
	iMysqlDepartment := iMysqlUserFactory.CreateDepartment()
	iMysqlUser.Insert(&user)
	iMysqlUser.GetUser(1)
	iMysqlDepartment.Insert(&department)
	iMysqlDepartment.GetDepartment(2)

	iSqlServerUserFactory := new(SqlServerUserFactory)
	iSqlServerUser := iSqlServerUserFactory.CreateUser()
	iSqlServerDepartment := iSqlServerUserFactory.CreateDepartment()
	iSqlServerUser.Insert(&user)
	iSqlServerUser.GetUser(1)
	iSqlServerDepartment.Insert(&department)
	iSqlServerDepartment.GetDepartment(2)
}
