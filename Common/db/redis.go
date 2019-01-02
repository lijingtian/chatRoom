package db

import (
	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool
type RedisModel struct {
	Conn redis.Conn
}

func init(){
	redisPool = &redis.Pool{
		MaxIdle: 10,
		MaxActive : 0,
		IdleTimeout: 300,
		Dial: func() (redisHandle redis.Conn, err error) {
			redisHandle, err = redis.Dial("tcp", redisHost)
			return
		},
	}
}

func NewRedisModel() (redisModel *RedisModel){
	redisModel = new(RedisModel)
	redisModel.getConn()
	return
}

func(this *RedisModel) getConn(){
	this.Conn = redisPool.Get()
}