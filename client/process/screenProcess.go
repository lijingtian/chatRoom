package process

import (
	"chatRoom/client/model"
	"fmt"
)
var DrawScreenProcess *DrawScreen
type DrawScreen struct {}

func init(){
	DrawScreenProcess = new(DrawScreen)
}

func(this *DrawScreen) ChatRoomScreen(){
	var key int
	var loop bool = true
	for loop{
		fmt.Println("----------欢迎来到聊天室-------------")
		fmt.Println("---1. 显示用户在线列表---")
		fmt.Println("---2. 发送消息---")
		fmt.Println("---3. 信息列表---")
		fmt.Println("---4. 退出系统---")
		fmt.Println("请选择（1-4）")
		fmt.Scanf("%d\n", &key)
		if key == 1{
			userList := model.UserListModel.GetUserList()
			for _, model := range userList{
				fmt.Printf("%s | 在线\n", model.UserName)
			}
		} else if key == 2{
			fmt.Println("请输入要发送的消息：")
			var mes string
			fmt.Scanf("%s\n", &mes)
			SmsProcessModel.SendGroupMes(mes)
		} else if key == 3{

		} else if key == 4{
			loop = false
		}
	}
}

func(this *DrawScreen) DrawHomeScreen(){
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
		this.LoginChatRoom()
	} else if key == 2{
		//注册用户
		this.Register()
	} else if key == 3{
		return
	}
}

func(this *DrawScreen) LoginChatRoom(){
	var userName string
	var userPwd string
	fmt.Println("请输入用户名:")
	fmt.Scanf("%s \n", &userName)
	fmt.Println("请输入密码")
	fmt.Scanf("%s \n", &userPwd)
	userProcessModel := new(UserProcess)
	userProcessModel.CheckLogin(userName, userPwd)
}
func(this *DrawScreen) Register(){
	var userName string
	var userPwd string
	fmt.Println("请输入用户名:")
	fmt.Scanf("%s \n", &userName)
	fmt.Println("请输入密码")
	fmt.Scanf("%s \n", &userPwd)
	userProcessModel := new(UserProcess)
	userProcessModel.Register(userName, userPwd)
}