package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Jlog struct {
	Pid    uint64
	Log    string
	Action string
}

func main() {
	session, err := mgo.Dial("mongodb://172.16.9.221:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cron").C("cron")
	// err = c.Insert(
	// 	&Jlog{Pid: 10, Log: "ppppppp", Action: "do10"},
	// 	&Jlog{Pid: 11, Log: "zzzzzzzz", Action: "do11"})
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// log.Printf("Insert Success")
	var jlog Jlog
	err = c.Find(bson.M{"pid": 11}).One(&jlog)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("jlog: ", jlog)
}
