package strategy

import "fmt"

//股票1结构体
type Stock1 struct {
}

//实现sell方法
func (*Stock1) Sell() {
	fmt.Println("股票1卖出")
}

//实现buy方法
func (*Stock1) Buy() {
	fmt.Println("股票1买入")
}
