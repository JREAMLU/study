package main

import (
	"log"
	"time"

	"git.corp.plu.cn/plugo/infrastructure/redis"

	"github.com/JREAMLU/study/queryip/service"
	"github.com/wangtuanjie/ip17mon"
)

func init() {
	if err := ip17mon.Init("mydata4vipweek2.dat"); err != nil {
		panic(err)
	}

	if err := redis.LoadConfig("", "relationship", "cache", "vector", "login", "profile", "live"); err != nil {
		log.Panic(err)
	}

}

func main() {
	// ips := ip17mon.IPs[0:10]
	ips := ip17mon.IPs
	beginTime := time.Now().UnixNano()
	err := service.BatchIPS(ips)
	if err != nil {
		log.Println(err)
	}
	endTime := time.Now().UnixNano()
	takeTime := endTime - beginTime
	log.Printf("all time: %vms , %vns", takeTime/1000000, takeTime)
	log.Printf("total ips: %v", len(ips))
}
