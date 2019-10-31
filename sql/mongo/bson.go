package mongo

import "gopkg.in/mgo.v2/bson"

// M 新的
type M bson.M

// Set .
func (x M) Set() M {
	return M{"$set": x}
}

// Inc .
func (x M) Inc() M {
	return M{"$inc": x}
}

// Match .
func (x M) Match() M {
	return M{"$match": x}
}

// Group 组合起来, GroupBy 语句, 根据某个值来
// usage M{ "_id", "$keyid"} // Select * From Table order by [keyid]
func (x M) Group() M {
	return M{"$group": x}
}

// Sort 排列
func (x M) Sort() M {
	return M{"$sort": x}
}

// OrderBy 按顺序排列  1 从小到到  -1 从大到小
func (x M) OrderBy(s string, n int) M {
	return M{"$sort": M{s: n}}
}

// Limit .
func (x M) Limit(n int) M {
	return M{"$limit": n}
}

// Where .
func (x M) Where(s string, n int) M {
	switch s {
	case "=", "==", "$eq", "eq":
		return M{"$eq": n} // 等于

	case "$lt", "<":
		return M{"$lt": n} // 小于

	case "<=", "$lte", "lte":
		return M{"$lte": n} // 小于或等于

	case ">", "$gt", "gt":
		return M{"$gt": n} // 大于

	case ">=", "$gte", "gte":
		return M{"$gte": n} // 大于或等于

	case "<>", "!=", "~=", "$ne", "ne":
		return M{"$ne": n} // 不等于
	}
	return M{}
}
