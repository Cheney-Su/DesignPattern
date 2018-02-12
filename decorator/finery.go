package strategy

//衣服父类
type Finery struct {
	person IComponent
}

//设置与人的关系
func (this *Finery) Decorate(per IComponent) {
	this.person = per
}

//实现Icomponent的show方法
func (this *Finery) Show() {
	if this == nil {
		return
	}
	if this.person != nil {
		this.person.Show()
	}
}
