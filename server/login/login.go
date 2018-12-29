package login

import (
	"chatRoom/Common/Message"
	"encoding/json"
	"fmt"
)

func GetLoginMessage(mes string) (loginMes Message.LoginMes, err error) {
	err = json.Unmarshal([]byte(mes), &loginMes)
	if err != nil {
		fmt.Println("get login message unmarshal err:", err)
		return
	}
	return loginMes, nil
}