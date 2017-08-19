package main

import (
	"fmt"
	"time"
)

func main() {
	getPayUserLtvDate()
}

func test(){
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().AddDate(0, 0, -1).Unix() - time.Now().Unix())
}


func  getPayUserLtvDate(regtime int64) map[string]string {
	genLtvDateEndSec :=func(startDate time.Time ,diffDay int, ltvStart int64) int64{
		ltvDateStr:=startDate.AddDate(0,0, diffDay).Format("2006-01-02")
		ltvDate,_:=time.Parse("2006-01-02", ltvDateStr)
		return ltvDate.Unix()+oneDaySec-1
	}
	
	oneDaySec := 24*60*60-1	
	ltvDates :=make(map[string]string)
	
	regDateStr:=time.Unix(regtime, 0).Format("2006-01-02")
	regDate,_:=time.Parse("2006-01-02", regDateStr)
	
	yesterdayStr := time.Now().AddDate(0,0,-1).Format("2006-01-02")
	lastWeekStr:= time.Now().AddDate(0,0,-6).Format("2006-01-02")
	ltv3DayStr:=regDate.AddDate(0,0,2).Format("2006-01-02")
	ltv7DayStr :=regDate.AddDate(0,0,6).Format("2006-01-02")
	ltv15DayStr :=regDate.AddDate(0,0,14).Format("2006-01-02")
	ltv30DayStr :=regDate.AddDate(0,0,29).Format("2006-01-02")
	ltv60DayStr :=regDate.AddDate(0,0,59).Format("2006-01-02")
	ltv90DayStr :=regDate.AddDate(0,0,89).Format("2006-01-02")
	ltv150DayStr :=regDate.AddDate(0,0,179).Format("2006-01-02")
	
	
	yesterday,_:=time.Parse("2006-01-02", yesterdayStr)
	lastWeek,_:=time.Parse("2006-01-02", lastWeekStr)
	ltv3Day :=time.Parse("2006-01-02",ltv3DayStr )
	ltv7Day :=time.Parse("2006-01-02",ltv7DayStr )
	ltv7Day :=time.Parse("2006-01-02",ltv7DayStr )
	ltv7Day :=time.Parse("2006-01-02",ltv7DayStr )
	ltv7Day :=time.Parse("2006-01-02",ltv7DayStr )
	ltv7Day :=time.Parse("2006-01-02",ltv7DayStr )
	
	ltvDates["total"] = fmt.Sprintf("%d;%d",regtime,yesterday.Unix()+oneDaySec-1)
	ltvDates["last_week"] = fmt.Sprintf("%d;%d",lastWeek.Unix(),yesterday.Unix()+oneDaySec-1)
	ltvDates["ltv_3"] := fmt.Sprintf("%d;%d",regtime,  ltv3Day.Unix()+oneDaySec-1)
	ltvDates["ltv_7"] := fmt.Sprintf("%d;%d",regtime,  ltv7Day.Unix()+oneDaySec-1)
	ltvDates["ltv_15"] := fmt.Sprintf("%d;%d",regtime,  ltv7Day.Unix()+oneDaySec-1)
	ltvDates["ltv_30"] := fmt.Sprintf("%d;%d",regtime,  ltv7Day.Unix()+oneDaySec-1)
	ltvDates["ltv_60"] := fmt.Sprintf("%d;%d",regtime,  ltv7Day.Unix()+oneDaySec-1)
	ltvDates["ltv_90"] := fmt.Sprintf("%d;%d",regtime,  ltv7Day.Unix()+oneDaySec-1)
	ltvDates["ltv_180"] := fmt.Sprintf("%d;%d",regtime,  ltv7Day.Unix()+oneDaySec-1)
}