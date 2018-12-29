package main

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"chatRoom/Common/db"
	"chatRoom/server/model"
	"chatRoom/server/process"
	"encoding/json"
	"fmt"
	"net"
)

func main(){
	go UserMysqlToRedis()
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

func UserMysqlToRedis(){
	rows, err := db.MysqlDBPool.Query("SELECT id,user_name,user_pwd FROM user")
	if err != nil{
		fmt.Println("main 56 err:", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		userModel := model.NewUserModel("", "")
		rows.Scan(&userModel.UserID, &userModel.UserName, &userModel.UserPwd)
		userInfo, err := json.Marshal(userModel)
		if err != nil{
			fmt.Println("main 66 err", err)
			continue
		}
		db.NewRedisModel().Conn.Do("Hset", "userInfo", userModel.UserName, string(userInfo))
	}
}