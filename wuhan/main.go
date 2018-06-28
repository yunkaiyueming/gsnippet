package main

import (
	"fmt"

	"encoding/json"
	//"time"

	"github.com/astaxie/beego/httplib"
)

func main() {
	url := "http://liuyan.cjn.cn/threads/queryThreadsList"

	req := httplib.Post(url)
	req.Param("fid", "18")
	req.Param("lastTid", "159239")
	rets, err := req.String()
	if err != nil {
		fmt.Print(err.Error())
	}

	data := make(map[string]interface{})
	json.Unmarshal([]byte(rets), &data)
	fmt.Println(data)

	data_2 := data["responseData"].(string)

	//data_3 := data_2[0]["threads"]
	fmt.Println(data_2)
}
