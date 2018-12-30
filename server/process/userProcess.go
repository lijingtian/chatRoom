package process

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"chatRoom/server/model"
	"fmt"
	"net"
	"strconv"
)

type UserProcess struct {
	Conn net.Conn
	Mes Message.Message
	//该prcess归属的用户信息
	//UserInfo *Message.UserStatusNotify
	UserName string
	UserID int
}

func NewUserProcess(conn net.Conn, mes Message.Message) (*UserProcess){
	return &UserProcess{
		Conn:conn,
		Mes:mes,
		//UserInfo:&Message.UserStatusNotify{},
	}
}

/*
 * 验证用户登录时的账号密码是否正确
*/
func(this *UserProcess) CheckLogin(){
	//解析登录信息
	tempLoginModel, err := Message.MesDecode([]string{Message.LoginMesType, this.Mes.Data})
	if err != nil{
		fmt.Println("userProcess 33 err:", err)
		return
	}
	loginModel, ok := tempLoginModel.(*Message.LoginMes)
	if !ok{
		fmt.Println("user process 38 transfer to loginModel fail")
		return
	}

	//处理登录
	userModel := model.NewUserModel("", "")
	ok, _ = userModel.CheckLogin(loginModel.UserName, loginModel.UserPwd)
	var code, errStr string
	if ok{
		code = "200"
		errStr = ""

		userInfo, err := model.GetUserInfoByNameOnRedis(loginModel.UserName)
		if err != nil{
			fmt.Println("NotifyUserOnlineToClient 1 err:", err)
			return
		}

		this.UserID = userInfo.UserID
		this.UserName = userInfo.UserName
		userMgr.AddOnlineUser(this)
		go this.NotifyUserOnlineToClient()
	} else {
		code = "500"
		errStr = "user name or password err"
	}
	data, err := Message.MesEncode([]string{Message.LoginResMesType, code, errStr})
	//封装登录返回的消息
	if err != nil{
		fmt.Println("server process marshal login resmessage err:", err)
		return
	}
	//封装返回的message
	data, err = Message.MesEncode([]string{Message.MessageType, Message.LoginResMesType, data})
	if err != nil{
		fmt.Println("server process marshal resmessage err:", err)
		return
	}
	Socket.SendMessage(this.Conn, string(data))
}

/*
 * 用户注册
*/
func(this *UserProcess) Register(){
	data := this.Mes.Data
	decodeModel, err := Message.MesDecode([]string{Message.RegisterMesType, data})
	if err != nil{
		fmt.Println("user process register json unmarshal err:", err)
		return
	}
	registerMes, ok := decodeModel.(*Message.RegisterMes)
	if !ok{
		fmt.Println("user process 62 err:", err)
		return
	}
	userModel := model.NewUserModel(registerMes.UserName, registerMes.UserPwd)
	isOK, err := userModel.Register()


	var code, errStr string
	if isOK{
		code = "200"
		errStr = ""
	} else {
		code = "500"
		errStr = "register fail"
	}
	//封装登录返回的消息
	resData, err := Message.MesEncode([]string{Message.LoginResMesType, code, errStr})
	if err != nil{
		fmt.Println("server process marshal login resmessage err:", err)
		return
	}
	//封装返回的message
	resData, err = Message.MesEncode([]string{Message.MessageType, Message.RegisterMesType, resData})

	if err != nil{
		fmt.Println("server process marshal resmessage err:", err)
		return
	}
	Socket.SendMessage(this.Conn, string(resData))
}

/*
 * 将当前用户登录成功的消息通知到客户端
*/

func (this *UserProcess) NotifyUserOnlineToClient(){
	userInfo, err := model.GetUserInfoByNameOnRedis(this.UserName)
	if err != nil{
		fmt.Println("NotifyUserOnlineToClient 1 err:", err)
		return
	}
	//封装UserStatusNotifyType
	args := []string{
		Message.UserStatusNotifyType,
		strconv.Itoa(userInfo.UserID),
		userInfo.UserName,
		strconv.Itoa(Message.UserStatusNotifyOnlineStatus),
	}
	data, err := Message.MesEncode(args)
	if err != nil{
		fmt.Println("NotifyUserOnlineToClient 2 err:", err)
		return
	}

	//封装ServerNotify
	data, err = Message.MesEncode([]string{Message.ServerNotifyType, Message.UserStatusNotifyType, data})
	if err != nil{
		if err != nil{
			fmt.Println("NotifyUserOnlineToClient 3 err:", err)
			return
		}
	}
	//封装MessageType
	data, err = Message.MesEncode([]string{Message.MessageType, Message.ServerNotifyType, data})
	if err != nil{
		if err != nil{
			fmt.Println("NotifyUserOnlineToClient 3 err:", err)
			return
		}
	}
	fmt.Println(userMgr.GetAllOnlineUser())
	for uid, up := range userMgr.GetAllOnlineUser(){
		fmt.Println(uid)
		fmt.Println(this.UserID)
		//不给自己发
		if uid != this.UserID{
			Socket.SendMessage(up.Conn, data)
		}
	}
}