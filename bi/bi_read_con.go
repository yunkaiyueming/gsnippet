package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/yunkaiyueming/Starfire/cron"
)

func main() {
	recordRoutine()
	beego.SetLogger("multifile", `{"filename":"./run.log"}`)
	cron.CreateCronJob(5*time.Second, recordRoutine, "recordRoutine")
}

func recordRoutine() string {
	resp, err := httplib.Get("http://bi.tt.com:65534/v1/maintain/get_routines").String()
	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &ret)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}

	fmt.Println(ret["data"])
	logs.Info(ret["data"])
	return strconv.FormatFloat(ret["data"].(float64), 'f', 2, 64)
}
