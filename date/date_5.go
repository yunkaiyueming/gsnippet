package main

import (
	"fmt"
	//"strconv"
	"time"
)

func main() {
	today := time.Now().Format("2006-01-02")
	updateDates := getUpdateDates(today)
	fmt.Println(updateDates)

	updateDates = getUpdateDates2(today)
	fmt.Println(updateDates)

	fmt.Println("------------------------")
	OriginalDate(updateDates, getLtvDay())
	fmt.Println("------------------------")
	OriginalDate2(updateDates, getLtvDay())
}

func getUpdateDates(date string) []string {
	updateDates := make([]string, 0)
	days := getLtvDay()
	for _, day := range days {
		dateTime, _ := time.Parse("2006-01-02", date)
		//subTime := time.Hour * time.Duration(-24*day)
		//updateDate := dateTime.Add(subTime).Format("2006-01-02")
		updateDate := dateTime.AddDate(0, 0, -1*int(day)).Format("2006-01-02")
		updateDates = append(updateDates, updateDate)
	}
	return updateDates
}

func getUpdateDates2(date string) []string {
	updateDates := make([]string, 0)
	days := getLtvDay()
	for _, day := range days {
		dateTime, _ := time.Parse("2006-01-02", date)
		subTime := time.Hour * time.Duration(-24*day)
		updateDate := dateTime.Add(subTime).Format("2006-01-02")
		//updateDate := dateTime.AddDate(0, 0, -1*int(day)).Format("2006-01-02")
		updateDates = append(updateDates, updateDate)
	}
	return updateDates
}

func getLtvDay() []int64 {
	return []int64{0, 1, 2, 3, 4, 5, 6, 14, 29, 59, 89, 119, 149, 179, 239, 299, 359, 419, 479, 539, 599, 659, 719}
}

func OriginalDate(updateDates []string, ltvDay []int64) {
	for key, updateDate := range updateDates {
		updateDateTime, _ := time.Parse("2006-01-02", updateDate)
		calDate := updateDateTime.AddDate(0, 0, int(ltvDay[key]))
		fmt.Println(calDate)
	}
}

func OriginalDate2(updateDates []string, ltvDay []int64) {
	for key, updateDate := range updateDates {
		updateDateTime, _ := time.Parse("2006-01-02", updateDate)
		calDate := updateDateTime.Add(time.Duration(24*ltvDay[key]) * time.Hour)
		fmt.Println(calDate)
	}
}
