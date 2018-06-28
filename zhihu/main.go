package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/astaxie/beego/httplib"

	"github.com/PuerkitoBio/goquery"
)

var chanLog = make(chan []string, 20)
var wg = sync.WaitGroup{}
var globalUser = make(map[string]bool)

func main() {
	go recordLog()
	//st := 18551016
	//et := st + 1000
	//getQuestion(int64(st), int64(et))
	//getQuestionTopic(int64(st), int64(et))
	getUserFollowers("halo9pan")

	//getZhuanLanArticle("zhangxin1023")
	//getUserFollowZhuanLan("jixin")

	//	wg.Wait()
	//	close(chanLog)
	//	time.Sleep(5 * time.Second)
}

func getQuestionTopic(st, et int64) {
	limitCh := make(chan int, 8)
	errorTime := 0

	for i := st; i <= et; i++ {
		limitCh <- 1
		time.Sleep(400 * time.Millisecond)
		wg.Add(1)

		go func(id int64) {
			url := "https://www.zhihu.com/topic/" + strconv.FormatInt(id, 10)
			req := httplib.Get(url).SetTimeout(30*time.Second, 30*time.Second)
			req.Header("X-Forwarded-For", "118.144.12.12")
			rets, err := req.String()
			if err != nil {
				fmt.Print(err.Error())
			}

			reg := regexp.MustCompile(`true">(.*)知乎</t`)
			questionRet := reg.FindAllString(rets, -1)
			if len(questionRet) > 0 {
				endPos := len(questionRet[0])
				questionTitle := questionRet[0][6 : endPos-3]

				fmt.Println(id, "==>", questionTitle, url)
				logdata := fmt.Sprintf("%d ==> %s  %s \n", id, questionTitle, url)
				fileid := getFieldId(int64(id))
				chanLog <- []string{strconv.FormatFloat(fileid, 'f', 0, 64), "topic", logdata}
				errorTime = 0
			} else {
				errorTime++
				fmt.Println("失败 ==> ", id)
				authWarning(int64(errorTime))
			}
			<-limitCh
			wg.Done()
		}(i)
	}
}

func getQuestion(st, et int64) {
	limitCh := make(chan int, 8)
	errorTime := 0

	for i := st; i <= et; i++ {
		limitCh <- 1
		time.Sleep(400 * time.Millisecond)
		wg.Add(1)

		go func(id int64) {
			url := "https://www.zhihu.com/question/" + strconv.FormatInt(id, 10)
			req := httplib.Get(url).SetTimeout(30*time.Second, 30*time.Second)
			req.Header("X-Forwarded-For", "118.144.12.12")
			rets, err := req.String()
			if err != nil {
				fmt.Print(err.Error())
			}

			reg := regexp.MustCompile(`true">(.*)知乎</t`)
			questionRet := reg.FindAllString(rets, -1)
			if len(questionRet) > 0 {
				endPos := len(questionRet[0])
				questionTitle := questionRet[0][6 : endPos-3]

				fmt.Println(id, "==>", questionTitle, url)
				logdata := fmt.Sprintf("%d ==> %s  %s  \n", id, questionTitle, url)
				fileid := getFieldId(id)
				chanLog <- []string{strconv.FormatFloat(fileid, 'f', 0, 64), "question", logdata}
				errorTime = 0
			} else {
				errorTime++
				fmt.Println("失败 ==> ", id)
				authWarning(int64(errorTime))
			}
			<-limitCh
			wg.Done()
		}(i)
	}
}

func getUser(userId string) {
	url := "https://www.zhihu.com/people/" + userId
	req := httplib.Get(url).SetTimeout(30*time.Second, 30*time.Second)
	req.Header("X-Forwarded-For", "118.144.12.12")
	rets, err := req.String()
	if err != nil {
		fmt.Print(err.Error())
	}

	reg := regexp.MustCompile(`true">(.*)知乎</t`)
	questionRet := reg.FindAllString(rets, -1)
	if len(questionRet) > 0 {
		endPos := len(questionRet[0])
		questionTitle := questionRet[0][6 : endPos-3]

		fmt.Println(userId, "==>", questionTitle, url)
		logdata := fmt.Sprintf("%s ==> %s  %s  \n", userId, questionTitle, url)
		chanLog <- []string{"1", "user", logdata}
	} else {
		fmt.Println("失败 ==> ", userId)
	}
}

func getUserFollowers(userId string) {
	url := "https://www.zhihu.com/people/" + userId + "/following"
	fmt.Println(url)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	fowers := make(map[string]string)
	doc.Find("a.UserLink-link").Each(func(index int, sel *goquery.Selection) {
		name := sel.Text()
		href, _ := sel.Attr("href")
		//fmt.Println(href)

		href = strings.Replace(href, "www.zhihu.com/org", "", -1)
		href = strings.Replace(href, "www.zhihu.com/people", "", -1)
		userId := strings.Replace(href, "/", "", -1)

		if name != "" && userId != "" {
			if !globalUser[userId] {
				if _, ok := fowers[userId]; !ok {
					fowers[userId] = name
				}
			}
		}
	})

	fmt.Println(fowers)
	if len(fowers) > 0 {
		fmt.Println(fowers)
		for uid, name := range fowers {
			if globalUser[uid] {
				continue
			}

			globalUser[uid] = true
			userUrl := "https://www.zhihu.com/people/" + uid
			logdata := fmt.Sprintf("%s ==> %s  %s  \n", uid, name, userUrl)
			chanLog <- []string{"5", "user", logdata}

			time.Sleep(400 * time.Millisecond)
			getUserFollowers(uid)
		}
	} else {
		userUrl := "https://www.zhihu.com/people/" + userId
		fmt.Printf("失败 ==> %s   %s", userId, userUrl)
	}
}

func recordLog() {
	var logfile *os.File

	for logMsg := range chanLog {
		filename := "./" + logMsg[1] + "/zhihu.log." + logMsg[0]
		var fh *os.File

		if checkFileExist(filename) {
			fh, _ = os.OpenFile(filename, os.O_APPEND, 0666)
		} else {
			fh, _ = os.Create(filename)
		}

		logger := log.New(fh, "", 0)
		logger.Output(1, logMsg[2])
	}

	logfile.Close()
}

func checkFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func getFieldId(id int64) float64 {
	for i := 11; i > 3; i-- {
		if float64(id) >= math.Pow10(i) {
			return math.Floor(float64(id)/math.Pow10(i)) * math.Pow10(i)
		}
	}
	return math.Pow10(4)
}

func authWarning(errorTime int64) {
	if errorTime >= 20 {
		fmt.Printf("Auth Warning %d Times", errorTime)
	}
}

func getZhuanLanArticle(zlid string) {
	url := "https://zhuanlan.zhihu.com/" + zlid
	fmt.Println(url)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	//fowers := make(map[string]string)
	//	doc.Find("h3.ArticleItem-Title").Each(func(index int, sel *goquery.Selection) {
	//		name := sel.Text()
	//		href, _ := sel.Attr("href")
	//		//fmt.Println(href)

	//		href = strings.Replace(href, "www.zhihu.com/org", "", -1)
	//		href = strings.Replace(href, "www.zhihu.com/people", "", -1)
	//		userId := strings.Replace(href, "/", "", -1)

	//		if name != "" && userId != "" {
	//			if _, ok := fowers[userId]; !ok {
	//				fowers[userId] = name
	//			}
	//		}
	//	})

	fmt.Println(doc.Text())
	doc.Find("li.ArticleItem.a").Each(func(index int, sel *goquery.Selection) {
		name := sel.Text()
		href, _ := sel.Attr("href")
		//fmt.Println(href)

		//		href = strings.Replace(href, "www.zhihu.com/org", "", -1)
		//		href = strings.Replace(href, "www.zhihu.com/people", "", -1)
		//		userId := strings.Replace(href, "/", "", -1)

		fmt.Println(name, href)
	})
}

func getUserFollowZhuanLan(userId string) {
	url := "https://www.zhihu.com/people/" + userId + "/following/columns"
	fmt.Println(url)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(doc.Text())
	fowers := make(map[string]string)

	//div#Popover-4859-13800-toggle
	doc.Find("a.ColumnLink").Each(func(index int, sel *goquery.Selection) {
		name := sel.Find("div#Popover-4859-13800-toggle").Text()
		href, _ := sel.Attr("href")
		//fmt.Println(href)

		href = strings.Replace(href, "www.zhihu.com/org", "", -1)
		href = strings.Replace(href, "www.zhihu.com/people", "", -1)
		userId := strings.Replace(href, "/", "", -1)

		if name != "" && userId != "" {
			if _, ok := fowers[userId]; !ok {
				fowers[userId] = name
			}
		}
	})
}
