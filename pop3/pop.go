package main

import (
	"fmt"

	pop "github.com/taknb2nch/go-pop3"
)

func main() {
	server_addr := ""
	c, err := pop.Dial(server_addr)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		c.Quit()
		c.Close()
	}()

	c.User("")
	c.Pass("")

	num, size, _ := c.List(2)
	fmt.Println(num, size)

	infos, _ := c.ListAll()
	for _, info := range infos {
		fmt.Println(info.Number)
	}

	//emailMsg, _ := c.Retr(1) //quoted-printable
	//fmt.Println(emailMsg)
	//fmt.Printf("%s", emailMsg)
	//如何解析获取到的经过quoted-printable编码的邮件信息？？？go的
}
