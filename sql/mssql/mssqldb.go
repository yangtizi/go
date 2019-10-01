package mssql

import (
	"database/sql"
	"errors"

	log "github.com/yangtizi/go/log/zaplog"
	"github.com/yangtizi/go/sql/scanner"
)

// TMsSQLDB 单个的数据库
type TMsSQLDB struct {
	chpool     chan int
	strConnect string
	pDB        *sql.DB
}

// NewDB 创建新的MSSQL数据库类
func NewDB(strReadConnect string) *TMsSQLDB {
	p := &TMsSQLDB{}
	p.init(strReadConnect)
	return p
}

func (self *TMsSQLDB) init(strConnect string) {
	db, err := sql.Open("mssql", strConnect)
	if err == nil {
		self.pDB = db
		self.strConnect = strConnect
		self.chpool = make(chan int, 30)
		return
	}

	log.Errorf("数据库连接出现问题 connect = [%s]", strConnect, " err = ", err)
}

func (self *TMsSQLDB) queryRow(strQuery string, args ...interface{}) (*scanner.TRow, error) {
	if self.pDB == nil {
		log.Errorf("queryRow [db == nil]")
		return nil, errors.New("不存在DB")
	}

	self.chpool <- 1
	row := self.pDB.QueryRow(strQuery, args...)
	<-self.chpool

	return scanner.NewRow(row), nil
}

func (self *TMsSQLDB) queryRows(strQuery string, args ...interface{}) (*scanner.TRows, error) {
	if self.pDB == nil {
		log.Errorf("queryRows [db == nil]")
		return nil, errors.New("不存在DB")
	}
	self.chpool <- 1
	rows, err := self.pDB.Query(strQuery, args...)
	<-self.chpool
	return scanner.NewRows(rows), err
}

func (self *TMsSQLDB) exec(strQuery string, args ...interface{}) (*scanner.TResult, error) {
	if self.pDB == nil {
		log.Errorf("exec [db == nil]")
		return nil, errors.New("不存在DB")
	}

	self.chpool <- 1
	rs, err := self.pDB.Exec(strQuery, args...)
	<-self.chpool
	return scanner.NewResult(rs), err
}
