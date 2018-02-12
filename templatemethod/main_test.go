package strategy

import (
	"testing"
)

//模板方法模式，定义一些操作中的算法的骨架，并将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
//1、优点：模板方法通过把不变的行为搬到父类，去除了子类的重复代码，有子类实现自己的业务代码；
//		  通过父类调用子类的实现操作，通过子类扩展新的功能，复合开闭原则
//2、缺点：每个不同的实现都需要编写一个子类，类的数目增多
//3、使用场景：在某些类的算法中，用了相同的方法，造成代码的重复；控制子类扩展，子类必须遵守算法规则。
func Test_main(t *testing.T) {
	ITestPaperA := TestPaperA{}
	ITestPaperA.ITestPaper = &ITestPaperA //不是很懂这里为什么要取址，取址是什么意义呢？
	// 结构体定义的是值对象，ITestPaper是interface类型，对应的应该是指针对象，那么在赋值过程就需要将结构体取址才能赋予interface
	ITestPaperA.TestQuestion1()
	ITestPaperA.TestQuestion2()
	ITestPaperA.TestQuestion3()

	//ITestPaperB := NewTestPaperB()
	ITestPaperB := TestPaperB{}
	ITestPaperB.ITestPaper = ITestPaperB //将TestPaperB中方法的指针去掉这里就可以不用取址
	ITestPaperB.TestQuestion1()
	ITestPaperB.TestQuestion2()
	ITestPaperB.TestQuestion3()
}
