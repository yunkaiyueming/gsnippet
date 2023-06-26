package main

var msg string
var done bool

func setup() {
	msg = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	println(msg)
}
