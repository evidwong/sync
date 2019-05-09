package cache

import (
	"encoding/json"
	"time"

	"baisiyi.net/config"
	"github.com/gomodule/redigo/redis"
)

//Redis redis cache
type Redis struct {
	Conn *redis.Pool
}

//NewRedis 实例化
func NewRedis(cnf *config.Config) *Redis {
	pool := &redis.Pool{
		MaxIdle:   4,
		MaxActive: 0,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", cnf.Redis.Host+":"+cnf.Redis.Port)
			if err != nil {
				return nil, err
			}
			auth := cnf.Redis.Auth
			if auth != "" {
				if _, err := conn.Do("AUTH", auth); err != nil {
					conn.Close()
					return nil, err
				}
			}
			dbNum := cnf.Redis.DB
			if dbNum == "" {
				dbNum = "0"
			}
			if _, err := conn.Do("SELECT", dbNum); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
	}
	return &Redis{pool}
}

//SetConn 设置conn
func (r *Redis) SetConn(conn *redis.Pool) {
	r.Conn = conn
}

//Get 获取一个值
func (r *Redis) Get(key string) interface{} {
	conn := r.Conn.Get()
	defer conn.Close()

	var data []byte
	var err error
	if data, err = redis.Bytes(conn.Do("GET", key)); err != nil {
		return nil
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}

	return reply
}

//Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	conn := r.Conn.Get()
	defer conn.Close()

	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}

	_, err = conn.Do("SETEX", key, int64(timeout/time.Second), data)

	return
}

//IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	conn := r.Conn.Get()
	defer conn.Close()

	a, _ := conn.Do("EXISTS", key)
	i := a.(int64)
	if i > 0 {
		return true
	}
	return false
}

//Delete 删除
func (r *Redis) Delete(key string) error {
	conn := r.Conn.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}
