package main

import (
	"fmt"

	"github.com/toolkits/net"
)

func main() {
	p := fmt.Println
	p(net.IntranetIP())        //获取本地机器的局域网IP
	p(net.GrabEphemeralPort()) //获取当前的程序的运行端口

}

/*
range c产生的迭代值为Channel中发送的值，
它会一直迭代知道channel被关闭。上面的例子中如果把close(c)注释掉，程序会一直阻塞在for …… range那一行。
*/
func goroutine_test() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
