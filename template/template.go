package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	TestEscapge()
	templateString()
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
	str := template.HTMLEscapeString("<br>hello</br>") //转义
	fmt.Println(str)

	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<br>hello</br><script>alert('you have been pwned')</script>") //不转义
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()

	t2 := template.New("foo2")
	t2, err = t2.Parse(`这个是Hello2, <br>hello</br><script>alert('you have been pwned')</script>!`)
	err = t2.Execute(os.Stdout, nil)
	fmt.Println()

	jsStr := template.JSEscapeString("<script>alert('123')</script>") //js转义
	fmt.Println(jsStr)

}
