package main

import (
	"fmt"
	"log"
	"strconv"
	_ "strings"

	_ "time"

	"github.com/PuerkitoBio/goquery"
)

var PageNum = 20
var isDone = make(chan int, PageNum)

func main() {
	for i := 1; i <= PageNum; i++ {
		clawPageUrl := ParsePageUrl(i)
		go StartClaw(clawPageUrl, i)

		//time.Sleep(time.Second * 1)
		fmt.Println(i)
	}

	for i := 1; i <= PageNum; i++ {
		<-isDone
		fmt.Println(i, "Done")
	}
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

func StartClawGolang(clawUrl string, CurrentPage int) {
	doc, err := goquery.NewDocument(clawUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========= page: ", CurrentPage, "=========")
	doc.Find("h3.searchItemTitle").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		url, _ := s.Find("a").Attr("href")
		fmt.Printf("Review %d: %s - %s \n", i, text, url)
	})

	isDone <- 1
}

func StartClaw(clawUrl string, CurrentPage int) {
	doc, err := goquery.NewDocument(clawUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========= page: ", CurrentPage, "=========")
	// Find the review items
	doc.Find("a.titlelnk").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		text := s.Text()
		href, _ := s.Attr("href")
		fmt.Printf("Review %d: %s - %s\n", i, text, href)
	})

	isDone <- 1
}
