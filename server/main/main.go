package main

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"chatRoom/server/process"
	"encoding/json"
	"fmt"
	"net"
)

func main(){
	listenHandle, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil{
		fmt.Println("listen err:", err)
	}
	defer listenHandle.Close()
	for{
		conn, err := listenHandle.Accept()
		if err != nil{
			fmt.Println("accept err:", err)
			continue
		}
		go Process(conn)
	}
}

func Process(conn net.Conn){
	defer conn.Close()
	mes := Socket.GetMessage(conn)
	var socketMessage Message.Message
	err := json.Unmarshal(mes, &socketMessage)
	if err != nil{
		fmt.Println("get message unmarshal err:", err)
		return
	}

	switch socketMessage.Type {
	case Message.LoginMesType:
		//登录
		//userModel = new()
		loginProcess := process.NewUserProcess(conn, socketMessage)
		loginProcess.CheckLogin()
	case Message.RegisterMesType:
		//注册
		//registerProcess := process.ProcessFactory[Message.RegisterMesType]
		registerProcess := process.NewUserProcess(conn, socketMessage)
		registerProcess.Register()
	}
}

//
//func process(conn net.Conn) {
//	defer conn.Close()
//	mes := Socket.GetMessage(conn)
//	var socketMessage Message.Message
//	err := json.Unmarshal(mes, &socketMessage)
//	if err != nil{
//		fmt.Println("get message unmarshal err:", err)
//		return
//	}
//	switch socketMessage.Type{
//		case Message.LoginMesType :
//			loginMes, err := login.GetLoginMessage(socketMessage.Data)
//			if err != nil{
//				fmt.Println("server process get login message err:", err)
//				return
//			}
//			var loginResMes Message.LoginResMes
//			isOK := login.CheckLogin(loginMes)
//			if isOK{
//				loginResMes.Code = 200
//				loginResMes.Error = ""
//			} else {
//				loginResMes.Code = 500
//				loginResMes.Error = "user name or password err"
//			}
//			//封装登录返回的消息
//			data, err := json.Marshal(loginResMes)
//			if err != nil{
//				fmt.Println("server process marshal login resmessage err:", err)
//				return
//			}
//			//封装返回的message
//			var mes Message.Message
//			mes.Type = Message.LoginResMesType
//			mes.Data = string(data)
//			data, err = json.Marshal(mes)
//			if err != nil{
//				fmt.Println("server process marshal resmessage err:", err)
//				return
//			}
//			Socket.SendMessage(conn, string(data))
//	}
//}