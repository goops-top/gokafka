package controller

import (
	"github.com/goops-top/utils/kafka"
	log "github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"
)

// DEPRECATED: this function should be replaced with the method  DescribeConsumerGroup

func ConsumerMsgTopics(brokers, topics []string, groupName, offset string) {
	// consumerApi := modules.NewConsumerApi(brokers)
	// consumerApi := kafka.NewConsumerApi(brokers)
	consumerApi := kafka.NewConsumerApi(brokers, groupName, offset)

	defer consumerApi.Close()

	c := consumerApi.ConsumerMsgFromTopics(topics)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	c()

}

func (c ClusterApi) ConsumerMsgTopics(topics []string, groupName, offset string) {
	defer c.ConsumerApi.Close()
	consumer := c.ConsumerApi.ConsumerMsgFromTopics(topics)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	for i := 0; i < 5; i++ {
		go consumer()
	}

}
