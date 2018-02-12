package strategy

import "fmt"

//股票2结构体
type Stock2 struct {
}

//实现sell方法
func (*Stock2) Sell() {
	fmt.Println("股票2卖出")
}

//实现buy方法
func (*Stock2) Buy() {
	fmt.Println("股票2买入")
}
