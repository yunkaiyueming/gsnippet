package main

import (
	"gsnippet/pod_event"
)

func main() {
	//pod := pod_event.NewPodEvent() //外界不需要new
	log1 := pod_event.LogField{Timestamp: "123456", Depid: 12, Item: "POD", Event: pod_event.POD_START, Content: "aaaa", SN: "sdedesd"}
	log2 := pod_event.LogField{Timestamp: "123456567", Depid: 12, Item: "POD", Event: pod_event.POD_START, Content: "aaaa", SN: "sdedesd"}
	data := []pod_event.LogField{log1, log2}

	pod_event.OperationLogAdd(pod_event.PodEventObj, data) //如果用OperationLogAdd2，连pod_event.PodEventObj外界都不关心
	pod_event.OperationLogAdd2(data)
}
