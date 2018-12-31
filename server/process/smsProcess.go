package process

import (
	"chatRoom/Common/Socket"
	"net"
)

type SmsProcess struct {
	Conn net.Conn
	Message string
}

func NewSmsProcess(mes string)(*SmsProcess){
	return &SmsProcess{
		Message:mes,
	}
}

func (this *SmsProcess) TranferGroupMes(){
	for _, user := range userMgr.GetAllOnlineUser(){
		Socket.SendMessage(user.Conn, this.Message)
	}
}