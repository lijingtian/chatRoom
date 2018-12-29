package Message

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//服务器应答消息体
type LoginResMes struct {
	Code int `json:"code"`	//500用户未注册， 200登录成功
	Error string `json:"error"`
}

func NewLoginResMes()(*LoginResMes){
	return &LoginResMes{}
}

func(this *LoginResMes) ModelInit(args []string){
	if code, err := strconv.Atoi(args[0]); err != nil{
		fmt.Println("login message 22 err:", err)
		this.Code = 0
	} else {
		this.Code = code
	}
	this.Error = args[1]
}

func(this *LoginResMes) Encode() (mes string, err error){
	data, err := json.Marshal(this)
	if err != nil{
		fmt.Println("login message 34 err:", err)
		return
	}
	mes = string(data)
	return
}

func(this *LoginResMes) DeCode(data string)(err error){
	err = json.Unmarshal([]byte(data), this)
	if err != nil{
		fmt.Println("message decode err:", err)
		return
	}
	return nil
}