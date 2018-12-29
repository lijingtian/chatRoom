package Message

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//用户登录消息体
type LoginMes struct {
	UserID int `json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

func NewLoginMes()(*LoginMes){
	return &LoginMes{}
}

func(this *LoginMes) ModelInit(args []string){
	if id, err := strconv.Atoi(args[0]); err != nil{
		fmt.Println("login message 22 err:", err)
		this.UserID = 0
	} else {
		this.UserID = id
	}
	this.UserName = args[1]
	this.UserPwd = args[2]
}

func(this *LoginMes) Encode() (mes string, err error){
	data, err := json.Marshal(this)
	if err != nil{
		fmt.Println("login message 34 err:", err)
		return
	}
	mes = string(data)
	return
}

func(this *LoginMes) DeCode(data string)(err error){
	err = json.Unmarshal([]byte(data), this)
	if err != nil{
		fmt.Println("message decode err:", err)
		return
	}
	return nil
}