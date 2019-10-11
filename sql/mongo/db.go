package mongo

import (
	"errors"
	"sync"

	_ "github.com/denisenkom/go-mssqldb" // mssql 数据库
	log "github.com/yangtizi/go/log/zaplog"
	"gopkg.in/mgo.v2"
)

var mapMongo sync.Map

// GetTable 获取表格
func GetTable(strAgent, strDatabase, strCollection string) (*mgo.Collection, error) {
	log.Debugf("%s, %s, %s", strAgent, strDatabase, strCollection)
	v, ok := mapMongo.Load(strAgent)
	if !ok {
		log.Errorf("Exec 不存在索引")
		return nil, errors.New("不存在的DB索引")
	}
	return v.(*TMongoDB).pDB.DB(strDatabase).C(strCollection), nil
}

// InitDB 初始化DB (strAgent 代理商编号, strConnect 从库连接字符串)
func InitDB(strAgent string, strConnect string) {
	_, ok := mapMongo.Load(strAgent)

	if !ok {
		// * 创建新的DB指针
		pMongo := NewDB(strConnect)

		log.Infof("正在连接数据库 strAgent = [%s], strConnect = [%s]", strAgent, strConnect)
		mapMongo.Store(strAgent, pMongo)
		return
	}

	log.Println("已经存在确有重复创建")
}
