package simplefactory

//相乘结构体
type Multi struct {
	First float64
	Second float64
}

//实现operation接口的GetResult
func (this *Multi) GetResult() float64  {
	return this.First * this.Second
}

//实现operation接口的Init
func (this *Multi) Init(first, second float64)  {
	this.First = first
	this.Second = second
}
