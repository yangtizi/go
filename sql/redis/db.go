package redis

import (
	"errors"
	"sync"

	"github.com/yangtizi/go/sysutils"

	log "github.com/yangtizi/go/log/zaplog"
)

// TValues 快速解析
type TValues []interface{}

// I 快速获取整数
func (self TValues) I(n int) int {
	if n > len(self) {
		return 0
	}
	return sysutils.StrToIntDef(string(self[n].([]byte)), 0)
}

// S 快速获取字符串
func (self TValues) S(n int) string {
	if n > len(self) {
		return ""
	}
	return string(self[n].([]byte))
}

var mapRedis sync.Map

// Do (strAgent 代理商编号, strCommand sql脚本, args 脚本参数)
func Do(strAgent string, strCommand string, args ...interface{}) (interface{}, error) {
	v, ok := mapRedis.Load(strAgent)
	if !ok {
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TRedisDB).do(strCommand, args...)
}

// HMGet (strAgent 代理商编号, args 脚本参数)
func HMGet(strAgent string, args ...interface{}) (TValues, error) {
	v, ok := mapRedis.Load(strAgent)
	if !ok {
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TRedisDB).hmget(args...)
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
