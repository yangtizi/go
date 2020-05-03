package mysql

import (
	"database/sql"
	"errors"

	log "yangtizi/log/zaplog"
)

// TMySQLDB 单个的数据库
type TMySQLDB struct {
	chpool     chan int
	strConnect string
	pDB        *sql.DB
}

// NewDB 创建新的MSSQL数据库类
func NewDB(strReadConnect string) *TMySQLDB {
	p := &TMySQLDB{}
	p.init(strReadConnect)
	return p
}

func (self *TMySQLDB) init(strConnect string) {
	db, err := sql.Open("mysql", strConnect)
	if err == nil {
		self.pDB = db
		self.strConnect = strConnect
		self.chpool = make(chan int, 30)
		return
	}

	log.Println("数据库连接出现问题 connect = ", strConnect, " err = ", err)
}

func (self *TMySQLDB) queryRow(strQuery string, args ...interface{}) (*sql.Row, error) {
	if self.pDB == nil {
		return nil, errors.New("不存在DB")
	}

	self.chpool <- 1
	row := self.pDB.QueryRow(strQuery, args...)
	<-self.chpool
	return row, nil
}

func (self *TMySQLDB) queryRows(strQuery string, args ...interface{}) (*sql.Rows, error) {
	if self.pDB == nil {
		return nil, errors.New("不存在DB")
	}
	self.chpool <- 1
	rows, err := self.pDB.Query(strQuery, args...)
	<-self.chpool
	return rows, err
}

func (self *TMySQLDB) exec(strQuery string, args ...interface{}) (sql.Result, error) {

	if self.pDB == nil {
		return nil, errors.New("不存在DB")
	}

	self.chpool <- 1
	rs, err := self.pDB.Exec(strQuery, args...)
	<-self.chpool
	return rs, err
}
