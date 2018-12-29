package util

import (
	"fmt"
	"time"
)

func GetNowTime()(time.Time, error){
	local, err := time.LoadLocation("Local")
	if err != nil{
		fmt.Println("utilFunc.go 11 err:", err)
	}
	var timeNow time.Time = time.Now().In(local)
	return timeNow, nil
}