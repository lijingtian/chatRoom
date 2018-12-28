package process

import (
	"chatRoom/Common/Socket"
	"fmt"
	"net"
	"time"
)

type ChatRoomProcess struct {
	Conn net.Conn
}

func NewChatRoomProcess()(*ChatRoomProcess){
	conn, err := Socket.GetConn()
	if err != nil{
		fmt.Println("chatRoomProcess get socket conn err:", err)
		return nil
	}
	return &ChatRoomProcess{
		Conn: conn,
	}
}

func (this *ChatRoomProcess) Process(){
	defer this.Conn.Close()
	go this.WaitServerMes()
	fmt.Println("This is chat room process process()")
	time.Sleep(88888)
}

func (this *ChatRoomProcess) WaitServerMes(){
	mes := Socket.GetMessage(this.Conn)
	fmt.Println(mes)
}