package main

type Study interface {
	LearnYuWen()
	LearnEnglish()
}

type Student struct {
	Name       string
	goToSchool func(age int) int //字段是函数
	Study                        //字段是接口，可以引用赋值实现该接口的实例
}

func GoToSchool(age int) int {
	print(age, "gotoschool11111\n")
	return age
}

func GoToSchool2(age int) int {
	print(age, "gotoschool2222\n")
	return age
}

func (s Student) Fly() string {
	print("im flying")
	return "im flying"
}

type Skill struct{}

func (s Skill) LearnYuWen() {
	print("stduy yuwenk ok 11111\n")
}
func (s Skill) LearnEnglish() {
	print("stduy yuwenk very good11111\n")
}

type Skill2 struct{}

func (s Skill2) LearnYuWen() {
	print("stduy yuwenk ok 2222222\n")
}
func (s Skill2) LearnEnglish() {
	print("stduy yuwenk very good 2222222\n")
}

// 实现Study接口的类型都可以传递和引用
func adapter(s Study) {
	s.LearnYuWen()
}

// 调用Student类型的方法
func adapterGoSchool(s Student) {
	s.goToSchool(22)
}

// Go语言鼓励了基于接口的编程，提供了更灵活和简洁的代码设计。任何类型只要满足接口的方法要求，都可以被传递和使用，无需显式地指定类型与接口的关系。这使得代码更具扩展性和复用性
func main() {
	zs := Student{
		Name:       "zs",
		goToSchool: GoToSchool,
	}

	zs2 := Student{
		Name:       "zs2",
		goToSchool: GoToSchool2,
	}
	//调用结构体的不同方法实现
	adapterGoSchool(zs)
	adapterGoSchool(zs2)

	//结构体实现接口 赋予接口体的不同实现方法，实现不同功能目的。
	zs.Study = Skill{}
	zs.LearnYuWen()

	zs2.Study = Skill2{}
	zs2.LearnYuWen()

	adapter(Skill{}) //接口多态 结构体有不同接口实现，调用interface传参的地方，可以来传参这些结构体，调用这些结构体的不同实现
	adapter(Skill2{})

}
