package Message

type MessageInferface interface {
	ModelInit([]string)
	Encode()(string, error)
	DeCode(string)(error)
}

var MessageFactory = make(map[string]MessageInferface, 10)

func init(){
	MessageFactory[MessageType] = NewMessage()
	MessageFactory[LoginMesType] = NewLoginMes()
	MessageFactory[LoginResMesType] = NewLoginResMes()
	MessageFactory[RegisterMesType] = NewRegisterMes()
}