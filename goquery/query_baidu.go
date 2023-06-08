package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	_ "time"
)

// downloadImage 从网站爬取图片，放到本地目录中
func downloadImage(url string, index int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http get error:", err.Error())
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body error:", err.Error())
		return
	}

	filename := "./" + "image" + strconv.Itoa(index) + ".jpg"
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println("Write file error:", err.Error())
		return
	}
}

// getImagesURL 从网站爬取图片的URL地址
func getImagesURL(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http get error:", err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body error:", err.Error())
		os.Exit(1)
	}

	str := string(data)
	htmlSplit := strings.Split(str, "src=\"http")
	images := make([]string, 0, 10)

	for _, v := range htmlSplit {
		if strings.Contains(v, ".jpg") {
			images = append(images, "http"+v[:strings.Index(v, ".jpg")+4])
		}
	}

	return images
}

func main() {
	// 从网站爬取图片URL
	images := getImagesURL("http://image.baidu.com/")

	var wg sync.WaitGroup

	// 用并发多协程爬取图片
	for index, url := range images {
		wg.Add(1)
		go func(index int, url string) {
			defer wg.Done()

			// 下载图片
			downloadImage(url, index)

			fmt.Println("Downloaded image:", url, index)
		}(index, url)
	}

	wg.Wait()

	fmt.Println("Downloaded", len(images), "images in", "seconds")
}
