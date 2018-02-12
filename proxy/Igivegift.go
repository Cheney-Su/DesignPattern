package strategy

//送礼物接口，GiveDolls，GiveFlowers，GiveChocolate三个待实现方法
type IGiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}