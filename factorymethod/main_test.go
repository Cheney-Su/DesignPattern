package factorymethod

import (
	"testing"
	"fmt"
)

//工厂方法模式，定义一个用于创建对象的接口，让子类决定实例化哪个类。工厂方法使一个类的实例化延迟到其子类。
//1、优点：工厂方法很好的减轻了工厂类的责任，用户只需要知道产品所需要的具体工厂，无需关心具体的创建过程；
//		  新增子产品功能时，只需要添加对应的产品类和工厂类且继承对应的接口，无需对原工厂进行修改，很好地符合了开闭原则。
//2、缺点：每次新增子产品功能时，都需要添加一个对应的产品类和工厂类，使得系统中类的个数成倍增多，增加系统的复杂度，同时也增加系统具体类的依赖。
//3、使用场景：类不知道自己要创建哪一个对象，客户端需要清楚自己创建哪一个对象，并用子类来指定创建的对象。
func Test_main(t *testing.T) {
	//相加工厂
	iAddFactory := new(AddFactory)
	iAddOperation := iAddFactory.CreateOperation()
	iAddOperation.Init(1, 2)
	fmt.Println(iAddOperation.GetResult())
	//相乘工厂
	iMultiFactory := new(MultiFactory)
	iMtltiOperation := iMultiFactory.CreateOperation()
	iMtltiOperation.Init(1, 2)
	fmt.Println(iMtltiOperation.GetResult())
}
