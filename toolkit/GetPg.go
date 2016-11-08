/*获取项目下的所有文件里的import的包*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	_ "regexp"
	"strings"
	"sync"
)

var ProjectPath string
var ImportPg []string
var GoFiles []string
var isDone chan int
var lock = sync.RWMutex{}

func init() {
	ProjectPath = "E:/GO_PATH/src/beego_action"
}

func main() {
	GetFilesByDir(ProjectPath)

	isDone = make(chan int) //必须要初始化
	fmt.Println(len(GoFiles))
	for _, filePath := range GoFiles {
		go GetFileImport(filePath)
	}

	doneNum := 1
Loop:
	for {
		select {
		case <-isDone:
			fmt.Println("done: ", doneNum)
			if doneNum == len(GoFiles) {
				break Loop
			}
			doneNum++
		}
	}

	fmt.Println(ImportPg)
	fmt.Println("regexp finish")
}

func GetFilesByDir(dirPath string) {
	rFile, _ := os.Open(dirPath)
	fileInfos, _ := rFile.Readdir(-1)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			GetFilesByDir(dirPath + "/" + fileInfo.Name())
		} else {
			if strings.Contains(fileInfo.Name(), ".go") {
				GoFiles = append(GoFiles, dirPath+"/"+fileInfo.Name())
			}
		}
	}
}

func GetFileCon(filePath string) string {
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func GetFileImport(fielPath string) {
	GetFileCon(fielPath)
	pgs := []string{"a", "b", "c", "z"}
	SafeSliceSet(pgs)
	isDone <- 1
}

func SafeSliceSet(addPgs []string) {
	lock.Lock()
	defer lock.Unlock()

	for _, addPg := range addPgs {
		if !inSlice(addPg, ImportPg) {
			ImportPg = append(ImportPg, addPg)
		}
	}
}

func inSlice(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
