package controller

import (
	"testing"
)

func TestDescribeConsumerMember(t *testing.T) {
	DescribeConsumerGroup([]string{"log-kafka-1.bgbiao.cn:9092"}, []string{"auto-commentV2-0", "active_action_user_groupid"})
}
