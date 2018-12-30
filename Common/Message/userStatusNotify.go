package Message

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const(
	UserStatusNotifyOnlineStatus = 1
)

/*
 * 用户状态变更消息体
 * UserStatus 1- 在线
*/
type UserStatusNotify struct {
	UserID int `json:"userid"`
	UserName string `json:"username"`
	UserStatus int `json:"userstatus"`
}

func NewUserStatusNotify()(*UserStatusNotify){
	return &UserStatusNotify{}
}

func(this *UserStatusNotify) ModelInit(args []string){
	if id, err := strconv.Atoi(args[0]); err != nil{
		fmt.Println("serverNotify string to int err:", err)
		return
	} else {
		this.UserID = id
	}
	this.UserName = args[1]
	if status, err := strconv.Atoi(args[2]); err != nil{
		fmt.Println("serverNotify string to int err:", err)
		return
	} else {
		this.UserStatus = status
	}
}

func(this *UserStatusNotify) Encode() (mes string, err error){
	data, err := json.Marshal(this)
	mes = string(data)
	return
}

func(this *UserStatusNotify) DeCode(data string)(err error){
	err = json.Unmarshal([]byte(data), this)
	if err != nil{
		fmt.Println("message decode err:", err)
		return
	}
	return nil
}


