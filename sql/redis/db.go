package redis

import (
	"errors"
	"sync"

	log "github.com/yangtizi/go/log/zaplog"
)

var mapRedis sync.Map

// Do (strAgent 代理商编号, strCommand sql脚本, args 脚本参数)
func Do(strAgent string, strCommand string, args ...interface{}) (interface{}, error) {
	v, ok := mapRedis.Load(strAgent)
	if !ok {
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TRedisDB).do(strCommand, args...)
}

// InitDB 初始化DB (strAgent 代理商编号, strReadConnect 从库连接字符串, strWriteConnect 主库连接字符串)
func InitDB(strAgent string, strConnect string) {
	_, ok := mapRedis.Load(strAgent)
	if !ok {
		// * 创建新的DB指针
		pRedis := &TRedisDB{}
		pRedis.init(strConnect)
		mapRedis.Store(strAgent, pRedis)
		return
	}

	log.Println("已经存在确有重复创建")
}
