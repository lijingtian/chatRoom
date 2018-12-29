package Message

/*
 * 封装消息结构体为string
 * @param data [0] 消息类型 [1:] 结构体初始化参数
 * @return string 封装后的数据
 * @return error 过程中产生的异常
*/
func MesEncode(data []string)(string, error){
	mesModel := MessageFactory[data[0]]
	mesModel.ModelInit(data[1:])
	return mesModel.Encode()
}

/*
 * 解析消息结构体
 * @param data [0] 消息类型 [1] 待解析的数据
 * @return MessageInferface 解析数据附着的结构体
 * @return error 过程中产生的异常
*/
func MesDecode(data []string) (MessageInferface, error){
	mesModel := MessageFactory[data[0]]
	return mesModel, mesModel.DeCode(data[1])
}