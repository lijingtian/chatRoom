package main

import (
	"bufio"
	"fmt"
	"encoding/binary"
	"net"
	"os"
)

func main()  {
	//连接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil{
		fmt.Println("dial tcp err:", err)
	}
	defer conn.Close()
	//从控制台输入
	var inPutString string
	inputHandle := bufio.NewReader(os.Stdin)
	inPutString, err = inputHandle.ReadString('\n')
	if err != nil{
		fmt.Println("input string err:", err)
	}

	buf := make([]byte, 0, 0)
	buf = []byte(inPutString)
	//先发送输入消息的长度
	inputLen := len(buf)
	inputLenByte := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(inputLenByte, uint32(inputLen))
	fmt.Println(inputLenByte)
	_, err = conn.Write(inputLenByte)
	if err != nil{
		fmt.Println("发送消息长度失败, err:", err)
		return
	} else {
		//发送输入长度正常之后，发送数据
		conn.Write(buf)
	}
}
