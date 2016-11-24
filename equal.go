package main

import (
	"reflect"
)

func main() {
	TestEqual()
}

func TestEqual() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"a": 1, "c": 3, "b": 2}

	println(compareMap(map1, map2))
	println(reflect.DeepEqual(map1, map2))

	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{2, 3, 4, 1}
	println(compareSlice(slice1, slice2))
	println(reflect.DeepEqual(slice1, slice2))
}

func compareMap(map1, map2 map[string]int) bool {
	for k1, v1 := range map1 {
		if v2, has := map2[k1]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	for k2, v2 := range map2 {
		if v1, has := map1[k2]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func compareSlice(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
