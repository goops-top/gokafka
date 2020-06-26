package controller

import (
	// "gokafka/modules"

	log "github.com/sirupsen/logrus"
	"github.com/goops-top/utils/kafka"
)

func ProducerMsgFromString(brokers []string, topic, msg string) {

	// producerApi := modules.NewProducerApi(brokers)
	producerApi := kafka.NewProducerApi(brokers)

	defer producerApi.Close()

	log.Infof("Produce msg:%v to topic:%v\n", msg, topic)
	producerApi.PutFromString(topic, msg)

}
