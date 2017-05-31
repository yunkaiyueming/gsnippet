package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	pwd  string `json:"pwd"`
	sex  int64  `json:"sex"`
}

type Teacher struct {
	Person
	CourseName string `json:"course_name"`
}

type MapClass map[string]string

func (p Person) String() string {
	return fmt.Sprintf("id:%s,name:%s,age:%s,pwd:%s,sex:%s \n", p.Id, p.Name, p.Age, p.pwd, p.sex)
}

func (p Person) DoWork() string {
	return fmt.Sprintf("i'm a person,do work hard")
}

func (t *Teacher) DoWork() string {
	return fmt.Sprintf("i'm a teacher,teacher student")
}

func (m MapClass) RangeData() string {
	var out string
	for k, v := range m {
		out += "key:" + k + ",v:" + v + "\t"
	}
	return out
}

func reflectStruct() {
	p := Person{321414, "张三", 25, "abcdefpwd", 0}
	t := new(Teacher)
	t.Person = p
	t.CourseName = "语文"

	fmt.Println("-------------------------type------------------------")
	GetAllTypInfo(p)

	fmt.Println("-------------------------val-------------------------")
	GetAllValInfo(p)

	fmt.Println("-------------------------type------------------------")
	GetAllTypInfo(*t) //是指类型才行，是指针变量的话，会panic

	fmt.Println("-------------------------val-------------------------")
	GetAllValInfo(*t)

	//当前面的CanSet是p.Elem()(实际上就是*p)，它就是可以寻址的,可寻址的才可以改变
	fmt.Println(reflect.ValueOf(&(p.Name)).Elem().CanSet())
}

func reflectMap() {
	classMsg := make(map[string]string)
	classMsg["one"] = "一年级"
	classMsg["two"] = "二年级"
	classMsg["three"] = "三年级"
	//	fmt.Println("-------------------------type------------------------")
	//	GetAllTypInfo(classMsg)//不是结构体

	//	fmt.Println("-------------------------val-------------------------")
	//	GetAllValInfo(classMsg)
	fmt.Println(classMsg)
}

func main() {
	reflectMap()
	reflectStruct()
}

func GetAllTypInfo(i interface{}) {
	typ := reflect.TypeOf(i)
	fmt.Println(typ)
	fmt.Println(typ.NumField())
	fmt.Println(typ.NumMethod())

	//遍历结构体的字段
	fmt.Println("---------获取所有字段-------------")
	for w := 0; w < typ.NumField(); w++ {
		fmt.Println("-----", i, "----")
		field := typ.Field(w)
		fmt.Println(field.Index, field.Name, field.Offset, field.PkgPath, field.Tag, field.Type)
	}

	// 遍历对象中的方法
	fmt.Println("---------获取所有方法-------------")
	for w := 0; w < reflect.TypeOf(i).NumMethod(); w++ {
		method := reflect.TypeOf(i).Method(w)
		fmt.Println(method.Type, method.Name, method.Type.NumIn(), method.Type.In(0)) // func(*main.MyStruct) string,GetName,参数个数，类型
	}

	met, _ := typ.MethodByName("DoWork")
	fmt.Println(met.Func, met.Index, met.Name, met.PkgPath, met.Type)
}

func GetAllValInfo(i interface{}) {
	val := reflect.ValueOf(i)
	fmt.Println(val, i)

	fmt.Println(val.NumMethod())
	fmt.Println(val.NumField())

	for i := 0; i < val.NumMethod(); i++ {
		b := val.Method(i).Call([]reflect.Value{})
		fmt.Println(i, b)
	}

	//当前面的CanSet是一个指针的时候（p）它是不可寻址的，但是当是p.Elem()(实际上就是*p)，它就是可以寻址的
	fmt.Println("name", val.FieldByName("Name").CanSet())
	fmt.Println("id", val.FieldByName("id").CanSet())
	fmt.Println("sex", val.FieldByName("sex").CanSet())

}
