/*
 * 处理用户登录和注册相关的业务逻辑
*/
package process

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"chatRoom/client/model"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}




//检查用户登录是否正确
func(this *UserProcess) CheckLogin(userName string, userPwd string){
	var err error
	this.Conn, err = Socket.GetConn()
	if err != nil{
		fmt.Println("tcp dial err:", err)
	}
	//按照协议封装登录数据
	data, err := Message.MesEncode([]string{Message.LoginMesType, "0", userName, userPwd})
	if err != nil{
		fmt.Println("user check login marshal err:", err)
		return
	}

	//按照协议封装Message消息
	dataStr, err := Message.MesEncode([]string{Message.MessageType, Message.LoginMesType, data})
	if err != nil{
		fmt.Println("client login go 32 err", err)
		return
	}
	Socket.SendMessage(this.Conn, dataStr)

	//获取server返回的验证登录信息
	socketMes := Socket.GetMessage(this.Conn)
	//解析为Message类型
	mes, err := Message.MesDecode([]string{Message.MessageType, string(socketMes)})
	if err != nil{
		fmt.Println("check login get server message unmarshal err:", err)
		return
	}
	mesModel, isOk := mes.(*Message.Message)
	if !isOk{
		fmt.Println("login 47 tranfer to Message.Message Err")
		return
	}
	if mesModel.Type != Message.LoginResMesType{
		fmt.Println("server message type err")
		return
	}
	loginResMesModel, err := Message.MesDecode([]string{Message.LoginResMesType, mesModel.Data})
	loginModel, ok := loginResMesModel.(*Message.LoginResMes)
	if !ok{
		fmt.Println("login 54 err")
	}

	if err != nil {
		fmt.Println("登录失败", err)
	} else if loginModel.Code != 200 {
		fmt.Println(loginModel.Error)
	} else if loginModel.Code == 200 {
		fmt.Println("登录成功")
		go this.WaitServerNotify()
		this.SendAllOnlineUserReqToServer()
		DrawScreenProcess.ChatRoomScreen()
	}
	return
}

//注册用户
func(this *UserProcess) Register(userName string, userPwd string) (){
	var err error
	this.Conn, err = Socket.GetConn()
	if err != nil{
		fmt.Println("tcp dial err:", err)
	}
	//按照协议封装数据
	data, err := Message.MesEncode([]string{Message.RegisterMesType, userName, userPwd})
	if err != nil{
		fmt.Println("user check login marshal err:", err)
		return
	}

	data, err = Message.MesEncode([]string{Message.MessageType, Message.RegisterMesType, data})
	Socket.SendMessage(this.Conn, string(data))

	//获取server返回的验证登录信息
	socketMes := Socket.GetMessage(this.Conn)
	//解封为Message类型
	socketMessage, err := Message.MesDecode([]string{Message.MessageType, string(socketMes)})
	if err != nil{
		fmt.Println("register get server message 1 unmarshal err:", err)
		return
	}
	socketMesModel, isOK := socketMessage.(*Message.Message)
	if !isOK{
		fmt.Println("login 91 err:", err)
	}
	if socketMesModel.Type != Message.RegisterMesType{
		fmt.Println("server message type err")
		return
	}
	loginResMes := new(Message.LoginResMes)
	err = json.Unmarshal([]byte(socketMesModel.Data), &loginResMes)
	if err != nil{
		fmt.Println("register get server message 2 unmarshal err:", err)
		return
	}
	if err != nil {
		fmt.Println("注册失败", err)
	} else if loginResMes.Code != 200 {
		fmt.Println(loginResMes.Error)
	} else if loginResMes.Code == 200 {
		fmt.Println("注册成功")
		DrawScreenProcess.DrawHomeScreen()
	}
	return
}

/*
 *  登录成功之后，建立一条C到S的连接，等待S发送通知
*/
func (this *UserProcess) WaitServerNotify(){
	for{
		notifyByte := Socket.GetMessage(this.Conn)
		//解析MessageType
		mes, err := Message.MesDecode([]string{Message.MessageType, string(notifyByte)})
		if err != nil{
			fmt.Println("home screen 124 err:", err)
			break
		}
		mesModel, ok := mes.(*Message.Message)
		if !ok{
			fmt.Println("home screen 129 err:", err)
			break
		}
		if mesModel.Type == Message.GroupMesType {
			SmsProcessModel.GetGroupMes(mesModel.Data)
			continue
		}
		if mesModel.Type != Message.ServerNotifyType{
			fmt.Println("接收来自服务器的消息类型异常")
			continue
		}
		//解析ServerNotifyType
		tempNotifyModel,err := Message.MesDecode([]string{Message.ServerNotifyType, mesModel.Data})
		if err != nil{
			fmt.Println("客户端接收到来自服务器的消息解析异常")
			break
		}
		notifyModel, ok := tempNotifyModel.(*Message.ServerNotify)
		switch notifyModel.Type {
		case Message.UserStatusNotifyType:
			//解析UserStatusNotifyType
			tempUserStatus, err := Message.MesDecode([]string{Message.UserStatusNotifyType, notifyModel.Content})
			if err != nil{
				fmt.Println("home screen err:", err)
				break
			}
			userStatusModel, ok := tempUserStatus.(*Message.UserStatusNotify)
			if !ok{
				fmt.Println("home screen tarfer to userStatusModel fail")
				break
			}
			//这里如果userStatusModel使用指针类型，则用户列表中全部变成最后一次的用户信息，具体原因，待查。
			model.UserListModel.AddUserList(*userStatusModel)
			fmt.Println(userStatusModel.UserName + " 上线了")
		case Message.GetAllOnlineUserResType:
			this.GetAllOnlineUserFromServer(notifyModel)
		default:
			fmt.Println("未定义的服务器通知类型")
		}
	}
}

/*
 *  登录成功之后，向服务器申请一次全部的在线用户信息列表
*/
func (this *UserProcess) SendAllOnlineUserReqToServer(){
	//按照协议封装Message消息
	dataStr, err := Message.MesEncode([]string{Message.MessageType, Message.GetAllOnlineUserType, ""})
	if err != nil{
		fmt.Println(err)
		return
	}
	Socket.SendMessage(this.Conn, dataStr)
}


/*
 *  接收来自服务器的全部的在线用户信息列表
*/
func(this *UserProcess) GetAllOnlineUserFromServer(notifyModel *Message.ServerNotify){
	//解析UserStatusNotifyType
	tempUserStatusArr := make([]string, 10)
	err := json.Unmarshal([]byte(notifyModel.Content), &tempUserStatusArr)
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, dataStr := range tempUserStatusArr{
		if dataStr == ""{
			continue
		}
		var userStatusModel Message.UserStatusNotify
		err = json.Unmarshal([]byte(dataStr), &userStatusModel)
		if err != nil{
			fmt.Println(err)
			continue
		}
		model.UserListModel.AddUserList(userStatusModel)
	}
}