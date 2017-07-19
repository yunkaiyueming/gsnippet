package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-vgo/robotgo"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	i := 0
	beego.SetLogger("multifile", `{"filename":"./run.log","level":3,"separate":["error"]}`)
	logs.Alert(fmt.Sprintf("run_time:%d,start_time:%s\n", i, time.Now().Format("2006-01-02 15:04:05")))

	for {
		lastX, lastY := robotgo.GetMousePos()
		//robotgo.Move(280, 400)
		robotgo.MoveMouse(290, 400)
		robotgo.MouseClick("left", false)
		robotgo.Move(lastX, lastY)

		i++
		fmt.Println(i)
		logs.Alert(fmt.Sprintf("run_time:%d,start_time:%s\n", i, time.Now().Format("2006-01-02 15:04:05")))

		<-ticker.C
	}

}
