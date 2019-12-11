package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/nameouyang/learning-go/conf"
	"strconv"
	"time"
)

var redisConn *redis.Pool

func GetRedisConn() *redis.Pool {
	return redisConn
}
func init() {
	redisConn = &redis.Pool{
		MaxIdle:     conf.RedisConf.MaxIdle,
		MaxActive:   conf.RedisConf.MaxActive,
		IdleTimeout: conf.RedisConf.IdleTimeout,
		Dial: func() (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", conf.RedisConf.Host+":"+strconv.Itoa(conf.RedisConf.Port))
			if err != nil {
				return nil, err
			}
			// 验证密码
			if conf.RedisConf.Password != "" {
				if _, err := conn.Do("AUTH", conf.RedisConf.Password); err != nil {
					_ = conn.Close()
					return nil, err
				}
			}
			if _, err := conn.Do("SELECT", conf.RedisConf.DBNum); err != nil {
				_ = conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func Set(key string, data string, seconds int) error {
	conn := GetRedisConn().Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, data)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, seconds)
	if err != nil {
		return err
	}
	return nil
}
func Get(key string) (string, error) {
	conn := GetRedisConn().Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("GET", key))
	if err != nil || err == redis.ErrNil {
		return "", err
	}
	return reply, nil
}
func Exists(key string) bool {
	conn := GetRedisConn().Get()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}
func Del(key string) (bool, error) {
	conn := GetRedisConn().Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
