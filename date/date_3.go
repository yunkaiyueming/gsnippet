package main

func getMonthDate(month, bigAppId string) map[string]string {
	monthDates := make(map[string]string)
	dateStr := month + "-01"
	firstDate := util.StrToBeijingTime(dateStr)
	endDate := firstDate.AddDate(0, 1, -1)

	if endDate.Unix() >= time.Now().Unix() {
		endDate = time.Now().AddDate(0, 0, -1)
	}

	monthDates["firstDate"] = firstDate.Format("2006-01-02")
	monthDates["endDate"] = endDate.Format("2006-01-02")

	//In返回采用loc指定的地点和时区，但指向同一时间点的Time， 这种操作unxi时间戳不会变，日期会变。
	timeZone, _ := time.LoadLocation(app.GetZoneLocation(bigAppId))
	firstDate = firstDate.In(timeZone)
	endDate = endDate.In(timeZone)

	monthDates["firstSeconds"] = strconv.FormatInt(firstDate.Unix(), 10)
	monthDates["endSeconds"] = strconv.FormatInt(endDate.Unix()+24*60*60-1, 10)
	return monthDates
}
