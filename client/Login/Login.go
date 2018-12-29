package Login

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"encoding/json"
	"fmt"
)

//检查用户登录是否正确
func CheckLogin(userName string, userPwd string) (loginRes Message.LoginResMes, err error){
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
	//解封为Message类型
	//var mes Message.Message
	//mes := Message.MessageFactory[Message.MessageType]
	//err = mes.DeCode(string(socketMes))
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
	&loginRes = loginModel
	//err = loginRes.DeCode(mesModel.Data)
	//if err != nil{
	//	fmt.Println("check login get server message unmarshal err:", err)
	//	return
	//}
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

	data, err = Message.MesEncode([]string{Message.MessageType, data})
	Socket.SendMessage(conn, string(data))

	//获取server返回的验证登录信息
	socketMes := Socket.GetMessage(conn)
	//解封为Message类型
	socketMessage, err := Message.MesDecode([]string{Message.MessageType, string(socketMes)})
	//err = socketMessage.DeCode(string(socketMes))
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