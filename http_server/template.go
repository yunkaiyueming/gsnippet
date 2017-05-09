package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	TestEscapge()
}

func templateString() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}

func TestEscapge() {
	str := template.HTMLEscapeString("<br>hello</br>")
	fmt.Println(str)

	jsStr := template.JSEscapeString("<script>alert('123')</script>")
	fmt.Println(jsStr)
}
