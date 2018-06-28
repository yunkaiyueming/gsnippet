package main

import (
	"fmt"
	"os"

	"github.com/chenqinghe/baidu-ai-go-sdk/version/ocr"
)

const (
	//	APIKEY    = "5RijeBzVjQ82uPx8gxGGfeNXlfRt7yH6"
	//	APISECRET = "keiyq3oKrkYsSPUcrf0gtRKneeTxjuqV"

	APIKEY    = "yxhhb5oxZBIj5tYxfiS3q2HK"
	APISECRET = "fh4KOy0MTkD0BNBnvprTqwWLGmi4P14q"
)

func main() {
	client := ocr.NewOCRClient(APIKEY, APISECRET)

	f, err := os.OpenFile("q1.png", os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	rs, err := client.GeneralRecognizeBasic(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(rs))
}
