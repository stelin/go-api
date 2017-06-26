package mgo

import (
	"fmt"
	//	"time"

	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

var mgoSession *mgo.Session
var mgoDataBase string

func init() {
	var err error
	mgoSession, err = mgo.Dial("mongodb://192.168.101.3:27017/user?maxPoolSize=20")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	mgoSession.SetMode(mgo.Monotonic, true)

	mgoDataBase = "user"
}

func Adds(colName string, docs ...interface{}) error {

	//	fmt.Fprintln("insert -----", bson.M)
	session := mgoSession.Copy()
	defer session.Close()
	colection := session.DB(mgoDataBase).C(colName)
	err := colection.Insert(docs...)
	return err
}

func FindOne(colName string, where map[string]interface{}, result interface{}) error {
	session := mgoSession.Copy()
	defer session.Close()
	colection := session.DB(mgoDataBase).C(colName)
	data := colection.Find(where).One(result)
	return data
}

func Query(colName string, where map[string]interface{}, result interface{}) error {
	session := mgoSession.Copy()
	defer session.Close()
	colection := session.DB(mgoDataBase).C(colName)
	data := colection.Find(where).All(result)
	return data
}
