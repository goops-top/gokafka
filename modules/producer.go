package modules

import (
	_ "gokafka/api"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type ProducerApi struct {
	ProducerSyncApi sarama.SyncProducer
}

// 初始化一个生产者api
func NewProducerApi(brokers []string) *ProducerApi {
	config := newConfig()
	// 指定生产者参数: acks = all
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 指定分区器: Random
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, clientErr := sarama.NewSyncProducer(brokers, config)
	if clientErr != nil {
		log.Infof("生产者链接失败:%v\n", clientErr)
		panic(clientErr)
	}

	return &ProducerApi{ProducerSyncApi: client}

}

//
func (p *ProducerApi) Close() {
	p.ProducerSyncApi.Close()
}

// 发送消息
func (p *ProducerApi) PutFromString(topic, msg string) bool {
	// 构造kafka消息体
	message := &sarama.ProducerMessage{}
	message.Topic = topic
	message.Value = sarama.StringEncoder(msg)

	// 像kafka topic发送消息
	_, offset, err := p.ProducerSyncApi.SendMessage(message)

	if err != nil {
		log.Infof("topic:%v send data failed with:%v\n", topic, err)
		return false
	}

	log.Infof("topic:%v send ok with offset:%v\n", topic, offset)

	return true

}
