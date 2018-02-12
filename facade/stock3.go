package strategy

import "fmt"

//股票3结构体
type Stock3 struct {
}

//实现sell方法
func (*Stock3) Sell() {
	fmt.Println("股票3卖出")
}

//实现buy方法
func (*Stock3) Buy() {
	fmt.Println("股票3买入")
}
