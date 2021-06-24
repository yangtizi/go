package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache" // 使用前先import包
)

var ccc *gocache.Cache

// 5*time.Minute, 10*time.Minute
func New(defaultExpiration, cleanupInterval time.Duration) {
	ccc = gocache.New(defaultExpiration, cleanupInterval)
}

// Set s如果key不存在，增加一个kv记录；如果key已经存在，用新的value覆盖旧的value。
func Set(k string, x interface{}) {
	ccc.Set(k, x, gocache.DefaultExpiration)
}

// SetWithTime s如果key不存在，增加一个kv记录；如果key已经存在，用新的value覆盖旧的value。
// 对于有效时间d，如果是0（DefaultExpiration）使用默认有效时间；如果是-1（NoExpiration），表示没有过期时间。
func SetWithTime(k string, x interface{}, d time.Duration) {
	ccc.Set(k, x, d)
}

func Get(k string) (interface{}, bool) {
	return ccc.Get(k)
}

// Inc 对于cache中value是int, int8, int16, int32, int64, uintptr, uint,uint8, uint32, or uint64, float32,float64这些类型记录，可以使用该接口，将value值 +n。如果key不存在或value不是上述类型，会返回error。
func Inc(k string, n int64) error {
	return ccc.Increment(k, n)
}

// Dec 对于cache中value是int, int8, int16, int32, int64, uintptr, uint,uint8, uint32, or uint64, float32,float64这些类型记录，可以使用该接口，将value值 -n。如果key不存在或value不是上述类型，会返回error。
func Dec(k string, n int64) error {
	return ccc.Decrement(k, n)
}

// Delete 按照key删除记录，如果key不存在直接忽略，不会报错。
func Delete(k string) {
	ccc.Delete(k)
}

// Clean 将cache清空，删除所有记录。
func Clean() {
	ccc.Flush()
}
