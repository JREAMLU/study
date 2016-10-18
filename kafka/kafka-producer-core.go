package main

import "study/core"

func main() {
	core.KafkaProducerPush("cron", `{"name":"jream","age":23}`, "172.16.9.4:9092,172.16.9.4:9092")
}
