package main

import (
	"encoding/json"
	"fmt"
)

type UserServants struct {
	Uid       int64  `json:"uid"`
	Info      string `json:"info"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserServants2 struct {
	Uid       int64              `json:"uid"`
	info      map[string]Servant `json:"info"`
	UpdatedAt int64              `json:"updated_at"`
}

type Servant struct {
	Lv     int64  `json:"lv"`
	Hasexp int64  `json:"hasexp"`
	Total  string `json:"total"`
}

func main() {
	s := Servant{
		Lv:     1,
		Hasexp: 1,
		Total:  "1231",
	}

	sJson, _ := json.Marshal(map[string]Servant{"1": s})
	sJsonStr := string(sJson)
	data1 := UserServants{Uid: 111, Info: sJsonStr, UpdatedAt: 123131}

	d, _ := json.Marshal(data1)
	fmt.Println(string(d))
}
