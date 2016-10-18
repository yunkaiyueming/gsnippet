package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Id     string `json:"userid"`
	Name   string
	Avatar string
}

type User struct {
	UserType int `json:"usertype"`
	UserInfo `json:"user_info"`
}

func main() {
	json_str := `{"usertype":5,
	"user_info":{"userid":"test@qq.com","name":"qq","avatar":"http:\/\/shp.qpic.cn\/bizmp\/elWs7iccclX3WCeUK3ZBycA\/"},
	"corp_info":{"corpid":"wxffdf"}
	}`

	userModel := User{}
	json.Unmarshal([]byte(json_str), &userModel)
	fmt.Println(userModel)
}
