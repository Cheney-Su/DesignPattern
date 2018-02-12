package singleton

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var once sync.Once

//每个携程都需要获取锁
func GetLockInstance() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		return &singleton{}
	}
	return instance
}

//避免每个携程都需要获取锁，做双重校验
func GetDoubleLockInstance() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			return &singleton{}
		}
	}
	return instance
}

//通过sync.Once帮助，用一个unit32的变量作为标识位，如果为0的时候，执行方法后用原子操作将标识位置为1，比较推荐这种方法
func GetOnceLockInstance() *singleton {
	once.Do(func() {
		if instance == nil {
			instance = &singleton{}
		}
	})
	return instance
}

func (*singleton) Test() {
	fmt.Println(111)
}
