package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache" // 使用前先import包
	"github.com/yangtizi/go/log/zaplog"
)

type TCache struct {
	cache *gocache.Cache
}

var ccc TCache

func init() {
	if ccc.cache == nil {
		ccc.cache = gocache.New(12*time.Hour, 10*time.Minute)
	}
}

// 5*time.Minute, 10*time.Minute
func NewCache(defaultExpiration, cleanupInterval time.Duration) *TCache {
	zaplog.Infof("缓存正在初始化, 超时时间=[%v], 清理时间=[%v]", defaultExpiration, cleanupInterval)
	ccc.cache = gocache.New(defaultExpiration, cleanupInterval)
	return &ccc
}

// Set s如果key不存在，增加一个kv记录；如果key已经存在，用新的value覆盖旧的value。
func Set(k string, x interface{}) {
	ccc.cache.Set(k, x, gocache.DefaultExpiration)
}

// SetWithTime s如果key不存在，增加一个kv记录；如果key已经存在，用新的value覆盖旧的value。
// 对于有效时间d，如果是0（DefaultExpiration）使用默认有效时间；如果是-1（NoExpiration），表示没有过期时间。
func SetWithTime(k string, x interface{}, d time.Duration) {
	ccc.cache.Set(k, x, d)
}

func Get(k string) (interface{}, bool) {
	return ccc.cache.Get(k)
}

// Inc 对于cache中value是int, int8, int16, int32, int64, uintptr, uint,uint8, uint32, or uint64, float32,float64这些类型记录，可以使用该接口，将value值 +n。如果key不存在或value不是上述类型，会返回error。
func Inc(k string, n int64) error {
	return ccc.cache.Increment(k, n)
}

// Dec 对于cache中value是int, int8, int16, int32, int64, uintptr, uint,uint8, uint32, or uint64, float32,float64这些类型记录，可以使用该接口，将value值 -n。如果key不存在或value不是上述类型，会返回error。
func Dec(k string, n int64) error {
	return ccc.cache.Decrement(k, n)
}

// Delete 按照key删除记录，如果key不存在直接忽略，不会报错。
func Delete(k string) {
	ccc.cache.Delete(k)
}

// Clean 将cache清空，删除所有记录。
func Clean() {
	ccc.cache.Flush()
}

// 保存到文件
func SaveToFile(strFilename string) error {
	return ccc.cache.SaveFile(strFilename)
}

//
func LoadFromFile(strFilename string) error {
	return ccc.cache.LoadFile(strFilename)
}
