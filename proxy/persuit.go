package strategy

import "fmt"

//追求者类
type Persuit struct {
	SchoolGirl SchoolGirl
}

//追求者类构造方法
func NewPersuit(schoolGirl *SchoolGirl) *Persuit {
	return &Persuit{SchoolGirl: *schoolGirl}
}

//实现接口的送洋娃娃方法
func (this *Persuit) GiveDolls() {
	fmt.Println(this.SchoolGirl.Name + " 送你洋娃娃")
}

//实现接口的送鲜花方法
func (this *Persuit) GiveFlowers() {
	fmt.Println(this.SchoolGirl.Name + " 送你鲜花")
}

//实现接口的送巧克力方法
func (this *Persuit) GiveChocolate() {
	fmt.Println(this.SchoolGirl.Name + " 送你巧克力")
}
