package mongo

import (
	log "github.com/yangtizi/go/log/zaplog"
	"gopkg.in/mgo.v2"
)

// TMongoDB .
type TMongoDB struct {
	chpool     chan int
	strConnect string
	pDB        *mgo.Session
}

// NewDB 新DB
func NewDB(strConnect string) *TMongoDB {
	p := &TMongoDB{}
	p.init(strConnect)
	return p
}

func (m *TMongoDB) init(strConnect string) {
	mongo, err := mgo.Dial(strConnect)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	m.pDB = mongo
	m.strConnect = strConnect
}
