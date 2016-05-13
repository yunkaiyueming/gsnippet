package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dir_path := "E:/www2/GitHub/go_code/src"
	file_infos, _ := ioutil.ReadDir(dir_path)

	//	Name() string       // base name of the file
	//	Size() int64        // length in bytes for regular files; system-dependent for others
	//	Mode() FileMode     // file mode bits
	//	ModTime() time.Time // modification time
	//	IsDir() bool        // abbreviation for Mode().IsDir()
	//	Sys() interface{}

	for i, file_info := range file_infos {
		fmt.Println(i, file_info.Name(), file_info.Size(), file_info.Mode(), file_info.IsDir())
	}

	log.Fatal("log fatal")
}
