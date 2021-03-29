/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:describe.go
*Author:Xuebiao Xu
*Date:2020年05月18日
*Description:
*
================================================================*/
package controller

import (
	"github.com/goops-top/utils/kafka"

	"fmt"
)

// DEPRECATED: this function should be replaced with the method  DescribeConsumerGroup
func CreateTopic(brokers []string, topic string, part int32, replicas int16, topicConfig map[string]string) {
	// kfkAdmin := modules.NewClusterAdmin(brokers)
	kfkAdmin := kafka.NewClusterAdmin(brokers)
	defer kfkAdmin.Close()
	isok, err := kfkAdmin.CreateCustomTopic(topic, part, replicas, topicConfig)
	if isok {
		fmt.Println(isok)
	} else {
		fmt.Println(err)
	}
}

func (c ClusterApi) CreateTopic(topic string, part int32, replicas int16, topicConfig map[string]string) {
	defer c.AdminApi.Close()
	isok, err := c.AdminApi.CreateCustomTopic(topic, part, replicas, topicConfig)
	if isok {
		fmt.Println(isok)
	} else {
		fmt.Println(err)
	}
}
