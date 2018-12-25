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
	//按照协议封装数据
	var loginMes Message.LoginMes
	loginMes.UserName = userName
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("user check login marshal err:", err)
		return
	}
	var socketMessage Message.Message
	socketMessage.Type = Message.LoginMesType
	socketMessage.Data = string(data)

	data, err = json.Marshal(socketMessage)
	Socket.SendMessage(conn, string(data))
	//获取server返回的验证登录信息
	socketMes := Socket.GetMessage(conn)
	//解封为Message类型
	var mes Message.Message
	err = json.Unmarshal(socketMes, &mes)
	if err != nil{
		fmt.Println("check login get server message unmarshal err:", err)
		return
	}
	if mes.Type != Message.LoginResMesType{
		fmt.Println("server message type err")
		return
	}
	err = json.Unmarshal([]byte(mes.Data), &loginRes)
	if err != nil{
		fmt.Println("check login get server message unmarshal err:", err)
		return
	}
	return
}