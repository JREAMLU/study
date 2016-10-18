package main

import (
	"fmt"
	"strings"
	"study/core"
	"time"
)

func prehookF() {
	fmt.Println("prehook")
}

func handleF(msg string) {
	fmt.Println(msg)
}

func posthookF() {
	fmt.Println("posthook")
}

func main() {
	topic := "cron"
	groupname := "go"
	zkroot := ""
	zkaddrs := "172.16.9.4:2181,172.16.9.4:2181"
	pt := 10
	ci := 10

	var kafkaConsumer core.KafkaConsumer
	kafkaConsumer.Topic = topic
	kafkaConsumer.GoupName = groupname
	kafkaConsumer.ZKRoot = zkroot
	kafkaConsumer.ZKAddrs = strings.Split(zkaddrs, ",")
	kafkaConsumer.ProcessingTimeout = time.Duration(pt) * time.Second
	kafkaConsumer.CommitInterval = time.Duration(ci) * time.Second
	kafkaConsumer.InitKafkaConsumer(prehookF, handleF, posthookF)

	for {
	}
}
