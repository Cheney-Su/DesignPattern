package simplefactory

//简单工厂类
type SimpleFactory struct {
}

//创建一个操作类对象，根据操作符类型创建对应的对象
func (*SimpleFactory) createOperation(operationStr string) IOperation {
	var operation IOperation
	switch operationStr {
	case "+":
		operation = new(Add)
	case "*":
		operation = new(Multi)
	default:
		operation = new(Add)
	}
	return operation
}
