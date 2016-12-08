package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	date := "2016-01-15"
	fmt.Println(GetDateMap(date))
}

func GetDateMap(date string) map[string]string {
	//	Today, _ := time.Parse("2016-01-02", date)
	//	fmt.Println(Today)

	Today := StrToLocalTime(date)
	fmt.Println(Today)

	yesterday := Today.AddDate(0, 0, -1)
	fmt.Println(yesterday)

	yesterday_before_day := Today.AddDate(0, 0, -2)
	first_day_this_month, _ := GetMonthFirstDay(yesterday_before_day)
	fmt.Println(first_day_this_month)

	yesterday_last_month := Today.AddDate(0, -1, 0)
	first_day_last_month := first_day_this_month.AddDate(0, -1, 0)
	last_day_last_month := first_day_this_month.AddDate(0, 0, -1)

	return map[string]string{
		"yesterday":            yesterday.Format("2006-01-02"),
		"yesterday_before_day": yesterday_before_day.Format("2006-01-02"),
		"first_day_this_month": first_day_this_month.Format("2006-01-02"),
		"yesterday_last_month": yesterday_last_month.Format("2006-01-02"),
		"first_day_last_month": first_day_last_month.Format("2006-01-02"),
		"last_day_last_month":  last_day_last_month.Format("2006-01-02"),
	}
}

func GetMonthFirstDay(today time.Time) (time.Time, error) {
	year, month, _ := today.Date()
	monthInt := fmt.Sprintf("%d", month)
	firstDayStr := ""

	if month < 10 {
		firstDayStr = strconv.Itoa(year) + "-0" + string(monthInt) + "-01"
	} else {
		firstDayStr = strconv.Itoa(year) + "-" + monthInt + "-01"
	}
	return StrToLocalTime(firstDayStr), nil
}

func StrToLocalTime(StrDateTime string) time.Time {
	set_location, _ := time.LoadLocation("Asia/Shanghai")
	my_date_formate := "2006-01-02"
	t, _ := time.ParseInLocation(my_date_formate, StrDateTime, set_location)
	return t
}
