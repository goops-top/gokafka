package controller

import (
	"github.com/goops-top/utils/kafka"
	log "github.com/sirupsen/logrus"
)

func ProducerMsgFromString(brokers []string, topic, msg string) {

	producerApi := kafka.NewProducerApi(brokers)

	defer producerApi.Close()

	log.Infof("Produce msg:%v to topic:%v\n", msg, topic)
	producerApi.PutFromString(topic, msg)

}
