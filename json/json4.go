package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	var mapData map[int64][]int64
	mapData = parseAppSet(js_str)
	fmt.Println(mapData) //功能更强大
	println(mapData)
}

var js_str = `
{"1013":[10113,10213,10313,10413,10513,10613,10713],"1014":[10114,10214,10314,10414],"1018":[1018,10118],"1019":[1019],"1024":[10124,10224],"1025":[10125,10225,10325,10425,10525],"1026":[10126,10226,10526],"1027":[10127,10227,10327],"1028":[1028],"1029":[10129,10229,10329,10429],"1031":[10131,10231,10331],"1032":[10232],"1033":[10133,10233],"1034":[10134,10234],"1036":[10136,10236,10336,10436,10536,10636,10736,10836],"1040":[10140,10240]}
`

func parseAppSet(appIds string) map[int64][]int64 {
	var appConfig interface{}
	appMap := make(map[int64][]int64)
	json.Unmarshal([]byte(appIds), &appConfig)
	var subInt []int64

	appData := appConfig.(map[string]interface{})
	for big_app_id, sub_app_ids := range appData {
		for _, v := range sub_app_ids.([]interface{}) {
			subInt = append(subInt, int64(v.(float64)))
		}
		bigInt, _ := strconv.Atoi(big_app_id)
		appMap[int64(bigInt)] = subInt
	}
	return appMap
}

func test() {
	var appConfig interface{}
	//未知结构体的json解析
	json.Unmarshal([]byte(js_str), &appConfig)

	appData := appConfig.(map[string]interface{})
	for big_app_id, sub_app_ids := range appData {
		println(big_app_id)
		for i, v := range sub_app_ids.([]interface{}) {
			println(i, int64(v.(float64)))
		}
	}
}
