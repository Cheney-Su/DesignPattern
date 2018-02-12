package factorymethod

//相乘结构体
type Multi struct {
	first  float64
	secend float64
}

//实现operation接口的GetResult
func (this *Multi) Init(first, second float64) {
	this.first = first
	this.secend = second
}

//实现operation接口的Init
func (this *Multi) GetResult() float64 {
	return this.first * this.secend
}
