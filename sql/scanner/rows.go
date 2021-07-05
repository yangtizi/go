package scanner

import (
	"database/sql"
	"log"
	"time"
)

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

// Unknown 未知搜索
func (m *TRows) Unknown() ([]*string, error) {
	columns, err := m.GetRows().Columns()
	if err != nil {
		return nil, err
	}
	nLen := len(columns)

	// 根据长度来设置
	args2 := []interface{}{}

	for i := 0; i < nLen; i++ {
		args2 = append(args2, &sql.NullString{})
	}

	err = m.rows.Scan(args2...)
	if err != nil {
		return nil, err
	}

	args1 := []*string{}

	for _, v := range args2 {
		// 插入进来
		args1 = append(args1, &(v.(*sql.NullString).String))
	}

	return args1, nil
}

// Scan 新搜索
func (m *TRows) Scan(args ...interface{}) error {
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

	err := m.rows.Scan(args2...)
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
func (m *TRows) Next() bool {
	return m.rows.Next()
}

// Close 关闭
func (m *TRows) Close() error {
	return m.rows.Close()
}

// GetRows 得到Rows
func (m *TRows) GetRows() *sql.Rows {
	return m.rows
}
