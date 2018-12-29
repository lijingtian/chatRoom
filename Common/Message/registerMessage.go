package Message

import (
	"encoding/json"
	"fmt"
)

//注册用户消息体
type RegisterMes struct {
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

func NewRegisterMes()(*RegisterMes){
	return &RegisterMes{}
}

func(this *RegisterMes) ModelInit(args []string){
	this.UserName = args[0]
	this.UserPwd = args[1]
}

func(this *RegisterMes) Encode() (mes string, err error){
	data, err := json.Marshal(this)
	if err != nil{
		fmt.Println("login message 34 err:", err)
		return
	}
	mes = string(data)
	return
}

func(this *RegisterMes) DeCode(data string)(err error){
	err = json.Unmarshal([]byte(data), this)
	if err != nil{
		fmt.Println("message decode err:", err)
		return
	}
	return nil
}
