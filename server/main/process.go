package main

//func Process(conn net.Conn){
//	defer conn.Close()
//	mes := Socket.GetMessage(conn)
//	var socketMessage Message.Message
//	err := json.Unmarshal(mes, &socketMessage)
//	if err != nil{
//		fmt.Println("get message unmarshal err:", err)
//		return
//	}
//
//
//
//	switch socketMessage.Type {
//		case Message.LoginMesType:
//			//登录
//			//userModel = new()
//		case Message.RegisterMesType:
//			//注册
//			//registerProcess := process.ProcessFactory[Message.RegisterMesType]
//			registerProcess := process.NewUserProcess(conn, socketMessage)
//			registerProcess.Register()
//	}
//}