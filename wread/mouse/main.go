package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

/*

2017/06/11 00:02:27 [A] run_time:157,start_time:2017-06-11 00:02:27

2017/06/11 00:03:29 [A] run_time:158,start_time:2017-06-11 00:03:29

2017/06/11 00:04:28 [A] run_time:159,start_time:2017-06-11 00:04:28

2017/06/11 00:05:29 [A] run_time:160,start_time:2017-06-11 00:05:29

2017/06/11 00:06:27 [A] run_time:161,start_time:2017-06-11 00:06:27

2017/06/11 00:07:27 [A] run_time:162,start_time:2017-06-11 00:07:27

2017/06/11 00:08:28 [A] run_time:163,start_time:2017-06-11 00:08:28

2017/06/11 00:09:26 [A] run_time:164,start_time:2017-06-11 00:09:26

2017/06/11 00:10:27 [A] run_time:165,start_time:2017-06-11 00:10:27*/
func main() {
	robotgo.ScrollMouse(10, "up") //滚动鼠标：up (向上滚动)  down (向下滚动)
	robotgo.MouseClick("left", true)
	robotgo.MouseClick("right", true)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0) //向100,200移动鼠标，xSpeed=10,ySpeed=20
	robotgo.Move(400, 400)
	robotgo.MouseClick("left", false)

	x, y := robotgo.GetMousePos()
	fmt.Println(x, y)

	robotgo.MoveClick(300, y)
}
