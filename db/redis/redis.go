package redis

import (
	"Watermelon/common/config"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisClient *redis.Pool
)

func init() {
	RDhost := config.GetString("Redis_Host")
	RDport := config.GetString("Redis_Port")
	RDdb := config.GetString("Redis_DB")
	RDpassword := config.GetString("Redis_Password")
	RDMaxIdle := config.GetInt("Redis_MaxIdle")
	RDMaxActive := config.GetInt("Redis_MaxActive")
	/**********************redis初始化**************************************/
	RedisClient = &redis.Pool{ // 建立连接池
		MaxIdle:     RDMaxIdle, //从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxActive:   RDMaxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RDhost+":"+RDport, redis.DialPassword(RDpassword))
			if err != nil {
				return nil, err
			}
			c.Do("SELECT", RDdb)  // 选择db
			if RDpassword != "" { //密码校验
				if _, err := c.Do("AUTH", RDpassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
}

func RedisSet(key, value interface{}) error {
	rconn := RedisClient.Get()
	_, err := rconn.Do("SET", key, value)
	defer rconn.Close()
	return err

}
func RedisGet(key interface{}) (string, error) {
	rconn := RedisClient.Get()
	v, err := redis.String(rconn.Do("GET", key))
	defer rconn.Close()
	return v, err
}
func RedisExit(key interface{}) (bool, error) {
	rconn := RedisClient.Get()
	booler, err := redis.Bool(rconn.Do("EXISTS", key))
	defer rconn.Close()
	return booler, err
}
func RedisDelete(key interface{}) error {
	rconn := RedisClient.Get()
	_, err := rconn.Do("DEL", key)
	defer rconn.Close()
	return err
}
func RedisHMSET(key string, obj interface{}) error {
	rconn := RedisClient.Get()
	_, err := rconn.Do("HMSET", redis.Args{}.Add(key).AddFlat(obj)...)
	defer rconn.Close()
	return err
}
func RedisHMGET(key string) (map[string]string, error) {
	rconn := RedisClient.Get()
	v, err := redis.StringMap(rconn.Do("HGETALL", key))
	defer rconn.Close()
	return v, err
}
