package main

import "fmt"

var mapData map[string]int

func main() {
	data := map[string]string{"k1": "v1", "k3": "v3"}
	fmt.Println(data)
	changeMap(data)
	fmt.Println(data)

}

func InitMap() {
	//make方式创建
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//map类型创建
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func TestMap() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

func TestMap2() {
	//var mapData map[string]int = map[string]int{}

	//mapData = make(map[string]int)
	fmt.Println(mapData)
	mapData["s"] = 4
	fmt.Println(mapData)
}

func changeMap(data map[string]string) {
	data["k2"] = "d2"
}
