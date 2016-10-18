package main

import (
	"fmt"
	"log"

	"github.com/JREAMLU/core/db/mongodb"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mlog struct {
	Id        bson.ObjectId `bson:"_id"`
	Taskid    uint64        `bson:"taskid"`
	Pid       uint64        `bson:"pid"`
	Log       string        `bson:"log"`
	Shell     string        `bson:"shell"`
	ErrorLog  string        `bson:"errorlog"`
	CreatedAt int64         `bson:"create_at"`
}

var mongoClient *mongo.MongoClient

func InitMongo(url, dbName string) {
	mongoClient = mongo.NewMongoClient(url, dbName)
}

func GetLogs() []Mlog {
	session, err := mongoClient.Session()
	if err != nil {
		log.Println("mongodb err: ", err)
	}
	defer session.Close()
	c := session.DB(mongoClient.DBName).C("cron")

	var mlog []Mlog
	err = c.Find(bson.M{"log": "ppppppp"}).All(&mlog)
	if err != nil {
		log.Fatal(err)
	}

	return mlog
}

func AddLogs(j *Mlog) (id interface{}, err error) {
	session, err := mongoClient.Session()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(mongoClient.DBName).C("cron")

	if !mongoClient.CollectExists(session.DB(mongoClient.DBName), "cron") {
		cIndex := mgo.Index{Key: []string{"pid"}, Unique: false, Background: false}
		c.EnsureIndex(cIndex)
	}

	j.Id = bson.NewObjectId()

	err = c.Insert(j)

	if err != nil {
		return nil, err
	}
	return j.Id, nil
}

func AddLogsBatch(docs []interface{}) (id interface{}, err error) {
	session, err := mongoClient.Session()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(mongoClient.DBName).C("cron")

	if !mongoClient.CollectExists(session.DB(mongoClient.DBName), "cron") {
		cIndex := mgo.Index{Key: []string{"pid"}, Unique: false, Background: false}
		c.EnsureIndex(cIndex)
	}

	b := c.Bulk()
	b.Insert(docs...)
	// b.Insert(docs...)
	// b.Insert(docs...)
	res, err := b.Run()

	fmt.Println("res", res)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func AddLogsBat(docs []interface{}) (id interface{}, err error) {
	session, err := mongoClient.Session()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(mongoClient.DBName).C("cron")

	if !mongoClient.CollectExists(session.DB(mongoClient.DBName), "cron") {
		cIndex := mgo.Index{Key: []string{"pid"}, Unique: false, Background: false}
		c.EnsureIndex(cIndex)
	}

	err = c.Insert(docs...)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
