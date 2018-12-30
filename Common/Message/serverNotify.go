package Message

import (
	"encoding/json"
	"fmt"
)

/*
 * 服务器通知消息结构体
 * NType 1- 用户状态变更通知
*/
type ServerNotify struct {
	Type string `json:"type"`
	Content string `json:"content"`
}

func NewServerNotify()(*ServerNotify){
	return &ServerNotify{}
}

func(this *ServerNotify) ModelInit(args []string){
	this.Type = args[0]
	this.Content = args[1]
}

func(this *ServerNotify) Encode() (mes string, err error){
	data, err := json.Marshal(this)
	mes = string(data)
	return
}

func(this *ServerNotify) DeCode(data string)(err error){
	err = json.Unmarshal([]byte(data), this)
	if err != nil{
		fmt.Println("message decode err:", err)
		return
	}
	return nil
}