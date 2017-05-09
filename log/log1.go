package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	EchoConst()
	EchoDefaultStdInfo()
	RecordMyLog("现在开会了")
}

func EchoConst() {
	fmt.Println(log.Ldate, log.Ltime, log.Lmicroseconds, log.Llongfile, log.Lshortfile, log.LUTC, log.LstdFlags)
	fmt.Println(log.Ldate | log.Ltime)                     //按位或运算
	fmt.Println(log.Ldate | log.Ltime | log.Lmicroseconds) //按位或运算
	// 1 01
	// 2 10
	// 4 100
	// 8 1000
	// 16 10000
	// 32 100000
}

func EchoDefaultStdInfo() {
	fmt.Println(log.Flags())
	fmt.Println(log.Prefix())
	log.Println("it default value")
	log.Output(3, "输出三级信息")
}

func RecordMyLog(msg string) {
	logfile, err := os.OpenFile("E:/GO_PATH/src/gsnippet/log/output.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)
	logger.Output(3, "自定义调用log输出到日志:"+msg)
}
