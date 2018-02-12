package strategy

//代理类
type Proxy struct {
	Persuit Persuit
}

//代理类的构造方法，这里实例被代理类，重点
func NewProxy(schoolGirl *SchoolGirl) *Proxy {
	persuit := NewPersuit(schoolGirl)
	return &Proxy{Persuit: *persuit}
}

//实现接口的送洋娃娃方法，调用被代理类的实现
func (this *Proxy) GiveDolls() {
	this.Persuit.GiveDolls()
}

//实现接口的送鲜花方法，调用被代理类的实现
func (this *Proxy) GiveFlowers() {
	this.Persuit.GiveFlowers()
}

//实现接口的送巧克力方法，调用被代理类的实现
func (this *Proxy) GiveChocolate() {
	this.Persuit.GiveChocolate()
}
