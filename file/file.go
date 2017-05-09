package main

import (
	_ "bufio" //缓存IO
	_ "fmt"
	_ "io"
	"io/ioutil" //io 工具包
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
  from: http://www.isharey.com/?p=143
*/

func main() {
	var wireteString = "测试n"
	var filename = "E:/www2/GitHub/go_code/src/1.log"
	//var f *os.File
	//var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	//	if checkFileIsExist(filename) { //如果文件存在
	//		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	//		fmt.Println("文件存在")
	//	} else {
	//		f, err1 = os.Create(filename) //创建文件
	//		fmt.Println("文件不存在")
	//	}
	//	check(err1)
	//	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	//	check(err1)
	//	fmt.Printf("写入 %d 个字节n", n)

	/*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************/
	var d1 = []byte(wireteString)
	err2 := ioutil.WriteFile(filename, d1, 0666) //写入文件(字节数组)
	ioutil.WriteFile(filename, d1, 0666)         //写入文件(字节数组)
	ioutil.WriteFile(filename, d1, 0666)         //写入文件(字节数组)
	check(err2)

	/*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************/
	//	f, err3 := os.Create("./output3.txt") //创建文件
	//	check(err3)
	//	defer f.Close()
	//	n2, err3 := f.Write(d1) //写入文件(字节数组)
	//	check(err3)
	//	fmt.Printf("写入 %d 个字节n", n2)
	//	n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	//	fmt.Printf("写入 %d 个字节n", n3)
	//	f.Sync()

	//	/***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
	//	w := bufio.NewWriter(f) //创建新的 Writer 对象
	//	n4, err3 := w.WriteString("bufferedn")
	//	fmt.Printf("写入 %d 个字节n", n4)
	//	w.Flush()
	//	f.Close()
}

func (this *FileController) ListDir() {
	dirPath := "E:/GO_PATH/src/beego_action"
	rFile, _ := os.Open(dirPath)
	fileInfos, _ := rFile.Readdir(-1)

	this.Data["fileInfos"] = fileInfos
	this.MyRender("file/view_list_dir.html")
}

func ReadByOsRead(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func ReadByBufio(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func ReadByIoutil(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

// filecopy.go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    CopyFile("target.txt", "source.txt")
    fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}