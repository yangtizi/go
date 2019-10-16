package mssql

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/denisenkom/go-mssqldb" // mssql 数据库
	log "github.com/yangtizi/go/log/zaplog"
	"github.com/yangtizi/go/sql/scanner"
)

var mapMSSQL sync.Map

// QueryRow (strAgent 代理商编号, strQuery sql脚本, args 脚本参数)
func QueryRow(strAgent string, strQuery string, args ...interface{}) (*scanner.TRow, error) {
	log.Debugf("strAgent = [%s], strQuery = [%s]", strAgent, strQuery)
	log.Debug("[+] ", args)
	v, ok := mapMSSQL.Load(strAgent)
	if !ok {
		log.Errorf("QueryRow 不存在索引")
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TMsSQLDB).queryRow(strQuery, args...)
}

// QueryRows (strAgent 代理商编号, strQuery sql脚本, args 脚本参数)
func QueryRows(strAgent string, strQuery string, args ...interface{}) (*scanner.TRows, error) {
	log.Debugf("strAgent = [%s], strQuery = [%s]", strAgent, strQuery)
	log.Debug("[+] ", args)
	v, ok := mapMSSQL.Load(strAgent)
	if !ok {
		log.Errorf("QueryRows 不存在索引")
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TMsSQLDB).queryRows(strQuery, args...)
}

// Exec (strAgent 代理商编号, strQuery sql脚本, args 脚本参数)
func Exec(strAgent string, strQuery string, args ...interface{}) (*scanner.TResult, error) {
	log.Debugf("strAgent = [%s], strQuery = [%s]", strAgent, strQuery)
	log.Debug("[+] ", args)
	v, ok := mapMSSQL.Load(strAgent)
	if !ok {
		log.Errorf("Exec 不存在索引")
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TMsSQLDB).exec(strQuery, args...)
}

// Transaction 事务
func Transaction(agent interface{}) (*sql.Tx, error) {
	log.Debugf("Transaction begin")
	v, ok := mapMSSQL.Load(agent)
	if !ok {
		log.Errorf("Transaction 不存在索引")
		return nil, errors.New("不存在的DB索引")
	}

	return v.(*TMsSQLDB).transaction()
}

// InitDB 初始化DB (strAgent 代理商编号, strConnect 从库连接字符串)
func InitDB(strAgent string, strConnect string) {
	_, ok := mapMSSQL.Load(strAgent)
	if !ok {
		// * 创建新的DB指针
		pMsSQL := NewDB(strConnect)

		log.Infof("正在连接数据库 strAgent = [%s], strConnect = [%s]", strAgent, strConnect)
		mapMSSQL.Store(strAgent, pMsSQL)
		return
	}

	log.Println("已经存在确有重复创建")
}
