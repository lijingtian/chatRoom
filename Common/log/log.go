package log

import (
	"chatRoom/Common/util"
	"fmt"
	"github.com/logrus"
	"os"
)


var CommonLog = logrus.New()

func init(){
	CommonLog.Formatter = &logrus.JSONFormatter{
		TimestampFormat: util.YMDHIS,
	}
	rootPath, err := util.GetRootDir()
	fmt.Println(rootPath)
	if err != nil{
		CommonLog.Warn("日志文件创建时根目录获取失败")
		return
	}
	logPath := rootPath + "/log"
	fmt.Println(logPath)
	err = util.CreateDir(logPath)
	if err != nil {
		fmt.Println(err)
	}
	logHandle, err := os.OpenFile(logPath + "/runtime.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		CommonLog.Warn("日志文件创建失败")
		return
	}
	CommonLog.Out = logHandle
	CommonLog.Level = logrus.InfoLevel
}
