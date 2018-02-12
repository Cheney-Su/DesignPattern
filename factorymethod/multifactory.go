package factorymethod

//相加工厂结构体
type MultiFactory struct {
}

//实现IFactory接口的CreateOperation，创建自己的相加对象
func (*MultiFactory) CreateOperation() IOperation  {
	return new(Multi)
}