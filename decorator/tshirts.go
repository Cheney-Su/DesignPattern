package strategy

import "fmt"

//T恤子类
type TShirts struct {
	Finery	//实现继承关系，继承Finery
}

//实现Icomponent的show方法
func (this *TShirts) Show() {
	if this == nil {
		return
	}
	fmt.Print("T恤 ")
	this.Finery.Show()
}
