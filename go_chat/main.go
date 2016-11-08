// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/net/websocket"
)

const AddChatType = 1
const SendMsgType = 2
const LeaveChatType = 3

type User struct {
	Name string `json:"name"`
}

type ReceiveMsg struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type ResponseMsg struct {
	RespCode int    `json:"resp_code"`
	Action   int    `json:"action"`
	Users    []User `json:"users"`
	Uname    string `json:"name"`
	Msg      string `json:"msg"`
}

type WsCon struct {
	Con *websocket.Conn
}

var UserList = make([]User, 0, 50)
var WsPool = make([]WsCon, 0, 50)

func Router() {
	http.HandleFunc("/", RenderTemplate)
	http.Handle("/chat", websocket.Handler(EchoServer))
}

func ChatHandler() {
	Router()
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func EchoServer(ws *websocket.Conn) {
	var err error
	for {
		var receiveMsg string
		fmt.Println("payloadtype:", ws.PayloadType)
		//		if ws.PayloadType == 8 {
		//			RemovePool(WsCon{ws})
		//			ws.Close()
		//		}

		if err = websocket.Message.Receive(ws, &receiveMsg); err != nil {
			fmt.Println("Can't receive, reason:", err)
			RemovePool(WsCon{ws})
			ws.Close()
		}
		fmt.Println("client: ", receiveMsg)

		MsgResp := MsgReturn(receiveMsg, ws)
		respJson, _ := json.Marshal(MsgResp)
		fmt.Println("server:", respJson)

		//如果加入失败，只给该连接发送失败消息
		if MsgResp.RespCode == 0 {
			websocket.Message.Send(ws, string(respJson))
			return
		}

		for _, WsCon := range WsPool {
			if err = websocket.Message.Send(WsCon.Con, string(respJson)); err != nil {
				fmt.Println("Can't send:", err)
				RemovePool(WsCon)
				WsCon.Con.Close()
			}
		}
	}
}

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./chat.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func MsgReturn(msgStr string, ws *websocket.Conn) ResponseMsg {
	receiveMsg := &ReceiveMsg{}
	json.Unmarshal([]byte(msgStr), receiveMsg)
	fmt.Println(receiveMsg)

	respMsg := ResponseMsg{}
	switch receiveMsg.Type {
	case AddChatType:
		respMsg = HandleAddUser(receiveMsg.Name, ws)
		break

	case SendMsgType:
		respMsg = HandleSendMsg(receiveMsg.Name, receiveMsg.Msg)

	case LeaveChatType:
		respMsg = HandleLeaveMsg(receiveMsg.Name, ws)
	}

	fmt.Println(respMsg)
	return respMsg
}

func HandleAddUser(name string, ws *websocket.Conn) ResponseMsg {
	Resp := ResponseMsg{}
	for _, user := range UserList {
		if user.Name == name {
			Resp.RespCode = 0
			Resp.Action = 1
			Resp.Uname = name
			Resp.Msg = "该名字已被占用"
			return Resp
		}
	}

	UserList = append(UserList, User{Name: name})
	if !InPool(WsCon{ws}) {
		WsPool = append(WsPool, WsCon{ws})
	}
	fmt.Println(UserList)
	Resp = ResponseMsg{
		RespCode: 1,
		Action:   AddChatType,
		Users:    UserList,
		Uname:    name,
		Msg:      "加入成功",
	}

	return Resp
}

func HandleSendMsg(name string, msg string) ResponseMsg {
	Resp := ResponseMsg{
		RespCode: 1,
		Action:   SendMsgType,
		Users:    UserList,
		Uname:    name,
		Msg:      msg,
	}
	fmt.Println(Resp)
	return Resp
}

func HandleLeaveMsg(name string, ws *websocket.Conn) ResponseMsg {
	DeleteUser(name)
	ws.Close()
	RemovePool(WsCon{ws})
	Resp := ResponseMsg{
		RespCode: 1,
		Action:   LeaveChatType,
		Users:    UserList,
		Uname:    name,
	}
	fmt.Println(Resp)
	return Resp
}

func DeleteUser(name string) {
	var key int
	for i, u := range UserList {
		if u.Name == name {
			key = i
			break
		}
	}
	UserList = append(UserList[:key], UserList[key+1:]...)
}

func InPool(wsCon WsCon) bool {
	for _, con := range WsPool {
		if con == wsCon {
			return true
		}
	}
	return false
}

func RemovePool(leaveCon WsCon) {
	var key int
	for i, con := range WsPool {
		if con == leaveCon {
			key = i
		}
	}
	WsPool = append(WsPool[:key], WsPool[key+1:]...)
}

func main() {
	ChatHandler()
}
