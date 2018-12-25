package Socket

import (
	"encoding/binary"
	"fmt"
	"net"
)

var Host string = "0.0.0.0:8889"
func GetConn() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", Host)
	return
}

func SendMessage(conn net.Conn, message string){
	buf := make([]byte, 0, 0)
	buf = []byte(message)
	//先发送输入消息的长度
	inputLen := len(buf)
	inputLenByte := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(inputLenByte, uint32(inputLen))
	_, err := conn.Write(inputLenByte)
	if err != nil{
		fmt.Println("发送消息长度失败, err:", err)
		return
	} else {
		//发送输入长度正常之后，发送数据
		conn.Write(buf)
	}
}

func GetMessage(conn net.Conn)([]byte){
	//1. 获取消息的长度
	buf := make([]byte, 4, 4)
	conn.Read(buf)
	bufLen := binary.BigEndian.Uint32(buf)
	//2. 接收消息
	newsBuf := make([]byte, bufLen, bufLen)
	conn.Read(newsBuf)
	return newsBuf
}