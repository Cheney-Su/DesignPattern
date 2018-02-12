package test

import (
	"testing"
	"fmt"
)

type AAA struct {
	aa int
}


//go函数方法的调用都是值传递，
func Test_test(t *testing.T) {
	aaa := AAA{aa: 2}
	ChangeStruct(aaa)
	fmt.Println("函数传结构体不会影响结果", aaa)	//值
	ChangeStructIn(&aaa)
	fmt.Println("函数传结构体指针会影响结果", aaa)	//引用，地址
	str := "test"
	ChangeString(str)
	fmt.Println("函数传字符串不会影响结果", str)	//值
	arr := [5]string{"1","2","3"}	//长度为5的数组
	fmt.Println("参数外arr[0]的地址", &arr[0]) 	//因为传数组其实相当于会在方法中数组是一个新的地址，但是切片确实引用地址，所以切片会修改原有数组的值
	ChangeArr(arr)
	fmt.Println("函数传数组不会影响结果", arr)	//值
	slice := []string{"1","2","3"}	//切片
	fmt.Println("参数外slice[0]的地址", &slice[0])
	ChangeSlice(slice)
	fmt.Println("函数传切片会影响结果", slice)	//值
	ChangeArrIn(&arr)
	fmt.Println("函数传指针数组会影响结果", arr)	//引用，地址

	bbb := [5]AAA{aaa, aaa}
	ChangeArrSt(bbb)
	fmt.Println("函数传数组结构体不会影响结果", bbb)	//值

}

//改变结构体的值
func ChangeStruct(a AAA)  {
	a.aa = 1
}

//改变指针结构体的值
func ChangeStructIn(a *AAA)  {
	a.aa = 1
}

//改变字符串的值
func ChangeString(a string)  {
	a = "test1"
}

//改变数组的值
func ChangeArr(arr [5]string) {
	fmt.Println("参数内a[0]的地址", &arr[0])
	arr[0] = "3"
}

//改变数组结构体的值
func ChangeArrSt(arr [5]AAA) {
	arr[0].aa = 3
}

//改变指针数组的值
func ChangeArrIn(arr *[5]string) {
	(*arr)[0] = "4"
}

func ChangeSlice(slice []string)  {
	fmt.Println("参数内slice[0]的地址", &slice[0])
	slice[0] = "15"
}

