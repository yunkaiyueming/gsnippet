package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//ListDirFiles("E:/GO_PATH/src/go_code")

	//Mkdir("E:/GO_PATH/src/go_code/ok")
	//DeleteDir("E:/GO_PATH/src/go_code/ok")

	//CreateFile("E:/GO_PATH/src/go_code/1test.log")
	//DeleteFile("E:/GO_PATH/src/go_code/1test.log")

	//GetFileInfo("E:/GO_PATH/src/go_code/append.go")
	GetUserPwd()
}

//读取指定目录下的文件及目录
func ListDirFiles(dirPath string) {
	file_infos, _ := ioutil.ReadDir(dirPath)
	for i, file_info := range file_infos {
		fmt.Println(i, file_info.Name(), file_info.Size(), file_info.Mode(), file_info.IsDir())
	}
}

func Mkdir(dirName string) {
	os.MkdirAll(dirName, 755)
}

func DeleteDir(dirPath string) {
	os.Remove(dirPath)
}

func CreateFile(filename string) {
	os.Create(filename)
}

func DeleteFile(filename string) {
	os.Remove(filename)
}

func GetFileInfo(filePath string) {
	info, _ := os.Stat(filePath)
	fmt.Println(info)
}

func GetUserPwd() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
}
