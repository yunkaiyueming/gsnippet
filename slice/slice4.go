package main

import (
	"fmt"
)

func main() {
	//test_delete()
	//test_delete2()
	test3()
}

func test_delete() {
	bigAppIds := []string{"1009", "1011", "1013", "1015", "1017", "1009", "1013"}
	bigAppIdT := make([]string, 0)

	for i, subAppId := range bigAppIds {
		if subAppId == "1013" {
			bigAppIdT = append(bigAppIds[:i], bigAppIds[i+1:]...)
		}
		fmt.Println(i, bigAppIds)
		fmt.Println(i, bigAppIdT)
	}

	fmt.Println(bigAppIdT)
}

func test_delete2() {
	bigAppIds := []string{"1009", "1011", "1013", "1015", "1017", "1009", "1013"}
	bigAppIdT := make([]string, 0)

	for i, subAppId := range bigAppIds {
		if subAppId == "1011" {
			bigAppIdT = append(bigAppIds[:i], bigAppIds[i+1:]...)
		}
	}
	fmt.Println(bigAppIdT)
}

func test3() {
	var s = []string{"ca", "ab", "ec", "ec", "ca", "ab", "ec", "ca", "ab", "ec"}
	index := 0
	endIndex := len(s) - 1
	var result = make([]string, 0)
	for k, v := range s {
		if v == "ec" {
			result = append(result, s[index:k]...)
			index = k + 1
		} else if k == endIndex {
			result = append(result, s[index:endIndex+1]...)
		}
	}
	fmt.Println(result)
}
