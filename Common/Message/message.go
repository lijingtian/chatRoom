package Message

const(
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

//服务器客户端通用消息体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//用户登录消息体
type LoginMes struct {
	UserID int `json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

//用户登录聊天室时，连接服务器消息体
type CToSMes struct {
	UserID int `json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

//注册用户消息体
type RegisterMes struct {
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

//服务器应答消息体
type LoginResMes struct {
	Code int `json:"code"`	//500用户未注册， 200登录成功
	Error string `json:"error"`
}