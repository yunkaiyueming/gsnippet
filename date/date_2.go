package main

import "fmt"
import "time"
import "strconv"

func main() {
	fmt.Println(GetDateMap("2017-04-02"))
}

func StrToBeijingTime(strDateTime string) time.Time {
	setLocation, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", strDateTime, setLocation)
	return t
}

func GetDateMap(date string) map[string]string {
	Today := StrToBeijingTime(date)
	yesterday := Today.AddDate(0, 0, -1)
	yesterdayBeforeDay := Today.AddDate(0, 0, -2)
	firstDayThisMonth := GetMonthFirstDay(yesterday)

	yesterdayLastMonth := Today.AddDate(0, -1, -1)
	firstDayLastMonth := firstDayThisMonth.AddDate(0, -1, 0)
	lastDayLastMonth := firstDayThisMonth.AddDate(0, 0, -1)

	return map[string]string{
		"yesterday":            yesterday.Format("2006-01-02"),
		"yesterday_before_day": yesterdayBeforeDay.Format("2006-01-02"),
		"first_day_this_month": firstDayThisMonth.Format("2006-01-02"),
		"yesterday_last_month": yesterdayLastMonth.Format("2006-01-02"),
		"first_day_last_month": firstDayLastMonth.Format("2006-01-02"),
		"last_day_last_month":  lastDayLastMonth.Format("2006-01-02"),
	}
}

func GetMonthFirstDay(today time.Time) time.Time {
	year, month, _ := today.Date()
	firstDayStr := strconv.Itoa(year) + "-" + fmt.Sprintf("%02d", month) + "-01"
	return StrToBeijingTime(firstDayStr)
}
