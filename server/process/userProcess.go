package process

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"chatRoom/server/login"
	"chatRoom/server/model"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	Mes Message.Message
	//该prcess归属的用户名称
	UserID int
}

func NewUserProcess(conn net.Conn, mes Message.Message) (*UserProcess){
	return &UserProcess{
		Conn:conn,
		Mes:mes,
	}
}

func(this *UserProcess) CheckLogin(){
	loginMes, err := login.GetLoginMessage(this.Mes.Data)
	if err != nil{
		fmt.Println("server process get login message err:", err)
		return
	}

	//处理登录
	userModel := model.NewUserModel(loginMes.UserName, loginMes.UserPwd)
	isOK, _ := userModel.CheckLogin()
	var loginResMes Message.LoginResMes
	if isOK{
		loginResMes.Code = 200
		loginResMes.Error = ""
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "user name or password err"
	}
	//封装登录返回的消息
	data, err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("server process marshal login resmessage err:", err)
		return
	}
	//封装返回的message
	var mes Message.Message
	mes.Type = Message.LoginResMesType
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("server process marshal resmessage err:", err)
		return
	}
	Socket.SendMessage(this.Conn, string(data))
}

func(this *UserProcess) Register(){
	data := this.Mes.Data
	var registerMes Message.RegisterMes
	err := json.Unmarshal([]byte(data), &registerMes)
	if err != nil{
		fmt.Println("user process register json unmarshal err:", err)
		return
	}
	userModel := model.NewUserModel(registerMes.UserName, registerMes.UserPwd)
	isOK, err := userModel.Register()


	var loginResMes Message.LoginResMes
	if isOK{
		loginResMes.Code = 200
		loginResMes.Error = ""
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "register fail"
	}
	//封装登录返回的消息
	resData, err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("server process marshal login resmessage err:", err)
		return
	}
	//封装返回的message
	var mes Message.Message
	mes.Type = Message.RegisterMesType
	mes.Data = string(resData)
	resData, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("server process marshal resmessage err:", err)
		return
	}
	Socket.SendMessage(this.Conn, string(resData))
}