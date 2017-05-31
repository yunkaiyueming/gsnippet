package main

import (
	"fmt"
	"time"
)

const (
	FORMAT_YYYY_MM             = "2006-01"
	FORMAT_YYYY_MM_DD          = "2006-01-02"
	FORMAT_YYYY_MM_DD_HH_II_SS = "2006-01-02 15:04:05"
	Y                          = "2006"
	M                          = "01"
	D                          = "02"
	H                          = "15"
	I                          = "04"
	S                          = "05"
)

func main() {
	p := time.Now().Format(FORMAT_YYYY_MM)
	fmt.Println(p)

	p = time.Now().Format(FORMAT_YYYY_MM_DD)
	fmt.Println(p)

	p = time.Now().Format(FORMAT_YYYY_MM_DD_HH_II_SS)
	fmt.Println(p)
}
