package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var MapData = make(map[string]string)
var RWLock = sync.RWMutex{}

func Test() {
	for i := 1; i < 200; i++ {
		go AddKey(strconv.Itoa(i), strconv.Itoa(i))
		go GetKey(strconv.Itoa(i))
	}

	ticker := time.NewTicker(time.Second * 2)
	for {
		<-ticker.C
		break
	}
	fmt.Println("finish")
}

func AddKey(key, val string) {
	RWLock.Lock()
	defer RWLock.Unlock()
	MapData[key] = val
	fmt.Println(MapData)
}

func GetKey(key string) string {
	RWLock.RLock()
	defer RWLock.RUnlock()
	fmt.Println("read ok")
	return MapData[key]
}

func RemoveKey(key string) bool {
	if _, isExist := MapData[key]; isExist {
		return false
	} else {
		RWLock.Lock()
		defer RWLock.Unlock()

		delete(MapData, key)
		return true
	}
}

func UpdateKey(key, updateVal string) bool {
	RWLock.Lock()
	defer RWLock.Unlock()

	MapData[key] = updateVal
	return true
}
