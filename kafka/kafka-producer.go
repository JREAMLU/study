package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	log "github.com/thinkboy/log4go"
)

const (
	KAFKA_TOPIC = "live"
	KAFKA_ADDRS = `172.16.9.80:9092`
)

var (
	producer sarama.AsyncProducer
)

func InitKafka(kafkaAddrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err = sarama.NewAsyncProducer(kafkaAddrs, config)
	go handleSuccess()
	go handleError()
	return
}

func handleSuccess() {
	var (
		pm *sarama.ProducerMessage
	)
	for {
		pm = <-producer.Successes()
		if pm != nil {
			log.Info("producer message success, partition:%d offset:%d key:%v valus:%s", pm.Partition, pm.Offset, pm.Key, pm.Value)
		}
	}
}

func handleError() {
	var (
		err *sarama.ProducerError
	)
	for {
		err = <-producer.Errors()
		if err != nil {
			log.Error("producer message error, partition:%d offset:%d key:%v valus:%s error(%v)", err.Msg.Partition, err.Msg.Offset, err.Msg.Key, err.Msg.Value, err.Err)
		}
	}
}

func pushKafka(str string) (err error) {
	producer.Input() <- &sarama.ProducerMessage{Topic: KAFKA_TOPIC, Value: sarama.StringEncoder(str)}
	return
}

func main() {
	kafkaAddrs := []string{KAFKA_ADDRS}
	InitKafka(kafkaAddrs)
	str := "abc123"
	err := pushKafka(str)
	if err != nil {
		fmt.Println("pushKafka err: ", err)
	}
}
