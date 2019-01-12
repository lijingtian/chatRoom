/*
 * 加载配置文件，如果配置文件不存在，记录日志
*/
package db

import (
	"chatRoom/Common/log"
	"chatRoom/config"
	"github.com/Unknwon/goconfig"
	"github.com/logrus"
)
var mysqlDSN string
var redisHost string
func init(){
	var configPath string = config.RootPath + "/config/db.ini";
	cfg, err := goconfig.LoadConfigFile(configPath)
	if err != nil{
		log.CommonLog.WithFields(
			logrus.Fields{
				"err": err,
				"db_path": configPath,
			}).Warn("db.ini no find")
		return
	}
	mysqlName, err := cfg.GetValue("mysql", "username")
	if err != nil{
		log.CommonLog.WithFields(
			logrus.Fields{
				"err": err,
			},
			).Warn("get mysql username err")
		return
	}
	mysqlPwd, err := cfg.GetValue("mysql", "password")
	if err != nil{
		log.CommonLog.WithFields(
			logrus.Fields{
				"err": err,
			}).Warn("get mysql password err")
		return
	}
	mysqlUrl, err := cfg.GetValue("mysql", "url")
	if err != nil{
		log.CommonLog.WithFields(
			logrus.Fields{
				"err": err,
		}).Warn("get mysql url err")
		return
	}
	mysqlCharset, err := cfg.GetValue("mysql", "charset")
	if err != nil{
		log.CommonLog.WithFields(
			logrus.Fields{
				"err": err,
			}).Warn("get mysql charset err")
		return
	}
	mysqlDSN = mysqlName + ":" + mysqlPwd + "@tcp" + mysqlUrl + "?charset=" + mysqlCharset

	redisHost, err = cfg.GetValue("redis", "address")
	if err != nil{
		log.CommonLog.WithFields(
			logrus.Fields{
				"err": err,
			}).Warn("get redis host err")
		return
	}
}
