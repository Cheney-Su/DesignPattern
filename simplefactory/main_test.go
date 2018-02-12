package simplefactory

import (
	"testing"
	"fmt"
)

//简单工厂模式
//1、优点：帮助封装，真正的面向接口编程；解耦，通过简单工厂可以使客户端和具体实现类解耦。
//2、缺点：增加客户端的复杂度，需要知道构造工厂参数的具体意义和功能，增加了客户端的使用难度；不利于扩展，不符合开闭原则
//3、使用场景：工厂类负责创建的对象比较少；客户只知道传入工厂的参数，对于如何创建对象不关心；由于简单工厂很容易违反高内聚责任分配原则，因此一般只在简单的场景下使用。
func Test_main(t *testing.T) {
	factory := SimpleFactory{}
	//创建一个操作对象
	operation := factory.createOperation("*")
	//初始化参数
	operation.Init(1, 2)
	fmt.Println(operation.GetResult())
}
