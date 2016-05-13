package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash_type := md5.New()
	hash_type.Write([]byte("123456")) // 需要加密的字符串为 sharejs.com

	res1 := hex.EncodeToString(hash_type.Sum(nil))
	fmt.Println(res1)

	//第二种方法MD5加密
	data := []byte("123456")
	fmt.Printf("%x", md5.Sum(data))
}
