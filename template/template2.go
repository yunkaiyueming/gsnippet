package main

import (
	"fmt"
	htempate "html/template"
	"os"
	ttempate "text/template"
)

func main() {
	htmlOut()
	fmt.Println()
	textOut()
}

func textOut() {
	t, err := ttempate.New("foo").Parse(`{{define "T"}}Hello_text, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>") //text/template不转义
	if err != nil {
		fmt.Println(err.Error())
	}
}

func htmlOut() {
	t, err := htempate.New("foo").Parse(`{{define "T"}}Hello_template, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>") //html/template转义
	if err != nil {
		fmt.Println(err)
	}
}
