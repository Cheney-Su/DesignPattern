package strategy

import "fmt"

//垮裤子类
type BigTrouser struct {
	Finery //实现继承关系，继承Finery
}

//实现Icomponent的show方法
func (this *BigTrouser) Show() {
	if this == nil {
		return
	}
	fmt.Print("垮裤 ")
	this.Finery.Show()
}
