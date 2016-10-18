package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

var count int = 1

func init() {
	InitMongo("mongodb://172.16.9.221:27017", "cron")
}

func main() {
	insert()

	// for i := 0; i < 500; i++ {
	// 	go insert()
	// }
	//
	// select {}
}

func insert() {
	var docs []interface{}
	for i := 0; i < 100; i++ {
		var m Mlog
		m.Id = bson.NewObjectId()
		m.Taskid = 1
		docs = append(docs, m)
	}

	// id, err := AddLogsBatch(docs)
	id, err := AddLogsBat(docs)
	if err != nil {
		fmt.Println("mgo err: ", err)
	} else {
		count++
		fmt.Println("id:", id, "-", count)
	}

}

func bat(d interface{}) {
	var docs []interface{}
	docs = append(docs, d)
	fmt.Println(len(docs))
}
