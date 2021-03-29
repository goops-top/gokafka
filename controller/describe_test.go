package controller

import (
	"testing"

	"gokafka/api"
)

func TestDescribeConsumerMember(t *testing.T) {
	DescribeConsumerGroup([]string{"log-kafka-1.bgbiao.cn:9092"}, []string{"auto-commentV2-0", "active_action_user_groupid"})
}

func TestClusterApiDescribeLog(t *testing.T) {
	cluster := api.ClusterInfo{
		Name:         "test-kafka",
		Version:      "V2_5_0_0",
		Brokers:      []string{"172.16.66.50:9092"},
		Sasl:         false,
		SaslType:     "plaintext",
		SaslUser:     "MQKafkaAdmin",
		SaslPassword: "MQKafkaAdmin",
	}
	ctx := NewClusterContext(cluster)

	adminApi, _ := NewClusterApi(ctx, "admin")
	// adminApi.DescribeTopicLog()
	// adminApi.DescribeTopicListLog([]string{"__consumer_offsets"})
	// adminApi.ListConsumerGroupOffSet("im-live-dispatcher", "im-live-dispatcher-msg")
	adminApi.DescribeBroker()
}
