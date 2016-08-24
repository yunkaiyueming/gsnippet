package main

import (
	"fmt"
	"sync"
)

type User struct {
	Mu    sync.Mutex
	Name  string
	Age   int
	Likes []string
}

func main() {
	u := User{Name: "xhc", Age: 12}
	u.UpdateUserInfo()
	fmt.Println(u)
}

func (u *User) UpdateUserInfo() {
	u.Mu.Lock()
	u.Name = "bb"
	u.Mu.Unlock()
}
