package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string //`json:"fruits"`
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	Json2File(str, "a.json")
}

func test() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
}

func Json2File(jsonStr string, fileName string) {
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()

	enc := json.NewEncoder(file)
	err := enc.Encode(jsonStr)
	if err != nil {
		fmt.Println("Error in encoding gob")
	}
}
