package main

import (
	"fmt"
	"log"
	"strconv"
	_ "strings"

	"runtime"
	"time"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

var PageNum = 50
var JobGroup sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核实在CPU计算上更显优势
}

func main() {
	start_time := time.Now().Unix()
	for i := 1; i <= PageNum; i++ {
		JobGroup.Add(1)

		clawPageUrl := ParsePageUrl(i)
		go StartClaw(clawPageUrl, i)
		fmt.Println(i)
	}

	JobGroup.Wait()
	end_time := time.Now().Unix()
	fmt.Println("spend time: ", end_time-start_time)
	fmt.Println("All finish")
}

func ClawWebSite() string {
	return "http://www.cnblogs.com/sitehome"
}

func ClawGolangPageUrl(page int) string {
	return "http://zzk.cnblogs.com/s/blogpost?Keywords=golang&pageindex=" + strconv.Itoa(page)
}

func ParsePageUrl(page int) string {
	//http://www.cnblogs.com/sitehome/p/10
	return ClawWebSite() + "/p/" + strconv.Itoa(page)
}

func RecordRes() {

}

func StartClaw(clawUrl string, CurrentPage int) {
	doc, err := goquery.NewDocument(clawUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========= page: ", CurrentPage, "=========")
	doc.Find("a.titlelnk").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		text := s.Text()
		href, _ := s.Attr("href")
		fmt.Printf("Review %d: %s - %s\n", i, text, href)
	})

	JobGroup.Done()
}
