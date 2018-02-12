package strategy

//建造者结构体，初始化需要传入person参数
type Builder struct {
	Person IBuilder
}

//调用建造者person的实现方法
func (this *Builder) CreatePerson() {
	if this == nil {
		return
	}
	this.Person.Head()
	this.Person.Body()
	this.Person.Foot()
	this.Person.Hand()
}
