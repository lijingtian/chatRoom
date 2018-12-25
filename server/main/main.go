package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main(){
	listenHandle, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil{
		fmt.Println("listen err:", err)
	}
	defer listenHandle.Close()
	for{
		conn, err := listenHandle.Accept()
		if err != nil{
			fmt.Println("accept err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	//1. 获取消息的长度
	buf := make([]byte, 4, 4)
	conn.Read(buf)
	bufLen := binary.BigEndian.Uint32(buf)
	//2. 接收消息
	newsBuf := make([]byte, bufLen, bufLen)
	conn.Read(newsBuf)
	fmt.Println(string(newsBuf))
}