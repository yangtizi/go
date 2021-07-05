package scanner

import "database/sql"

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
func (m *TResult) RowsAffected() (int64, error) {
	return m.r.RowsAffected()
}
