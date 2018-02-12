package factorymethod

//相加工厂结构体
type AddFactory struct {
}

//实现IFactory接口的CreateOperation，创建自己的相加对象
func (*AddFactory) CreateOperation() IOperation  {
	return new(Add)
}