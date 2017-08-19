package main

import (
	"fmt"
	"time"
)

func main() {
	reg, _ := time.Parse("2006-01-02", "2014-02-01")
	fmt.Println("regtime time：", reg.Unix())
	s := getPayUserLtvDate(reg.Unix())
	fmt.Println(s)
}

func getPayUserLtvDate(regtime int64) map[string]string {
	oneDaySec := 24*60*60 - 1
	ltvDates := make(map[string]string)

	regDateStr := time.Unix(regtime, 0).Format("2006-01-02")
	regDate, _ := time.Parse("2006-01-02", regDateStr)
	lastWeekStr := time.Now().AddDate(0, 0, -6).Format("2006-01-02")
	lastWeek, _ := time.Parse("2006-01-02", lastWeekStr)

	genLtvDate := func(ltvStart int64, startDate time.Time, diffDay int) string {
		ltvDateStr := startDate.AddDate(0, 0, diffDay).Format("2006-01-02")
		ltvDate, _ := time.Parse("2006-01-02", ltvDateStr)

		return fmt.Sprintf("%d;%d", ltvStart, ltvDate.Unix()+int64(oneDaySec)-1)
	}

	ltvDates["total"] = genLtvDate(regtime, time.Now(), -1)
	ltvDates["last_week"] = genLtvDate(lastWeek.Unix(), time.Now(), -1)
	ltvDates["ltv_3"] = genLtvDate(regtime, regDate, 2)
	ltvDates["ltv_7"] = genLtvDate(regtime, regDate, 6)
	ltvDates["ltv_15"] = genLtvDate(regtime, regDate, 14)
	ltvDates["ltv_30"] = genLtvDate(regtime, regDate, 29)
	ltvDates["ltv_60"] = genLtvDate(regtime, regDate, 59)
	ltvDates["ltv_90"] = genLtvDate(regtime, regDate, 89)
	ltvDates["ltv_180"] = genLtvDate(regtime, regDate, 179)

	return ltvDates
}
