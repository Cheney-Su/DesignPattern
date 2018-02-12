package strategy

//正常活动类
type CashNormal struct {
}

//正常活动类的构造方法
func NewCashNormal() *CashNormal {
	return &CashNormal{}
}

//正常活动类获取最终的价格
func (*CashNormal) acceptCash(money float64) float64 {
	return money
}
