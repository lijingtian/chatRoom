package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var MysqlDBPool *sql.DB

func init(){
	var err error
	MysqlDBPool, err = sql.Open("mysql", mysqlDSN)
	if err != nil{
		fmt.Println("mysql open err:", err)
	}
	MysqlDBPool.SetMaxOpenConns(10)
	MysqlDBPool.SetMaxIdleConns(5)
	//使连接池 连接数据库
	MysqlDBPool.Ping()
}

