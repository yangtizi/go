package mongo

import (
	log "github.com/yangtizi/go/log/zaplog"
	"gopkg.in/mgo.v2"
)

type TMongoDB struct {
	chpool     chan int
	strConnect string
	pDB        *mgo.Session
}

func NewDB(strConnect string) *TMongoDB {
	p := &TMongoDB{}
	p.init(strConnect)
	return p
}

func (self *TMongoDB) init(strConnect string) {
	mongo, err := mgo.Dial(strConnect)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	self.pDB = mongo
	self.strConnect = strConnect
}
