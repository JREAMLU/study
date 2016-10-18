package main

import (
	"fmt"
	llog "log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/thinkboy/log4go"
	"github.com/wvanbergen/kafka/consumergroup"
)

type KafkaConsumer struct {
	Topic    string
	GoupName string
	ZKRoot   string
	ZKAddrs  []string
}

const (
	KAFKA_TOPIC                        = "live"
	KAFKA_ADDRS                        = `172.16.9.80:9092`
	KAFKA_CHROOT                       = "" //kafka
	KAFKA_ZKADDRS                      = `172.16.9.80:2181`
	KAFKA_GROUP_NAME                   = "go"
	OFFSETS_COMMIT_INTERVAL            = 10 * time.Second
	OFFSETS_PROCESSING_TIMEOUT_SECONDS = 10 * time.Second
)

func (kc *KafkaConsumer) InitKafka() error {
	log.Info("start topic:%s consumer", kc.Topic)
	log.Info("consumer group name:%s", kc.GoupName)
	sarama.Logger = llog.New(os.Stdout, "[Sarama] ", llog.LstdFlags)
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = OFFSETS_PROCESSING_TIMEOUT_SECONDS
	config.Offsets.CommitInterval = OFFSETS_COMMIT_INTERVAL
	config.Zookeeper.Chroot = kc.ZKRoot
	kafkaTopics := []string{kc.Topic}

	cg, err := consumergroup.JoinConsumerGroup(kc.GoupName, kafkaTopics, kc.ZKAddrs, config)
	if err != nil {
		return err
	}
	go func() {
		for err := range cg.Errors() {
			log.Error("consumer error(%v)", err)
		}
	}()
	go func() {
		for msg := range cg.Messages() {
			log.Info("deal with topic:%s, partitionId:%d, Offset:%d, Key:%s msg:%s \n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
			log.Info("handle msg: % s \n", msg.Value)
			cg.CommitUpto(msg)
		}
	}()
	return nil
}

func main() {
	var kafkaConsumer KafkaConsumer
	kafkaConsumer.Topic = KAFKA_TOPIC
	kafkaConsumer.GoupName = KAFKA_GROUP_NAME
	kafkaConsumer.ZKRoot = KAFKA_CHROOT
	kafkaConsumer.ZKAddrs = []string{KAFKA_ZKADDRS}

	err := kafkaConsumer.InitKafka()
	if err != nil {
		fmt.Println("InitKafka err: ", err)
	}
	select {}
}
