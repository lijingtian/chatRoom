package Message

const(
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserID int `json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

type LoginResMes struct {
	Code int `json:"code"`	//500用户未注册， 200登录成功
	Error error `json:"error"`
}