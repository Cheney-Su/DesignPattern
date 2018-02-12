package strategy

import "fmt"

//简历类
type Resume struct {
	Name     string
	Sex      string
	Age      string
	TimeArea string
	Company  string
}

//简历的构造方法
func NewResume(name string) *Resume {
	return &Resume{Name: name}
}

//设置个人信息
func (this *Resume) SetPersonalInfo(sex, age string) {
	this.Sex = sex
	this.Age = age
}

//设置工作经历
func (this *Resume) SetWorkExperience(timeArea, company string) {
	this.TimeArea = timeArea
	this.Company = company
}

//显示
func (this *Resume) Display() {
	fmt.Println("个人信息：", this.Name, this.Age, this.Sex)
	fmt.Println("工作经历：", this.TimeArea, this.Company)
}

//深复制
func (this *Resume) Clone() *Resume {
	cloneResume := *this
	return &cloneResume
}
