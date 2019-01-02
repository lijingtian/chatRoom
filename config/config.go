package config

import (
	"chatRoom/Common/util"
	"fmt"
)

const(
	YMD = "2006-01-02"
	YMDHIS = "2006-01-02 15:04:05"
)

var RootPath string
func init(){
	var err error
	RootPath, err = util.GetRootDir()
	if err != nil{
		fmt.Println(err)
		return
	}
}
