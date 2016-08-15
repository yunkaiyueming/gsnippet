package main

type A struct{}

func (this *A) what() string {
	return "a hello"
}

type B struct {
	A
}

func (this *B) what() string {
	return "b hello"
}

type C interface {
	what() string
}

type D struct {
	C
}

func (this *D) say() {
	println(this.what())
}

func NewSay(c C) *D {
	return &D{c}
}

func main() {
	d := NewSay(&B{})
	d.say()
}
