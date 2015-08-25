package dbexecutor

import (
	_ "labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
)

var mgoexec *mgoSessions

const (
	DBUrl     = ""
	DBConnNum = 10
)

func Init() {
	mgoSessions.open(DBUrl, DBConnNum)
}

func FindOne(collName string, query interface{}, result interface{}) error {
	ses := mgoSessions.fetch()
	defer mgoSessions.push(ses)

	return ses.C(collName).Find(query).One(result)
}

func FindAll(collName string, query interface{}, result interface{}) error {
	ses := mgoSessions.fetch()
	defer mgoSessions.push(ses)

	return ses.C(collName).Find(query).All(result)
}

func Insert(collName string, document interface{}) error {
	ses := mgoSessions.fetch()
	defer mgoSessions.push(ses)

	return ses.C(collName).Insert(document)
}

func Update(collName string, selector interface{}, update interface{}) error {
	ses := mgoSessions.fetch()
	defer mgoSessions.push(ses)

	return ses.C(collName).Update(selector, update)
}

type mgoSessions struct {
	isInit      bool
	sessionChan chan *mgo.Session
}

func (self *mgoSessions) open(connectUrl string, sessionCount int) error {

	self.isInit = true
	self.sessionChan = make(chan *mgo.Session, sessionCount)

	for i := 0; i < sessionCount; i++ {
		session, err := mgo.Dial(connectUrl)
		if err != nil {
			return err
		}
		session.SetMode(mgo.Monotonic, true)
		self.sessionChan <- &session
	}
}

func (self *mgoSessions) fetch() *mgo.Session {
	return <-self.sessionChan
}

func (self *mgoSessions) push(ses *mgo.Session) {
	self.sessionChan <- ses
}

func (self *mgoSessions) close() {
	for i := 0; i < cap(self.sessionChan); i++ {
		ses <- self.sessionChan
		ses.Close()
	}
	self.isInit = false
}
