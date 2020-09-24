/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:describe.go
*Author:Xuebiao Xu
*Date:2020年05月18日
*Description: describe the kafka concepts some infos.
*describeTopicLog
*describeTopic
*describeBroker
*
================================================================*/
package controller

import (
	"github.com/goops-top/utils/kafka"

	"fmt"
)

// describe topic logdir
func DescribeTopicLog(brokers []string) {
	// kfkAdmin := kafka.NewClusterAdmin(brokers)
	kfkAdmin := kafka.NewClusterAdmin(brokers)
	defer kfkAdmin.Close()
	// 默认获取全部topic的日志
	for _, data := range kfkAdmin.GetLogFromTopic("") {
		fmt.Println(data.BrokerIp)
		for _, logData := range data.LogDatas {
			fmt.Printf("logdir:%v\n", logData.Path)
			fmt.Printf("topic-part\t\tlog-size(M)\t\toffset-lag\n")
			fmt.Printf("----------\t\t-----------\t\t----------\n")
			for _, v := range logData.LogInfo {
				fmt.Printf("%v\t\t%v\t\t%v\t\t\n", v.TopicPart, v.LogSize, v.OffsetLag)
			}
		}
	}

}

func DescribeTopicListLog(brokers []string, topics []string) {
	kfkAdmin := kafka.NewClusterAdmin(brokers)
	defer kfkAdmin.Close()
	for _, topicsDatas := range kfkAdmin.GetLogFromTopics(topics) {
		fmt.Printf("topic:%v\n", topicsDatas.Name)
		for _, data := range topicsDatas.LogData {
			fmt.Println(data.BrokerIp)
			for _, logData := range data.LogDatas {
				fmt.Printf("logdir:%v\n", logData.Path)
				fmt.Printf("topic-part\t\tlog-size(M)\t\toffset-lag\n")
				fmt.Printf("----------\t\t-----------\t\t----------\n")
				for _, v := range logData.LogInfo {
					fmt.Printf("%v\t\t%v\t\t%v\t\t\n", v.TopicPart, v.LogSize, v.OffsetLag)
				}
			}
		}
	}
}

// list the consumer-group offset infos
// 其实也需要把part的leader副本的logsize拿到
func ListConsumerGroupOffSet(brokers []string, group, topic string) {
	kfkAdmin := kafka.NewClusterAdmin(brokers)
	defer kfkAdmin.Close()

	topicPartOffSet, partOffSetErr := kfkAdmin.ListConsumerGroupOffSet(group, topic)

	if partOffSetErr != nil {
		panic(partOffSetErr)
	}

	for _, v := range topicPartOffSet {
		fmt.Printf("topic-part:%v log-offsize:%v\n", v.TopicPart, v.OffSet)
	}
}

// describe the topic
func DescribeTopic(brokers, topics []string) {
	kfkAdmin := kafka.NewClusterAdmin(brokers)
	defer kfkAdmin.Close()
	topicMetas, metaErr := kfkAdmin.DescribeTopics(topics)
	if metaErr != nil {
		panic(metaErr)
	}

	for _, metadata := range topicMetas {
		//fmt.Printf("Topic:%v\n",metadata.Name)
		fmt.Printf("Topic-Part:%v-%v\tLeader:%v\tReplicas:%v\tISR:%v\tOfflineRep:%v\n",
			metadata.Name, metadata.PartId,
			metadata.PartLeader,
			metadata.PartReplicas,
			metadata.PartIsr,
			metadata.PartOfflineReplicas)
	}
}

// describe the brokers
func DescribeBroker(brokers []string) {
	admin := kafka.NewClusterAdmin(brokers)
	controllerId, brokerIds, brokerInfos := admin.GetBrokerIdList()
	fmt.Println("controller:", controllerId)
	fmt.Println("brokers num:", len(brokerIds))
	fmt.Println("broker list:", brokerIds)
	for _, v := range brokerInfos {
		fmt.Printf("id:%v\t\t broker:%v\t\n", v.BrokerId, v.BrokerIp)
	}
}

// describe the consumer-group infos
func DescribeConsumerGroup(brokers, consumerGroups []string) {
	admin := kafka.NewClusterAdmin(brokers)
	consumerMembers, _ := admin.DescribeConsumerGroup(consumerGroups)
	for _, consumer := range consumerMembers {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("consumer-group:%v consumer-state:%v\n", consumer.GroupID, consumer.State)
		if consumer.State == "Empty" {
			break
		}
		fmt.Printf("consumer-id\t\t\t\t\tconsumer-ip\t\t\ttopic-list\n")
		for _, consumerMember := range consumer.ClientInfo {
			fmt.Printf("%v\t%v\t\t%v\n", consumerMember.ClientID, consumerMember.ClientIP, consumerMember.TopicList)
		}
	}
}
