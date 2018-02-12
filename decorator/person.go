package strategy

import "fmt"

//人
type Person struct {
	Name string
}

//实现Icomponent的show方法
func (this *Person) Show() {
	fmt.Printf("装扮的%s", this.Name)
}
