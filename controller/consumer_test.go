package controller

import (
	"testing"
)

func TestConsumerMsg(t *testing.T) {
	ConsumerMsgTopics([]string{"172.29.203.62:9092"}, []string{"test-push"}, "consumer-group-x", "latest")
}
