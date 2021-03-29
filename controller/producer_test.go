package controller

import (
	"testing"

	"gokafka/api"
)

func TestClusterApiProducerMsgFromString(t *testing.T) {
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

	producerApi, _ := NewClusterApi(ctx, "producer")
	producerApi.ProducerMsgFromString("GoOps", "BGBiao1111")
}
