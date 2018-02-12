package strategy

import "fmt"

///试卷模板
type TestPaper struct {
	ITestPaper ITestPaper
}

//模板1，答案根据子类重写接口的方法Answer1
func (this *TestPaper) TestQuestion1() {
	fmt.Println("杨过得到，后来给了郭靖，炼成倚天剑、屠龙刀的玄铁可能是？ a.球磨铸铁 b.马口铁 c.高速合金钢 d.碳素纤维")
	fmt.Println("答案：", this.ITestPaper.Answer1())
}

//模板2，答案根据子类重写接口的方法Answer2
func (this *TestPaper) TestQuestion2() {
	fmt.Println("杨过、程英、陆无双铲除了情花，造成？ a.使这种植物不再害人 b.使一种稀物种灭绝 c.破坏了生态圈的生态平衡 d.造成该地区沙漠化")
	fmt.Println("答案：", this.ITestPaper.Answer2())
}

//模板3，答案根据子类重写接口的方法Answer3
func (this *TestPaper) TestQuestion3() {
	fmt.Println("蓝凤凰致使华山师徒、桃谷六仙呕吐不止，如果你是大夫，会给他们开什么药？ a.阿司匹林 b.牛黄解毒片 c.氟哌酸 d.让他们喝大量生牛奶 e.以上都不对")
	fmt.Println("答案：", this.ITestPaper.Answer3())
}
