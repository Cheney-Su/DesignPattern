package strategy

import "fmt"

//胖人结构体
type Fat struct {
}

//实现head方法
func (*Fat) Head() {
	fmt.Println("我的头很胖")
}

//实现body方法
func (*Fat) Body() {
	fmt.Println("我的身体很胖")
}

//实现foot方法
func (*Fat) Foot() {
	fmt.Println("我的脚很胖")
}

//实现hand方法
func (*Fat) Hand() {
	fmt.Println("我的收很胖")
}
