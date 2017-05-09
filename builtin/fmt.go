package main

import (
	"fmt"
	"time"
)

func main() {
	test_time()
}

func test_fmt() {
	test_b := "bbb"
	test_c := "ccc"
	int_d := 132

	fmt.Print("111", "aaa")
	num, _ := fmt.Printf("222 %s %s %d ", test_c, test_b, int_d)
	fmt.Println(num)
	fmt.Println("333")

	//fmt.Fprint("444")
	//fmt.Fprintf("555")
	//fmt.Fprintln("666")

	get_string := fmt.Sprintf("777 %s %d", test_b, int_d)
	fmt.Println(get_string)
	fmt.Sprintf("888")
	fmt.Sprintln("999")
}

func test_return() string {
	fmt.Println("test_return string")
	return "aa"
}

func test_time() {
	//	p := fmt.Println

	//now_time := time.Now()
	//	fmt.Println(now_time)

	//	zone_name, offset_time := now_time.Zone()
	//	fmt.Println(zone_name)
	//	p(offset_time)

	//	unix_time := now_time.Unix()
	//	fmt.Printf("%d", unix_time)
	//	//p(now_time.Clock())
	//	//p(now_time.Year())

	//	fmt.Println(now_time.String())

	formate_string := "2012-03-10 00:00:00"
	date_string := "2016-05-11 15:10:10"

	//	create_time := now_time.Format(formate_string)
	//	fmt.Println(create_time)

	t2, _ := time.Parse(formate_string, date_string)
	fmt.Println(t2)
}
