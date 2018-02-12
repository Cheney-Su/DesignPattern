package factorymethod

//定义一个工厂接口，共有方法是创建操作对象
//不需要知道具体的实现方法，让实现接口的结构体自己实现
type IFactory interface{
	CreateOperation() IOperation
}

