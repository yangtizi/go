package mongo

import "gopkg.in/mgo.v2/bson"

// BsonM 新的
type BsonM bson.M

func (x BsonM) set() BsonM {
	return BsonM{"$set": x}
}

func (x BsonM) inc() BsonM {
	return BsonM{"$inc": x}
}
