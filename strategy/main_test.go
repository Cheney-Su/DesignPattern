package strategy

import (
	"testing"
	"fmt"
)

//策略模式，它定义了算法家族，分别封装起来，让他们之间可以相互替换，此模式让算法的变化，不会影响到使用算法的客户。
//1、优点：将各个算法封装在独立的strategy中，使他们易于切换，易于理解，易于扩展又不相互影响；
//		  当不同的行为堆砌到一个类中，难免会使用条件语句选择，将行为封装到strategy中，可以有效地消除条件语句。
//2、缺点：客户端必须知道context的参数类型对应生成什么策略类。
//3、使用场景：多个类只区别在表现行为不同，可以用strategy模式封装；需要再不同情况下使用不同的策略，或者在未来用其他方式实现；
//			  对客户隐藏策略算法的实现细节
func Test_main(t *testing.T) {
	cashNormalContext := NewContext(1)
	fmt.Println(cashNormalContext.GetResult(100))
	cashRebateContext := NewContext(2)
	fmt.Println(cashRebateContext.GetResult(100))
	cashReturnContext := NewContext(3)
	fmt.Println(cashReturnContext.GetResult(100))
	cashReturn2Context := NewContext(3)
	fmt.Println(cashReturn2Context.GetResult(400))
}
