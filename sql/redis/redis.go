package redis

import (
	"time"

	r "github.com/gomodule/redigo/redis"
)

// TRedisDB 单个的数据库
type TRedisDB struct {
	// chpool     chan int
	strConnect string
	// redisConn   r.Conn
	redisClient *r.Pool
}

func (m *TRedisDB) init(strConnect string) {
	m.redisClient = &r.Pool{
		MaxIdle:     30,
		MaxActive:   1000, // 最大连接数，
		IdleTimeout: 3600 * time.Second,
		Wait:        true,
		Dial: func() (r.Conn, error) {
			con, err := r.Dial("tcp", strConnect)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

	m.strConnect = strConnect
}

func (m *TRedisDB) do(strCommand string, args ...interface{}) (TValues, error) {
	rc := m.redisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	if rc.Err() != nil {
		return nil, rc.Err()
	}

	return r.Values(rc.Do(strCommand, args...))
}
