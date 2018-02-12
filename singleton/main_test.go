package singleton

import (
	"testing"
)

//单例模式，保证一个类仅有一个实例，并提供一个访问它的全局访问点。
//1、优点：让类自身负责保存它的唯一实例，这个类可以保证没有其他实例可以被创建；节约系统资源。
//2、缺点：单例类的责任过重，在一定程度上违反了单一职责原则；代码实现需要注意线程安全问题
//3、使用场景：需要频繁实例化后销毁的对象；创建对象耗时耗资源，但又经常用到的对象；频繁访问数据库或文件的对象。
func Test_main(t *testing.T) {
	instance := GetUnSafeInstance()
	instance.Test()

	instance = GetOnceLockInstance()
	instance.Test()
}
