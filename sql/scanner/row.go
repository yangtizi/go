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
func (m *TRow) Scan(args ...interface{}) error {
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

	err := m.row.Scan(args2...)
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
