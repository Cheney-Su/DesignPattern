package simplefactory

//定义一个操作符接口，共有方法是初始化参数方法Init和获取操作结果GetResult
//不需要知道具体的实现方法，让实现接口的结构体自己实现
type IOperation interface{
	Init(first, second float64)
	GetResult() float64
}



