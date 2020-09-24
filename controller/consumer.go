package controller

import (
	"github.com/goops-top/utils/kafka"
	log "github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"
)

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
