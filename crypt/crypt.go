package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	MyMd5()
	Test()
}

func MyMd5() {
	hash_type := md5.New()
	hash_type.Write([]byte("123456")) // 需要加密的字符串为 sharejs.com

	res1 := hex.EncodeToString(hash_type.Sum(nil))
	fmt.Println(res1)

	//第二种方法MD5加密
	data := []byte("123456")
	fmt.Printf("%x", md5.Sum(data))
}

func Test() {
	//import "crypto/sha256"
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("%x", h.Sum(nil))

	//import "crypto/sha1"
	h = sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("%x", h.Sum(nil))

	//import "crypto/md5"
	h = md5.New()
	io.WriteString(h, "需要加密的密码")
	fmt.Printf("%x", h.Sum(nil))
}
