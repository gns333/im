package dbexecutor

import (
	_ "labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
)

func Init() {

}

func FindOne(collName string, query interface{}, result interface{}) error {

}

func FindAll(collName string, query interface{}, result interface{}) error {

}

func Insert(collName string, document interface{}) error {

}

func Update(collName string, selector interface{}, update interface{}) {

}
