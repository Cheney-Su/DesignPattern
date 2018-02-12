package singleton

import (
	"fmt"
)

type singleton struct {
}

var instance *singleton

//不安全的单例模式
func GetUnSafeInstance() *singleton {
	if instance == nil {
		return &singleton{}
	}
	return instance
}

func (*singleton) TestUnSafe() {
	fmt.Println(111)
}