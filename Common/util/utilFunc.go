package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var runOSType = runtime.GOOS

/**
 * 根据当前时区，获取当前时间
*/
func GetNowTime()(timeNow time.Time, err error){
	local, err := time.LoadLocation("Local")
	if err != nil{
		return
	}
	timeNow = time.Now().In(local)
	return
}

/**
 * 获取根地址的绝对路径
*/
func GetRootDir()(dir string, err error){
	dir, err = filepath.Abs(os.Args[0])
	if err != nil{
		fmt.Println(err)
	}
	mainDir, err := GetParentDir(dir)
	if err != nil{
		fmt.Println(err)
		return
	}
	serverDir, err := GetParentDir(mainDir)
	if err != nil{
		fmt.Println(err)
		return
	}
	return GetParentDir(serverDir)
}

/**
 * 获取传入路径的父路径
*/
func GetParentDir(path string)(dir string, err error){
	var dirFlag string = "/"
	if runOSType == "windows" {
		dirFlag = "\\"
	}
	return Substr(path, 0, strings.LastIndex(path, dirFlag))
}

/**
 * 截取字符串
 * @param str string 原始字符串
 * @param pos int 截取开始的字符串的下标(包含本下标)
 * @param length int 截取的字符串的长度
*/
func Substr(str string, pos int, length int)(resStr string, err error){
	strLen := len(str)
	if (pos + length > strLen){
		err = errors.New("字符串截取要求长度超过字符串的总长度")
		return
	} else {
		resStr = str[pos:pos+length]
	}
	return
}

/**
 * 判断路径是否存在，如果不存在，则创建
 * @param path string 路径
 * @return error err 执行过程中出现的问题
*/
func CreateDir(path string) (err error) {
	_, err = os.Stat(path)
	if err == nil {
		//目录已经存在
		return err
	}
	if os.IsNotExist(err) {
		//目录不存在,创建
		err = os.MkdirAll(path, 0755)
		return err
	}
	return err
}