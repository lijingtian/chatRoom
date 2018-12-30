package model

import "chatRoom/Common/Message"

var UserList map[int]*Message.UserStatusNotify

func init(){
	UserList = make(map[int]*Message.UserStatusNotify, 10)
}