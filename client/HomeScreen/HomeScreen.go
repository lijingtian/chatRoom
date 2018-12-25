package HomeScreen

import (
	"chatRoom/client/Login"
	"fmt"
)


func DrawHomeScreen(){
	var key int
	var loop bool  = true
	for loop{
		fmt.Println("-------------欢迎登陆多人聊天系统-------------")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key == 1{
		//登录聊天室
		LoginChatRoom()
	} else if key == 2{
		//注册用户
	}
}

func LoginChatRoom(){
	var userName string
	var userPwd string
	fmt.Println("请输入用户名:")
	fmt.Scanf("%s \n", &userName)
	fmt.Println("请输入密码")
	fmt.Scanf("%s \n", &userPwd)
	loginResMes, err := Login.CheckLogin(userName, userPwd)
	if err != nil {
		fmt.Println("登录失败", err)
	} else if loginResMes.Code != 200 {
		fmt.Println(loginResMes.Error)
	} else if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	}
}