package main

import (
	"fmt"
	"time"
)

func main() {
	str_time := "2016-10-07 23:40:41"
	fmt.Println(StrToLocalTime(str_time))

	end_time := "2016-10-07 22:40:41"
	fmt.Println(StrToLocalTime(end_time))

	fmt.Println(DiffUnxiTime(str_time, end_time))

	t1 := StrToFormateDate(str_time)
	t2 := StrToFormateDate(end_time)
	fmt.Println(t1.Sub(t2).String())
}

func StrToLocalTime(StrDateTime string) int64 {
	set_location, _ := time.LoadLocation("Local")
	my_date_formate := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(my_date_formate, StrDateTime, set_location)
	return t.Unix()
}

func StrToUtcTime(StrDateTime string) int64 {
	my_date_formate := "2006-01-02 15:04:05"
	t, _ := time.Parse(my_date_formate, StrDateTime)
	return t.Unix()
}

func DiffUnxiTime(StartDateTime, EndDateTime string) int64 {
	return StrToLocalTime(EndDateTime) - StrToLocalTime(StartDateTime)
}

func StrToFormateDate(StrDateTime string) time.Time {
	set_location, _ := time.LoadLocation("Local")
	my_date_formate := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(my_date_formate, StrDateTime, set_location)

	//return t.Format(my_date_formate)
	return t
}
