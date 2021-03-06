//用户类
package model

import (
	"chatRoom/Common/db"
	"chatRoom/Common/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type UserModel struct {
	UserID int `json:"userid"`
	UserName string `json:"username"`
	UserPwd string `json:"userpwd"`
}

func NewUserModel(username string, userpwd string)(*UserModel){
	return &UserModel{
		UserID:0,
		UserName:username,
		UserPwd:userpwd,
	}
}

/*
 * 验证用户登录时的账号密码是否正确
 * @param name 用户名
 * @param pwd 密码
 * @return isOK 登录是否成功 true-登录成功 false-登录失败
 * @err 登录过程中产生的异常
*/
func (this *UserModel) CheckLogin(name string, pwd string)(isOK bool, err error){
	redisUserInfo, err := GetUserInfoByNameOnRedis(name)
	if redisUserInfo.UserPwd == pwd{
		isOK = true
	}
	return
}

/*
 * 用户注册函数
 * @return isOK 注册是否成功
 * @err 注册过程中产生的异常
*/
func (this *UserModel) Register()(isOK bool, err error){
	//先从数据库查询用户是否已经被注册过
	isIsset, err := this.CheckUserIsset()
	if isIsset{
		isOK = false
		return
	}

	//将用户信息写入到Mysql中
	id, err := this.UserInfoToMysql()
	if err != nil{
		return
	}
	this.UserID = id
	redisModel := db.NewRedisModel()
	defer redisModel.Conn.Close()
	userInfo, err := json.Marshal(this)
	if err != nil{
		fmt.Println("user.go 59 err:", err)
		return
	}
	_, err = redisModel.Conn.Do("Hset", "userInfo", this.UserName, string(userInfo))
	if err != nil{
		fmt.Println("user model register hset redis err:", err)
		return
	}
	return true, nil
}

/*
 * 检查用户是否已经注册
 * @return isOK true-已经注册过 false-没有注册过
 * @err 过程中产生的异常
*/
func (this *UserModel) CheckUserIsset()(isOk bool, err error){
	var id int
	err = db.MysqlDBPool.QueryRow("SELECT id FROM user WHERE user_name = ?", this.UserName).Scan(&id)
	if err == sql.ErrNoRows{
		//无查询记录
		isOk = false
	} else if err != nil{
		fmt.Println("server user register mysql query err:", err)
		return
	} else if id > 0{
		isOk = true
	}
	return isOk, nil
}

/*
 * 将用户信息写入到Mysql中
 * @return id 0-写入数据库错误， 非1-写入数据库的ID
 * @err 过程中产生的异常
*/
func (this *UserModel) UserInfoToMysql()(id int, err error){
	stmt, err := db.MysqlDBPool.Prepare("INSERT INTO user(user_name, user_pwd, cdate) VALUES(?,?, ?)")
	if err != nil{
		fmt.Println("user into mysql err:", err)
		return
	}
	defer stmt.Close()
	time, _ := util.GetNowTime()
	timeNow := time.Format(util.YMDHIS)
	ret, err := stmt.Exec(this.UserName, this.UserPwd, timeNow)
	if err != nil{
		fmt.Println("user into mysql stmt err:", err)
		return
	}
	insertID, err := ret.LastInsertId()
	if  err != nil{
		fmt.Println("user go 97 err:", err)
		return
	}
	id = int(insertID)
	return id, nil
}

func GetUserInfoByNameOnRedis(name string)(redisUserInfo *UserModel, err error){
	redisUserInfo = new(UserModel)
	redisModel := db.NewRedisModel()
	userInfo, err := redisModel.Conn.Do("Hget", "userInfo", name)
	if err != nil{
		fmt.Println("user.go 30 err:", err)
		return
	} else if err == redis.ErrNil{
		return
	}

	err = json.Unmarshal(userInfo.([]byte), redisUserInfo)
	if err != nil{
		fmt.Println("GetUserInfoByNameOnRedis json decode err:", err)
		return
	}
	return
}