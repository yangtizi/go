package mongo

import (
	"errors"
	"sync"

	log "yangtizi/log/zaplog"

	_ "github.com/denisenkom/go-mssqldb" // mssql 数据库
	"gopkg.in/mgo.v2"
)

var mapMongo sync.Map
var instance *TMongoDB

// GetTable 获取表格
// func GetTable(agent interface{}, strDatabase, strCollection string) (*mgo.Collection, error) {
// 	if agent == nil {
// 		if instance == nil {
// 			return nil, errors.New("不存在的DB索引")
// 		}
// 		return instance.pDB.DB(strDatabase).C(strCollection), nil
// 	}
// 	log.Debugf("%v, %s, %s", agent, strDatabase, strCollection)
// 	v, ok := mapMongo.Load(agent)
// 	if !ok {
// 		log.Errorf("Exec 不存在索引")
// 		return nil, errors.New("不存在的DB索引")
// 	}
// 	return v.(*TMongoDB).pDB.DB(strDatabase).C(strCollection), nil
// }

// GetDB 获取DB
func GetDB(agent interface{}) (*mgo.Session, error) {
	if agent == nil {
		if instance == nil {
			log.Errorf("不存在的DB索引")
			return nil, errors.New("不存在的DB索引")
		}

		return instance.pDB.Copy(), nil
	}

	v, ok := mapMongo.Load(agent)
	if !ok {
		log.Errorf("Exec 不存在索引")
		return nil, errors.New("不存在的DB索引")
	}
	return v.(*TMongoDB).pDB.Copy(), nil
}

// InitDB 初始化DB (strAgent 代理商编号, strConnect 从库连接字符串)
func InitDB(agent interface{}, strConnect string) {
	if agent == nil {
		instance = NewDB(strConnect)
		log.Infof("正在连接数据库 agent = [%v], strConnect = [%s]", "默认", strConnect)
		return
	}

	_, ok := mapMongo.Load(agent)

	if !ok {
		// * 创建新的DB指针
		pMongo := NewDB(strConnect)

		log.Infof("正在连接数据库 agent = [%v], strConnect = [%s]", agent, strConnect)
		mapMongo.Store(agent, pMongo)
		return
	}

	log.Println("已经存在确有重复创建")
}
