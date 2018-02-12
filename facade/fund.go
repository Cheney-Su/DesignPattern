package strategy

//基金结构体
type Fund struct {
	Stock1 Stock1
	Stock2 Stock2
	Stock3 Stock3
}

//基金的构造方法
func NewFund() *Fund {
	stock1 := Stock1{}
	stock2 := Stock2{}
	stock3 := Stock3{}
	return &Fund{Stock1: stock1, Stock2: stock2, Stock3: stock3}
}

//购买基金
func (this *Fund) BuyFund() {
	this.Stock1.Buy()
	this.Stock2.Buy()
	this.Stock3.Buy()
}

//卖出基金
func (this *Fund) SellFund() {
	this.Stock1.Sell()
	//this.Stock2.Sell()
	this.Stock3.Sell()
}
