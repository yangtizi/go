package mysql

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/go-sql-driver/mysql" // mysql 数据库
	log "github.com/yangtizi/go/log/zaplog"
)

var mapMYSQL sync.Map

// QueryRow (strAgent 代理商编号, strQuery sql脚本, args 脚本参数)
func QueryRow(strAgent string, strQuery string, args ...interface{}) (*sql.Row, error) {
	v, ok := mapMYSQL.Load(strAgent)
	if !ok {
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TMySQLDB).queryRow(strQuery, args...)
}

// QueryRows (strAgent 代理商编号, strQuery sql脚本, args 脚本参数)
func QueryRows(strAgent string, strQuery string, args ...interface{}) (*sql.Rows, error) {
	v, ok := mapMYSQL.Load(strAgent)
	if !ok {
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TMySQLDB).queryRows(strQuery, args...)
}

// Exec (strAgent 代理商编号, strQuery sql脚本, args 脚本参数)
func Exec(strAgent string, strQuery string, args ...interface{}) (sql.Result, error) {
	v, ok := mapMYSQL.Load(strAgent)
	if !ok {
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TMySQLDB).exec(strQuery, args...)
}

// InitDB 初始化DB (strAgent 代理商编号, strConnect 从库连接字符串)
func InitDB(strAgent string, strConnect string) {
	_, ok := mapMYSQL.Load(strAgent)
	if !ok {
		// * 创建新的DB指针
		pMsSQL := NewDB(strConnect)

		log.Println("正在连接数据库 ", strAgent, strConnect)
		mapMYSQL.Store(strAgent, pMsSQL)
		return
	}

	log.Println("已经存在确有重复创建")
}
