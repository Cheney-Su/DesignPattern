package simplefactory

//相加结构体
type Add struct {
	First float64
	Second float64
}

//实现operation接口的GetResult
func (this *Add) GetResult() float64  {
	return this.First + this.Second
}

//实现operation接口的Init
func (this *Add) Init(first, second float64)  {
	this.First = first
	this.Second = second
}

