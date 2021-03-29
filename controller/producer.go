package controller

import (
	"github.com/goops-top/utils/kafka"
	log "github.com/sirupsen/logrus"
)

// DEPRECATED: this function should be replaced with the method  ProducerMsgFromString
func ProducerMsgFromString(brokers []string, topic, msg string) {

	producerApi := kafka.NewProducerApi(brokers)

	defer producerApi.Close()

	log.Infof("Produce msg:%v to topic:%v\n", msg, topic)
	producerApi.PutFromString(topic, msg)

}

func (c ClusterApi) ProducerMsgFromString(topic, msg string) {
	defer c.ProducerApi.Close()

	log.Infof("Produce msg:%v to topic:%v\n", msg, topic)
	c.ProducerApi.PutFromString(topic, msg)
}
