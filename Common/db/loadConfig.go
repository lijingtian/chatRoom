/*
 * 加载配置文件，如果配置文件不存在，报错
*/
package db

import (
	"chatRoom/config"
	"fmt"
	"github.com/Unknwon/goconfig"
)
var mysqlDSN string
var redisHost string
func init(){
	var configPath string = config.RootPath + "/config/db.ini";
	cfg, err := goconfig.LoadConfigFile(configPath)
	if err != nil{
		fmt.Println(err)
		return
	}
	mysqlName, err := cfg.GetValue("mysql", "username")
	if err != nil{
		fmt.Println(err)
		return
	}
	mysqlPwd, err := cfg.GetValue("mysql", "password")
	if err != nil{
		fmt.Println(err)
		return
	}
	mysqlUrl, err := cfg.GetValue("mysql", "url")
	if err != nil{
		fmt.Println(err)
		return
	}
	mysqlCharset, err := cfg.GetValue("mysql", "charset")
	if err != nil{
		fmt.Println(err)
		return
	}
	//"root:@tcp(127.0.0.1:3306)/chatRoom?charset=utf8"
	mysqlDSN = mysqlName + ":" + mysqlPwd + "@tcp" + mysqlUrl + "?charset=" + mysqlCharset

	redisHost, err = cfg.GetValue("redis", "address")
	if err != nil{
		fmt.Println(err)
		return
	}

}
