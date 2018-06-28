package main

type Study interface {
	LearnYuWen()
	LearnEnglish()
}

type Student struct {
	Name       string
	goToSchool func(age int) int //字段是函数
	Study                        //字段是接口
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

func adapter(s Study) {
	s.LearnYuWen()
}

func adapterGoSchool(s Student) {
	s.goToSchool(22)
}

func main() {
	zs := Student{
		Name:       "zs",
		goToSchool: GoToSchool,
	}

	zs2 := Student{
		Name:       "zs2",
		goToSchool: GoToSchool2,
	}

	zs.goToSchool(22)

	zs.Study = Skill{}
	zs.LearnYuWen()

	zs.Study = Skill2{}
	zs.LearnYuWen()

	adapter(Skill{}) //多态 实例化实现结构体的不同实现，来调用这些结构体的不同实现
	adapter(Skill2{})

	adapterGoSchool(zs) //赋予接口体的不同方法，实现不同功能目的。
	adapterGoSchool(zs2)
}
