package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	ParseTest()
	AppendTest()
	FormateToOther()
}

func ParseTest() {
	//Parse 系列函数把字符串转换为其他类型
	a, err := strconv.ParseBool("false")
	checkError(err)
	MyGetType(a)

	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	MyGetType(b)

	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	MyGetType(c)

	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	MyGetType(d)

	e, err := strconv.Atoi("1023")
	checkError(err)
	MyGetType(e)

	fmt.Println(a, b, c, d, e)
}

func AppendTest() {
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}

func FormateToOther() {
	//Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}

func MyGetType(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("type:", v.Type())
}
