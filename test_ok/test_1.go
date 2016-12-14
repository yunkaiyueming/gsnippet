package test_ok

import (
	"fmt"
)

var Bihost string = getHost()

func getHost() string {
	return "xx.com"
}

func aa() {
	fmt.Println(Bihost)
	fmt.Println("Hello, playground")
}
