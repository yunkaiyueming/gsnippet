package main

import (
	"fmt"
	"sort"
)

var intArr = []int{1, 4, 6, 8, 0, 5, 34}
var float64Arr = []float64{1.2, 22.2, 4.54, 66, 7.8, 3, 5}
var StrArr = []string{"hello", "word", "go"}

type mapItem []map[string]string

var mapData mapItem = []map[string]string{
	{"date": "2017-03", "cost": "22.35", "pro": "zhangsan"},
	{"date": "2017-11", "cost": "52.35", "pro": "lisi"},
	{"date": "2017-09", "cost": "242.35", "pro": "wangwu"},
	{"date": "2017-10", "cost": "212.35", "pro": "heiliu"},
}

func (d mapItem) Len() int {
	return len(d)
}
func (d mapItem) Less(i, j int) bool {
	return d[i]["date"] < d[j]["date"]
}
func (d mapItem) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	if !sort.IntsAreSorted(intArr) {
		sort.Ints(intArr) //升序
		fmt.Println(intArr)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(intArr)))
	fmt.Println(intArr)

	fmt.Println("====================")

	if !sort.Float64sAreSorted(float64Arr) {
		sort.Float64s(float64Arr) //升序
		fmt.Println(float64Arr)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(float64Arr)))
	fmt.Println(float64Arr)

	if !sort.IsSorted(mapData) {
		sort.Sort(mapData) //按日期升序
		fmt.Println(mapData)
	}
	sort.Sort(sort.Reverse(mapData))
	fmt.Println(mapData) //按日期降序
}
