package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	m map[string]interface{}
	sync.Mutex
}

func (m *myMap) push(key string, e interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	if _, exist := m.m[key]; exist {
		return m.m
	}
	m.m[key] = e
	return m.m
}

func (m *myMap) pop(key string) interface{} {
	m.Lock()
	defer m.Unlock()
	if _, exist := m.m[key]; exist {
		m.m[key] = nil
		return m.m
	}
	return m.m
}

func newMap() *myMap {
	return &myMap{m: make(map[string]interface{})}
}

func main() {
	m := newMap()
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.pop("hello"))
	fmt.Println(m.pop("hello"))
}
