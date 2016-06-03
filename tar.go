package main

import (
	"archive/tar"
	//"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sourceDirPath := "E:/www2/GitHub/go_code/src/beego_code"
	targetName := "test.tar"
	TarPathFile(sourceDirPath, targetName)

	//tarFile(sourceDirPath, targetName)
}

//TarPathFile 遍历给出的文件目录的所有文件，压缩为指定文件名
func TarPathFile(sourceDirPath, targetName string) {
	// file write
	os.OpenFile()

	fw, err := os.Create(sourceDirPath + "/" + targetName)
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	// tar write
	tw := tar.NewWriter(fw)
	defer tw.Close()

	// 打开文件夹
	dir, err := os.Open(sourceDirPath)
	if err != nil {
		panic(nil)
	}
	defer dir.Close()

	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	// 遍历文件列表
	for _, fi := range fis {
		// 逃过文件夹, 我这里就不递归了
		if fi.IsDir() || strings.Contains(fi.Name(), "tar") {
			continue
		}

		// 打印文件名称
		fmt.Println(fi.Name())

		// 打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		// 信息头
		h := new(tar.Header)
		h.Name = fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()

		// 写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}

		// 写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}
	}
}

???压缩指定文件
func tarFile(sourceFilePath, TargetFilePath string) {
	//dirname(sourceFilePath)

	fw, err := os.Create(TargetFilePath)
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	// tar write
	tw := tar.NewWriter(fw)
	defer tw.Close()

	// 打开文件夹
	fileInfo, err := os.Open(sourceFilePath)
	if err != nil {
		panic(nil)
	}
	defer fileInfo.Close()

	fmt.Println(fileInfo.Name())

	// 信息头
	h := new(tar.Header)
	h.Name = fileInfo.Name()
	//h.Size = fileInfo.Size()
	//h.Mode = int64(fileInfo.Mode())
	//h.ModTime = fileInfo.ModTime()

	// 写信息头
	err = tw.WriteHeader(h)
	if err != nil {
		panic(err)
	}

	// 写文件
	_, err = io.Copy(tw, fileInfo)
	if err != nil {
		panic(err)
	}
}
