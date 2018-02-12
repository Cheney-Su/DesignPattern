package strategy

//打折活动类
type CashRebate struct {
	Rebate float64
}

//打折活动类的构造方法
func NewCashRebate(rebate float64) *CashRebate {
	return &CashRebate{Rebate: rebate}
}

//打折活动类打折后的价格
func (this *CashRebate) acceptCash(money float64) float64 {
	return money * this.Rebate
}
