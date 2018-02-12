package strategy

import (
	"testing"
)

//原型模式，用原型实例指定创建对象的种类，并通过拷贝这些原型创建新的对象。
//1、优点：简化了对象的创建，简化对象的创建过程，提高效率；可以使用深克隆保持对象的状态
//2、缺点：实现深克隆可能需要比较复杂的代码，要考虑引用克隆；每一个类都必须配备一个克隆方法，而且这个克隆方法需要对类的功能进行通盘考虑。
//3、使用场景：创建对象的成本较大，可以用已有对象进行复制获得；对象的状态变化少，或者占用内存小也可以使用原型模式创建对象。
func Test_main(t *testing.T) {
	resume := NewResume("qiang")
	resume.SetPersonalInfo("man", "26")
	resume.SetWorkExperience("2", "maizuowang")
	resume.Display()

	cloneResume := resume.Clone()
	cloneResume.SetPersonalInfo("man", "28")
	cloneResume.Display()
	//原对象是没有变化的，深复制
	resume.Display()

	copyResume := resume
	copyResume.SetPersonalInfo("man", "22")
	//原对象变化了，浅复制
	resume.Display()
}
