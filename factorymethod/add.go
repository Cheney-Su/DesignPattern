package factorymethod

//相加结构体
type Add struct {
	first  float64
	secend float64
}

//实现operation接口的GetResult
func (this *Add) Init(first, second float64) {
	this.first = first
	this.secend = second
}

//实现operation接口的Init
func (this *Add) GetResult() float64 {
	return this.first + this.secend
}
