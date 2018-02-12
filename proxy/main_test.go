package strategy

import (
	"testing"
)

//代理模式，为其他对象提供一种代理以控制对这个对象的访问
//1、优点：代理模式能让代理对象和真正被调用的对象分离，在一定程度上降低了系统的耦合度；
//		  可以起到保护目标对象的作用，客户端不需要知道目标对象是什么
//2、缺点：在客户端和目标对象中加了个代理类，会造成请求处理速度变慢；增加系统的复杂度。
//3、使用场景：远程代理，客户端可以访问在远程主机上的对象，为一个对象在不同的地址空间提供局部代表；
//			 虚拟代理，对于一些占用系统资源较多或加载时间较长的对象，可以给这些对象提供一个虚拟代理；
//			 安全代理，用来控制真实对象访问时的权限；
//			 动态代理，利用反射机制，在运行时指定需要创建的对象
func Test_main(t *testing.T) {
	schoolGirl := SchoolGirl{Name: "羊羊"}
	proxy := NewProxy(&schoolGirl)
	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()
}
