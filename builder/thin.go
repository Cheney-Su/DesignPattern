package strategy

import "fmt"

//瘦人结构体
type Thin struct {
}

//实现head方法
func (*Thin) Head() {
	fmt.Println("我的头很瘦")
}

//实现body方法
func (*Thin) Body() {
	fmt.Println("我的身体很瘦")
}

//实现foot方法
func (*Thin) Foot() {
	fmt.Println("我的脚很瘦")
}

//实现hand方法
func (*Thin) Hand() {
	fmt.Println("我的收很瘦")
}
