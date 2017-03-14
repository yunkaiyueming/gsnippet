package main

import (
	"fmt"
	"strconv"
	"time"
)

func StrToBeijingTime(strDateTime string) time.Time {
	setLocation, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", strDateTime, setLocation)
	return t
}

func getMonthPeriods(startMonth, endMonth string) []map[string]string {
	var monthDatePeriods []map[string]string
	today := time.Now()
	thisMonth := time.Now().Format("2006-01")

	if today.Day() == 1 { //月首不显示数据
		thisMonth = time.Now().AddDate(0, 0, -1).Format("2006-01")
	}
	if thisMonth < endMonth {
		endMonth = thisMonth
	}

	for month := startMonth; month <= endMonth; {
		monthStartDate := StrToBeijingTime(month + "-01")
		monthStartDateStr := monthStartDate.Format("2006-01-02")
		monthEndDateStr := monthStartDate.AddDate(0, 1, -1).Format("2006-01-02")
		days := monthStartDate.AddDate(0, 1, -1).Day()

		monthDatePeriods = append(monthDatePeriods, map[string]string{"start_date": monthStartDateStr, "end_date": monthEndDateStr, "month": month, "days": strconv.Itoa(days)})
		month = monthStartDate.AddDate(0, 1, 0).Format("2006-01") //下一个月
	}

	return monthDatePeriods
}

func main() {
	fmt.Println(getMonthPeriods("2017-01", "2017-05"))
}
