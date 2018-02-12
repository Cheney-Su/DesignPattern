package strategy

//活动类接口
type ICashSuper interface {
	acceptCash(money float64) float64
}
