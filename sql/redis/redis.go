package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
	r "github.com/gomodule/redigo/redis"
)

// TRedisDB 单个的数据库
type TRedisDB struct {
	chpool     chan int
	strConnect string
	// redisConn   r.Conn
	redisClient *r.Pool
}

func (self *TRedisDB) init(strConnect string) {
	self.redisClient = &redis.Pool{
		MaxIdle:     30,
		MaxActive:   1000, // 最大连接数，
		IdleTimeout: 3600 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", strConnect)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

	self.strConnect = strConnect
}

func (self *TRedisDB) do(strCommand string, args ...interface{}) (TValues, error) {
	rc := self.redisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	if rc.Err() != nil {
		return nil, rc.Err()
	}

	return r.Values(rc.Do(strCommand, args...))
}
