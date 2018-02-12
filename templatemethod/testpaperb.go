package strategy

//试卷B
type TestPaperB struct {
	TestPaper
}

//func NewTestPaperB() TestPaperB {
//	return TestPaperB{}
//}

//重写Answer1方法
func (this TestPaperB) Answer1() string {
	return "d"
}

//重写Answer2方法
func (this TestPaperB) Answer2() string {
	return "a"
}

//重写Answer3方法
func (this TestPaperB) Answer3() string {
	return "a"
}


