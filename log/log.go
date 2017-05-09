package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func main() {
	GetHttpService()
	SavePprofShot()
}

func Mylog() {
	log.Println("123")

	fmt.Println(log.Ldate)
	fmt.Println(log.Ltime)
	fmt.Println(log.Lmicroseconds)
	fmt.Println(log.Llongfile)
	fmt.Println(log.Lshortfile)
	fmt.Println(log.LstdFlags)
}

func user_list(w http.ResponseWriter, req *http.Request) {
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log_msg := "hello"
	io.WriteString(w, log_msg)
	RecordLog(log_msg)
}

func RecordLog(logMsg string) {
	logfile, err := os.OpenFile("E:/GO_PATH/src/go_code/test.log", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}

	defer logfile.Close()

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	logger.Println(logMsg)
	logger.Println("oh....")
}

func GetHttpService() {
	http.HandleFunc("/user_list", user_list)
	http.HandleFunc("/hello", HelloServer)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListendAndServe: ", err)
	}
}

func SavePprofShot() {
	f, err := os.OpenFile("E:/GO_PATH/src/go_code/cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	pprof.StopCPUProfile()
	f.Close()
}
