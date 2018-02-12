package strategy

//满减价格类
type CashReturn struct {
	MoneyCondition float64
	MoneyReturn    float64
}

//满减价格类的构造方法
func NewCashReturn(moneyCondition, moneyReturn float64) *CashReturn {
	return &CashReturn{MoneyCondition: moneyCondition, MoneyReturn: moneyReturn}
}

//满减价格类 满MoneyCondition减MoneyReturn后的价格
func (this *CashReturn) acceptCash(money float64) float64 {
	result := 0.0
	if money > this.MoneyCondition {
		result = money - this.MoneyReturn
	} else {
		result = money
	}
	return result
}
