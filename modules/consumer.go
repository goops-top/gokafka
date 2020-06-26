package modules

import (
	"context"
	"fmt"
	"sync"
	"time"

	_ "gokafka/api"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type ConsumerGroupApi struct {
	ConsumerApi sarama.ConsumerGroup
}

// init a consumer api
func NewConsumerApi(brokers []string) *ConsumerGroupApi {
	config := newConfig()
	// 指定消费者组的消费策略
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	// 指定消费组读取消息的offset[OffsetNewest,OffsetOldest]
	// config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	// 指定队列长度
	config.ChannelBufferSize = 2

	consumerGroupApi, consumerGroupApiErr := sarama.NewConsumerGroup(brokers, "gokafka", config)
	if consumerGroupApiErr != nil {
		fmt.Println("consumer group api connection failed")
		panic(consumerGroupApiErr)
	}

	return &ConsumerGroupApi{ConsumerApi: consumerGroupApi}
}

// close the consumer api
func (c *ConsumerGroupApi) Close() {
	c.ConsumerApi.Close()
}

// consumerGroupHandler
// https://pkg.go.dev/github.com/Shopify/sarama?tab=doc#ConsumerGroupHandler
// ConsumerGroupHandler是一个包含Setup，Cleanup，ConsumeClaim方法的接口
func (c *ConsumerGroupApi) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupApi) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupApi) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		part := message.Partition
		offset := message.Offset
		msg := string(message.Value)

		log.Infof("part:%v offset:%v \nmsg: %s", part, offset, msg)
		time.Sleep(time.Second)

		session.MarkMessage(message, "")
	}
	return nil
}

// consumer topic some info
func (c *ConsumerGroupApi) ConsumerMsgFromTopics(topics []string) func() {
	// ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())

	// ready := make(chan bool)
	//初始化后的消费者组api
	consumerGroupClient := c.ConsumerApi
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer func() {
			wg.Done()
		}()

		for {
			// 因为结构体p实现了Setup,Cleanup,ConsumeClaim 三个方法，所以实现了ConsumerGroupHandler接口
			err := consumerGroupClient.Consume(ctx, topics, c)
			if err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			// ready = make(chan bool)
		}
	}()
	// <-ready
	log.Infoln("Sarama consumer up and running!...")

	return func() {
		log.Info("kafka close")
		cancel()
		wg.Wait()
		log.Infoln("close the broker info.")
	}

}
