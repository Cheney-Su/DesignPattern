package strategy

//试卷A
type TestPaperA struct {
	TestPaper
}

//重写Answer1方法
func (this *TestPaperA) Answer1() string {
	return "b"
}

//重写Answer2方法
func (this *TestPaperA) Answer2() string {
	return "c"
}

//重写Answer3方法
func (this *TestPaperA) Answer3() string {
	return "a"
}

