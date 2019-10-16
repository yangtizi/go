package tx

import (
	"database/sql"
)

// TTx 封装的
type TTx struct {
	tx *sql.Tx
}

// NewRow 新建
func NewRow(tx *sql.Tx) *TTx {
	t := &TTx{}
	t.tx = tx
	return t
}
