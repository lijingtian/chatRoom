package Login

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"chatRoom/client/model"
	"chatRoom/client/process"
	"encoding/json"
	"fmt"
	"net"
)

//检查用户登录是否正确
func CheckLogin(userName string, userPwd string){
	conn, err := Socket.GetConn()
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
	Socket.SendMessage(conn, dataStr)

	//获取server返回的验证登录信息
	socketMes := Socket.GetMessage(conn)
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
		go WaitServerNotify(conn)
		ChatRoomScreen()
	}
	return
}

//检查用户登录是否正确
func Register(userName string, userPwd string) (loginResMes Message.LoginResMes, err error){
	conn, err := Socket.GetConn()
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
	Socket.SendMessage(conn, string(data))

	//获取server返回的验证登录信息
	socketMes := Socket.GetMessage(conn)
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
	err = json.Unmarshal([]byte(socketMesModel.Data), &loginResMes)
	if err != nil{
		fmt.Println("register get server message 2 unmarshal err:", err)
		return
	}
	return
}

/*
 *  登录成功之后，建立一条C到S的连接，等待S发送通知
*/
func WaitServerNotify(conn net.Conn){
	fmt.Println("调用了WaitServerNotify()")
		//conn, _ := Socket.GetConn()
		for{
			notifyByte := Socket.GetMessage(conn)
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
			fmt.Println(mesModel.Type)
			if mesModel.Type != Message.ServerNotifyType{
				fmt.Println("接收来自服务器的消息类型异常")
				break
			}
			//解析ServerNotifyType
			tempNotifyModel,err := Message.MesDecode([]string{Message.ServerNotifyType, mesModel.Data})
			if err != nil{
				fmt.Println("客户端接收到来自服务器的消息解析异常")
				break
			}
			notifyModel, ok := tempNotifyModel.(*Message.ServerNotify)
			fmt.Println(notifyModel.Type)
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
				model.UserList[userStatusModel.UserID] = userStatusModel
				fmt.Println(userStatusModel.UserName + " 上线了")
			default:
				fmt.Println("未定义的服务器通知类型")
			}
		}
}

func ChatRoomScreen(){
	//go WaitServerNotify()
	var key int
	var loop bool = true
	for loop{
		fmt.Println("----------欢迎来到聊天室-------------")
		fmt.Println("---1. 显示用户在线列表---")
		fmt.Println("---2. 发送消息---")
		fmt.Println("---3. 信息列表---")
		fmt.Println("---4. 退出系统---")
		fmt.Println("请选择（1-4）")
		fmt.Scanf("%d\n", &key)
		chatRoomProcess := process.NewChatRoomProcess()
		go chatRoomProcess.Process()
		if key == 1{

		} else if key == 2{

		} else if key == 3{

		} else if key == 4{
			loop = false
		}
	}
}