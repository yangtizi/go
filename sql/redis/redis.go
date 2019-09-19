package redis

import (
	"errors"

	r "github.com/gomodule/redigo/redis"
)

// TRedisDB 单个的数据库
type TRedisDB struct {
	strConnect string
	redisConn  r.Conn
}

func (self *TRedisDB) init(strConnect string) {

	if conn, err := r.Dial("tcp", strConnect); err == nil {
		self.redisConn = conn
		self.strConnect = strConnect
	}
}

func (self *TRedisDB) do(strCommand string, args ...interface{}) (interface{}, error) {
	if self.redisConn == nil {
		return nil, errors.New("不存在DB")
	}

	return self.redisConn.Do(strCommand, args...)
}

func (self *TRedisDB) hmget(args ...interface{}) ([]interface{}, error) {
	if self.redisConn == nil {
		return nil, errors.New("不存在DB")
	}

	return r.Values(self.redisConn.Do("hmget", args...))
}
