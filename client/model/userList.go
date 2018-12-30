package model

import (
	"chatRoom/Common/Message"
)

type UserList struct {
	userInfoList map[int]Message.UserStatusNotify
}
var UserListModel *UserList

func init(){
	UserListModel = &UserList{
		userInfoList: make(map[int]Message.UserStatusNotify, 10),
	}
}

func (this *UserList) GetUserList()(map[int]Message.UserStatusNotify){
	return this.userInfoList
}

func (this *UserList) AddUserList(user Message.UserStatusNotify){
	this.userInfoList[user.UserID] = user
}
