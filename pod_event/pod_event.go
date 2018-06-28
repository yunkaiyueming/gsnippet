package pod_event

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

type Event interface {
	AddEvent(log []LogField)
}

type podEvent struct {
	FileHandle *os.File
	mutex      *sync.Mutex
	lastday    int64 //增加一个标记，标记FileHandle是哪天的FileHandle，跨天了就要更新FileHandle，外界不需要关心
}

type depEvent struct {
}

type LogField struct {
	Timestamp string
	Depid     int64
	Item      string
	Event     EVENT_ITEM
	Content   string
	SN        string
}

var PodEventObj = newPodEvent() //这里将PodEventObj开放对外，连newPodEvent都不需要开放

func init() {
	t := time.Now().Local()
	filename := LOG_BASE_PATH + "/" + FILENAME + t.Format(F_DATE)

	if checkFileExist(filename) {
		PodEventObj.FileHandle, _ = os.OpenFile(filename, os.O_APPEND, 0666)
	} else {
		PodEventObj.FileHandle, _ = os.Create(filename)
	}
	PodEventObj.mutex = new(sync.Mutex)
	PodEventObj.lastday = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix() //给其第一次赋值
	print(PodEventObj.lastday)
}

func newPodEvent() *podEvent {
	return &podEvent{}
}

func (po *podEvent) AddEvent(data []LogField) {
	t := time.Now().Local()

	msgPack := make(map[string]interface{})
	var jsonStr []byte
	msg := ""

	for i := 0; i < len(data); i++ {
		msgPack["Timestamp"] = t.Format(F_DATETIME)
		msgPack["Depid"] = data[i].Depid
		msgPack["Item"] = data[i].Item
		msgPack["Event"] = data[i].Event
		msgPack["Content"] = data[i].Content
		msgPack["SN"] = data[i].SN

		jsonStr, _ = json.Marshal(msgPack)
		msg += string(jsonStr) + "\n"

	}

	filename := LOG_BASE_PATH + "/" + FILENAME + t.Format(F_DATE)

	po.mutex.Lock() //这样用po，外面传的po也就是PodEventObj
	writeLog(filename, msg)
	po.mutex.Unlock()
}

func (de *depEvent) AddEvent(log []LogField) {

}

func OperationLogAdd(p Event, log []LogField) {
	p.AddEvent(log)
}

// 如果OperationLogAdd改成下面这样，外面连PodEventObj都不需要关心，只需要传[]LogField
func OperationLogAdd2(log []LogField) {
	PodEventObj.AddEvent(log)
}

func writeLog(filename, message string) {
	t := time.Now().Local()
	todaySt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix() //今日0点时间戳
	if PodEventObj.lastday != todaySt {                                                 //跨天了就更新FileHandle，之后这一天都不会再更新FileHandle了，不会重复打开
		if checkFileExist(filename) {
			PodEventObj.FileHandle, _ = os.OpenFile(filename, os.O_APPEND, 0666)
		} else {
			PodEventObj.FileHandle, _ = os.Create(filename)
		}
		PodEventObj.lastday = todaySt
	}

	io.WriteString(PodEventObj.FileHandle, message)
}

func checkFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
