package strategy

import (
	"testing"
)

//装饰模式，动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更加灵活。
//1、优点：装饰对象和真实对象实现相同接口，把真实对象的装饰分离出来，有效地把类的核心职责和装饰功能分开。
//2、缺点：常常会造成有大量的小类，数量太多。
//3、使用场景：当我们需要为某个现有对象，动态地增加功能或者职责时，可以使用装饰模式。
func Test_main(t *testing.T) {
	person := new(Person)
	person.Name = "小小"

	tshirts := new(TShirts)
	tshirts.Decorate(person)
	bigTrouser := new(BigTrouser)
	bigTrouser.Decorate(tshirts)
	bigTrouser.Show()
}
