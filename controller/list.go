/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:describe.go
*Author:Xuebiao Xu
*Date:2020年05月18日
*Description: list the kafka concepts some infos.
*clusterList
*topicList
*brokerList
================================================================*/

package controller

import (
	"github.com/goops-top/utils/kafka"

	"fmt"
	"strings"
)

// list the topic
// DEPRECATED: this function should be replaced with the method  ListTopic
func ListTopic(brokers, topics []string) {
	kfkAdmin := kafka.NewClusterAdmin(brokers)
	defer kfkAdmin.Close()
	topicsInfos, err := kfkAdmin.ListTopicsInfo(topics)
	if err != nil {
		panic(err)
	}
	for _, topic := range topicsInfos {
		confEntry := ""
		for confName, confValue := range topic.ConfigEntries {
			confEntry = confEntry + fmt.Sprintf("%v:%v,", confName, *confValue)
		}
		fmt.Printf("Topic:%v\tPartNum:%v\tReplicas:%v\tConfig:%v\n", topic.Name, topic.PartitionNum, topic.Replication, confEntry)
		for k, v := range topic.ReplicaAssignment {
			fmt.Printf("Topic-Part:%v-%v\tReplicaAssign:%v\n", topic.Name, k, v)
		}
	}
}

func (c ClusterApi) ListTopic(topics []string) {
	// 关闭早了，可能导致后面的listTopicsInfo 异常
	// 最好是级联进行控制相关的关闭和退出
	// defer c.AdminApi.Close()

	topicsInfos, err := c.AdminApi.ListTopicsInfo(topics)
	if err != nil {
		panic(err)
	}
	for _, topic := range topicsInfos {
		confEntry := ""
		for confName, confValue := range topic.ConfigEntries {
			confEntry = confEntry + fmt.Sprintf("%v:%v,", confName, *confValue)
		}
		fmt.Printf("Topic:%v\tPartNum:%v\tReplicas:%v\tConfig:%v\n", topic.Name, topic.PartitionNum, topic.Replication, confEntry)
		for k, v := range topic.ReplicaAssignment {
			fmt.Printf("Topic-Part:%v-%v\tReplicaAssign:%v\n", topic.Name, k, v)
		}
	}

}

// list the consumer group
// DEPRECATED: this function should be replaced with the method  ListConsumerGroup
func ListConsumerGroup(brokers []string) {
	kfkAdmin := kafka.NewClusterAdmin(brokers)
	defer kfkAdmin.Close()

	consumerGroups, consumerListErr := kfkAdmin.ListConsumerGroup()
	if consumerListErr != nil {
		fmt.Printf("获取消费组信息失败:%v\n", consumerListErr)
		panic(consumerListErr)
	}

	fmt.Printf("%v\n", strings.Join(consumerGroups, "\n"))

}

func (c ClusterApi) ListConsumerGroup() {
	defer c.AdminApi.Close()

	consumerGroups, consumerListErr := c.AdminApi.ListConsumerGroup()
	if consumerListErr != nil {
		fmt.Printf("获取消费组信息失败:%v\n", consumerListErr)
		panic(consumerListErr)
	}

	fmt.Printf("%v\n", strings.Join(consumerGroups, "\n"))

}
