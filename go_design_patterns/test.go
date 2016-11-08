package main

import (
	"fmt"
)

type User struct {
	UserName string
}

func ProxyUser(u *User) func(userName string) *User {
	fmt.Println("start ProxyUser:")

	return func(userName string) *User {
		u.UserName = userName
		return u
	}
}

func main() {
	u := new(User)
	up := ProxyUser(u)
	u = up("lisi")
	fmt.Println(u)
}
