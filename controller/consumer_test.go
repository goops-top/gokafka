package controller

import (
	"gokafka/api"
	"testing"
)

func TestConsumerMsg(t *testing.T) {
	ConsumerMsgTopics([]string{"172.29.203.62:9092"}, []string{"test-push"}, "consumer-group-x", "latest")
}

func TestClusterApiConsumer(t *testing.T) {
	cluster := api.ClusterInfo{
		Name:         "test-kafka",
		Version:      "V2_5_0_0",
		Brokers:      []string{"172.29.202.78:9092"},
		Sasl:         true,
		SaslType:     "plaintext",
		SaslUser:     "MQKafkaAdmin",
		SaslPassword: "MQKafkaAdmin",
	}
	ctx := NewClusterContext(cluster)

	producerApi, _ := NewClusterApi(ctx, "consumer")
	producerApi.ConsumerMsgTopics([]string{"GoOps"}, "consumer-group-x", "latest")
}
