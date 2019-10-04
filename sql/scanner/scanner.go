package scanner

import (
	"database/sql"
	"log"
	"time"
)

// TRow 封装的
type TRow struct {
	row *sql.Row
}

// NewRow 新建
func NewRow(row *sql.Row) *TRow {
	r := &TRow{}
	r.row = row
	return r
}

// Scan 新的检索
func (self *TRow) Scan(args ...interface{}) error {
	args2 := []interface{}{}

	for _, v := range args {
		switch d := v.(type) {
		case *int, *int64:
			args2 = append(args2, &sql.NullInt64{})
		case *string, *time.Time:
			args2 = append(args2, &sql.NullString{})
		default:
			log.Println(v, d, "不知道的类型")
		}
	}

	err := self.row.Scan(args2...)
	if err != nil {
		return err
	}

	for i, v := range args {
		switch d := v.(type) {
		case *int:
			*d = int(args2[i].(*sql.NullInt64).Int64)
		case *int64:
			*d = args2[i].(*sql.NullInt64).Int64
		case *string:
			*d = args2[i].(*sql.NullString).String
		case *time.Time:
			log.Println(args2[i].(*sql.NullString))
			*d, _ = time.ParseInLocation("2006-01-02T15:04:05.999Z", args2[i].(*sql.NullString).String, time.Local)
			log.Println(*d)
		default:
			log.Println("不知道的类型222")
		}
	}
	return nil
}

// TRows 封装
type TRows struct {
	rows *sql.Rows
}

// NewRows 新建
func NewRows(rows *sql.Rows) *TRows {
	r := &TRows{}
	r.rows = rows
	return r
}

// Scan 新搜索
func (self *TRows) Scan(args ...interface{}) error {
	args2 := []interface{}{}

	for _, v := range args {
		switch d := v.(type) {
		case *int, *int64:
			args2 = append(args2, &sql.NullInt64{})
		case *string, *time.Time:
			args2 = append(args2, &sql.NullString{})
		default:
			args2 = append(args2, &sql.NullString{})
			log.Println(v, d, "不知道的类型")
		}
	}

	err := self.rows.Scan(args2...)
	if err != nil {
		return err
	}

	for i, v := range args {
		switch d := v.(type) {
		case *int:
			*d = int(args2[i].(*sql.NullInt64).Int64)
		case *int64:
			*d = args2[i].(*sql.NullInt64).Int64
		case *string:
			*d = args2[i].(*sql.NullString).String
		case *time.Time:
			log.Println(args2[i].(*sql.NullString))
			*d, _ = time.ParseInLocation("2006-01-02T15:04:05.999Z", args2[i].(*sql.NullString).String, time.Local)
			log.Println(*d)
		default:
			log.Fatalln("不知道的类型222")
		}
	}
	return nil
}

// Next 下一步
func (self *TRows) Next() bool {
	return self.rows.Next()
}

// Close 关闭
func (self *TRows) Close() error {
	return self.rows.Close()
}

// GetRows 得到Rows
func (self *TRows) GetRows() *sql.Rows {
	return self.rows
}

// TResult 结果
type TResult struct {
	r sql.Result
}

// NewResult 新结果
func NewResult(r sql.Result) *TResult {
	rr := &TResult{}
	rr.r = r
	return rr
}

// RowsAffected 数量
func (self *TResult) RowsAffected() (int64, error) {
	return self.r.RowsAffected()
}
