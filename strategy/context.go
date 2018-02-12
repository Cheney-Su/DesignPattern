package strategy

//活动上下文工厂类
type Context struct {
	cashType  int
	cashSuper ICashSuper
}

//cashType 1:正常活动 2:打8折 3:满300-100
func NewContext(cashType int) *Context {
	var cashSuper ICashSuper
	switch cashType {
	case 1:
		cashSuper = NewCashNormal()
	case 2:
		cashSuper = NewCashRebate(0.8)
	case 3:
		cashSuper = NewCashReturn(300, 100)
	}
	return &Context{cashType: cashType, cashSuper: cashSuper}
}

//获取抵扣后的价格，调用各个子类的acceptCash方法
func (this *Context) GetResult(money float64) float64 {
	return this.cashSuper.acceptCash(money)
}
