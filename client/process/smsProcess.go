package process

import (
	"chatRoom/Common/Message"
	"chatRoom/Common/Socket"
	"fmt"
	"net"
)
//定义一个全局的聊天信息链接，防止发送一次消息连接服务器一次，造成连接浪费
var SmsProcessModel *SmsProcess

type SmsProcess struct {
	Conn net.Conn
}

func init(){
	SmsProcessModel = new(SmsProcess)
	var err error
	SmsProcessModel.Conn, err = Socket.GetConn()
	if err != nil{
		fmt.Println(err)
		return
	}
}

/*
 * 发送群聊消息
*/
func(this *SmsProcess) SendGroupMes(mes string){
	//封装消息为 GroupMes
	data, err := Message.MesEncode([]string{Message.GroupMesType, mes})
	if err != nil{
		fmt.Println(err)
		return
	}

	//封装消息为 Message
	data, err = Message.MesEncode([]string{Message.MessageType, Message.GroupMesType, data})
	if err != nil{
		fmt.Println(err)
		return
	}
	Socket.SendMessage(this.Conn, data)
}

func(this *SmsProcess)GetGroupMes(mes string){
	//解析 Message
	mesTemp, err := Message.MesDecode([]string{Message.GroupMesType, mes})
	if err != nil{
		fmt.Println(err)
		return
	}
	groupMes, ok := mesTemp.(*Message.GroupMes)
	if !ok{
		fmt.Println("group message 类型断言失败")
		return
	}
	//打印群发的消息
	fmt.Println(groupMes.Content)
}