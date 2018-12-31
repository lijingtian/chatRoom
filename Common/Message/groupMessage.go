package Message

import (
	"encoding/json"
	"fmt"
)

type GroupMes struct {
	Content string `json:"content"`
}

func NewGroupMes()(*GroupMes){
	return &GroupMes{}
}

func(this *GroupMes) ModelInit(args []string){
	this.Content = args[0]
}

func(this *GroupMes) Encode() (mes string, err error){
	data, err := json.Marshal(this)
	mes = string(data)
	return
}

func(this *GroupMes) DeCode(data string)(err error){
	err = json.Unmarshal([]byte(data), this)
	if err != nil{
		fmt.Println("message decode err:", err)
		return
	}
	return nil
}