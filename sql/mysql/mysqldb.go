package mysql

import (
	"database/sql"
	"errors"

	"github.com/yangtizi/go/log/zaplog"
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

func (m *TMySQLDB) init(strConnect string) {
	db, err := sql.Open("mysql", strConnect)
	if err == nil {
		m.pDB = db
		m.strConnect = strConnect
		m.chpool = make(chan int, 30)
		return
	}

	zaplog.Errorf("数据库连接出现问题 connect = [%s] err = []", strConnect, err)
}

func (m *TMySQLDB) queryRow(strQuery string, args ...interface{}) (*sql.Row, error) {
	if m.pDB == nil {
		return nil, errors.New("不存在DB")
	}

	m.chpool <- 1
	row := m.pDB.QueryRow(strQuery, args...)
	<-m.chpool
	return row, nil
}

func (m *TMySQLDB) queryRows(strQuery string, args ...interface{}) (*sql.Rows, error) {
	if m.pDB == nil {
		return nil, errors.New("不存在DB")
	}
	m.chpool <- 1
	rows, err := m.pDB.Query(strQuery, args...)
	<-m.chpool
	return rows, err
}

func (m *TMySQLDB) exec(strQuery string, args ...interface{}) (sql.Result, error) {

	if m.pDB == nil {
		return nil, errors.New("不存在DB")
	}

	m.chpool <- 1
	rs, err := m.pDB.Exec(strQuery, args...)
	<-m.chpool
	return rs, err
}
