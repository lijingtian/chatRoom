package Login

import (
	"chatRoom/client/Socket"
	"encoding/json"
	"fmt"
	"wechatRoom/common/message"
)

//检查用户登录是否正确
func CheckLogin(userName string, userPwd string) (err error){
	conn, err := Socket.GetConn()
	if err != nil{
		fmt.Println("tcp dial err:", err)
	}
	//按照协议封装数据
	var loginMes message.LoginMes
	loginMes.UserName = userName
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("user check login marshal err:", err)
		return
	}
	var socketMessage message.Message
	socketMessage.Type = message.LoginMesType
	socketMessage.Data = string(data)

	data, err = json.Marshal(socketMessage)
	Socket.SendMessage(conn, string(data))
	return nil
}