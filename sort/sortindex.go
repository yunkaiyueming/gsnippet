package main

import (
	"fmt"
)

func main() {
	mapData := map[string]int64{
		"bi":    2,
		"gpm":   3,
		"other": 0,
		"tm":    1,
		"so":    0,
	}
	sortIndex := make([]string, len(mapData))
	var masterVal, j int64 = 0, 0
	//count := len(mapData)

	for i := 0; i < len(mapData); i++ {
		masterVal = 0
		for i2, v2 := range mapData {
			if masterVal <= v2 {
				sortIndex[j] = i2
				masterVal = v2
			}
		}
		fmt.Println(sortIndex[j])
		delete(mapData, sortIndex[j])
		j++
	}

	fmt.Println(sortIndex)
}
