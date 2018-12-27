//用户类
package model

import (
	"chatRoom/Common/db"
	"encoding/json"
	"fmt"
)

type UserModel struct {
	UserID int `json:"userid"`
	UserName string `json:"username"`
	UserPwd string `json:"userpwd"`
}

func NewUserModel(username string, userpwd string)(*UserModel){
	return &UserModel{
		UserID:123,
		UserName:username,
		UserPwd:userpwd,
	}
}

func (this *UserModel) CheckLogin()(isOK bool, err error){
	redisModel := db.NewRedisModel()
	userInfo, err := json.Marshal(this)
	if err != nil{
		fmt.Println("user model json marshal err:", err)
		return
	}
	redisModel.Conn.Do("Hset", "userInfo", string(userInfo))

	return
}

func (this *UserModel) Register()(isOK bool, err error){
	redisModel := db.NewRedisModel()
	defer redisModel.Conn.Close()
	_, err = redisModel.Conn.Do("Hset", "userInfo", this.UserName, this.UserPwd)
	if err != nil{
		fmt.Println("user model register hset redis err:", err)
		return
	}
	return true, nil
}
