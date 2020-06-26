package controller

import (
	_ "gokafka/modules"

	log "github.com/sirupsen/logrus"
	"github.com/goops-top/utils/kafka"

	"os"
	"os/signal"
	"syscall"
)

func ConsumerMsgTopics(brokers, topics []string) {

	// consumerApi := modules.NewConsumerApi(brokers)
	consumerApi := kafka.NewConsumerApi(brokers)

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
