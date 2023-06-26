package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	_ "time"

	"github.com/PuerkitoBio/goquery"
	"github.com/qibin0506/Meizar/code/rule"
)

type Rule interface {
	UrlRule() (url string)
	PageRule(currentPage int) (page string)
	ImageRule(doc *goquery.Document, f func(image string))
}

type JandanRule struct{}

func (p *JandanRule) UrlRule() (url string) {
	return "http://jandan.net/ooxx/"
}

func (p *JandanRule) PageRule(currentPage int) (page string) {
	return "page-" + strconv.Itoa(currentPage)
}

func (p *JandanRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}

//返回初始化的url配置
func RuleProvider() Rule {
	return &JandanRule{}
}

func New(dir string, startPage int, r rule.Rule, cookie string, client *http.Client, pageSort int) *Meizar {
	return &Meizar{dir: dir, currentPage: startPage, userCookie: cookie, r: r, client: client, pageSort: pageSort}
}

type Meizar struct {
	dir         string
	currentPage int
	userCookie  string
	client      *http.Client
	r           rule.Rule
	pageSort    int
}

var (
	maxRoutineNum = 20
	ch            = make(chan int, maxRoutineNum)
)

func (p *Meizar) Start() {
	if !p.isExist(p.dir) {
		if err := os.Mkdir(p.dir, 0777); err != nil {
			fmt.Println("can not mkdir " + p.dir)
			os.Exit(2)
		}
	}

	for p.currentPage > 0 {
		p.parsePage(p.r.UrlRule() + p.r.PageRule(p.currentPage))
		if p.pageSort == 1 {
			p.currentPage++
		} else {
			p.currentPage--
		}
		fmt.Println("======= page:", p.currentPage, " ========")
		//time.Sleep(time.Second * 1)
		<-ch
	}
}

//解析URL下载图片
func (p *Meizar) parsePage(url string) {
	req := p.buildRequest(url)
	resp, err := p.client.Do(req)

	if err != nil {
		fmt.Println("failed parse " + url)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(url + "-->" + strconv.Itoa(resp.StatusCode))
		return
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	img, err := p.parseImageUrl(bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range img {
		go p.downloadImage(item)
	}
}

//构建请求
func (p *Meizar) buildRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("request img error")
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36")
	req.Header.Set("Cookie", p.userCookie)
	return req
}

//使用goquery解析出图片url
func (p *Meizar) parseImageUrl(reader io.Reader) (res []string, err error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	p.r.ImageRule(doc, func(image string) {
		res = append(res, image)
	})

	return res, nil
}

//下载图片到本地
func (p *Meizar) downloadImage(url string) {
	fileName := p.getNameFromUrl(url)
	if p.isExist(p.dir + fileName) {
		fmt.Println("already download " + fileName)
		return
	}

	req := p.buildRequest(url)
	resp, err := p.client.Do(req)
	if err != nil {
		fmt.Println("failed download " + url)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("failed download " + url)
		return
	}

	defer func() {
		resp.Body.Close()
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("begin download " + fileName)
	localFile, _ := os.OpenFile(p.dir+fileName, os.O_CREATE|os.O_RDWR, 0777)
	if _, err := io.Copy(localFile, resp.Body); err != nil {
		fmt.Println("failed save " + fileName)
	}

	fmt.Println("success download " + fileName)
	ch <- 1
}

//文件是否已存在
func (p *Meizar) isExist(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

//根据url获取图片名字
func (p *Meizar) getNameFromUrl(url string) string {
	arr := strings.Split(url, "/")
	return arr[len(arr)-1]
}
