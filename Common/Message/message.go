package Message

import (
	"encoding/json"
	"fmt"
)

const(
	MessageType = "message"
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

//服务器客户端通用消息体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func NewMessage()(*Message){
	return &Message{}
}

func(this *Message) ModelInit(args []string){
	this.Type = args[0]
	this.Data = args[1]
}

func(this *Message) Encode() (mes string, err error){
	data, err := json.Marshal(this)
	mes = string(data)
	return
}

func(this *Message) DeCode(data string)(err error){
	err = json.Unmarshal([]byte(data), this)
	if err != nil{
		fmt.Println("message decode err:", err)
		return 
	}
	return nil
}

//用户登录聊天室时，连接服务器消息体
type CToSMes struct {
	UserID int `json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}