package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}

func main() {
	p := fmt.Println
	var num float64 = 884
	p(SendMonthExpected("2016-12-04", num))
	//	fmt.Println(Round(12.154453, 1))
	//	fmt.Println(Round(112322.5253, 2))
	//	fmt.Println(Round(1454.0453, 3))
	//	fmt.Println(Round(12.40253, 4))
}

func SendMonthExpected(dateStr string, originNum interface{}) float64 {
	date := StrToBeijingTime(dateStr)
	firstDayMonth := GetMonthFirstDay(date)
	lastDay := firstDayMonth.AddDate(0, 1, -1).Day()
	today := date.Day()

	originNumFloat, ok := originNum.(float64)
	if !ok {
		originNumInt, intOk := originNum.(int64)
		if intOk {
			originNumFloat = float64(originNumInt)
		}
	}

	fmt.Println(originNumFloat, today, lastDay)
	expectedNum := originNumFloat / float64(today) * float64(lastDay)
	return Round(expectedNum, 1)
}

func StrToBeijingTime(strDateTime string) time.Time {
	setLocation, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", strDateTime, setLocation)
	return t
}

func GetMonthFirstDay(today time.Time) time.Time {
	year, month, _ := today.Date()
	firstDayStr := strconv.Itoa(year) + "-" + fmt.Sprintf("%02d", month) + "-01"
	return StrToBeijingTime(firstDayStr)
}
