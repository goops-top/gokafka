/*
Copyright 2020 BGBiao Ltd. All rights reserved.
@File   : alter.go
@Time   : 2021/03/02 15:13:44
@Update : 2021/03/02 15:13:44
@Author : BGBiao
@Version: 1.0
@Contact: weichaungxxb@qq.com
@Desc   : alter.go 用来定义一些相关的修改的操作，主要用于修改topic,broker相关的参数
*/
package controller

import (
	"github.com/goops-top/utils/kafka"
	xerrors "github.com/pkg/errors"
)

// 修改topic的分区数量，kafka中为保证数据的可靠性是不允许缩减分区的
// 比较多的分区可以提高整个消息的并发消费，但同时也可能对一定的网络时延有一定的影响
// DEPRECATED: this function should be replaced with the method  AlterTopicPartitionsNum
func AlterTopicPartitionsNum(brokers []string, topicName string, partNum int32) (bool, error) {

	admin := kafka.NewClusterAdmin(brokers)
	defer admin.Close()

	var resagnment [][]int32
	isOK, err := admin.AddPartitions(topicName, partNum, resagnment, false)

	if err != nil {
		return isOK, xerrors.Wrapf(err, "failed to alter the topic partitions with topic:%s part:%d ", topicName, partNum)
	}
	return isOK, err
}

func (c ClusterApi) AlterTopicPartitionsNum(topicName string, partNum int32) (bool, error) {
	defer c.AdminApi.Close()

	var resagnment [][]int32
	isOK, err := c.AdminApi.AddPartitions(topicName, partNum, resagnment, false)

	if err != nil {
		return isOK, xerrors.Wrapf(err, "failed to alter the topic partitions with topic:%s part:%d ", topicName, partNum)
	}
	return isOK, err
}

// 修改topic相关的参数
// config: map[string]string{"retention.ms":"43200000","unclean.leader.election.enable","true"}
// DEPRECATED: this function should be replaced with the method  AlterTopicConfigs
func AlterTopicConfigs(brokers []string, topicName string, configs map[string]string) (bool, error) {
	admin := kafka.NewClusterAdmin(brokers)
	defer admin.Close()

	isOK, err := admin.UpdateTopicConfig(topicName, configs, false)

	if err != nil {
		return isOK, xerrors.Wrapf(err, "failed to update the topic config with topic:%s configs:%v", topicName, configs)
	}

	return isOK, err

}

func (c ClusterApi) AlterTopicConfigs(topicName string, configs map[string]string) (bool, error) {
	defer c.AdminApi.Close()

	isOK, err := c.AdminApi.UpdateTopicConfig(topicName, configs, false)

	if err != nil {
		return isOK, xerrors.Wrapf(err, "failed to update the topic config with topic:%s configs:%v", topicName, configs)
	}

	return isOK, err

}

// delete a kafka topic
// DEPRECATED: this function should be replaced with the method  DeleteTopic
func DeleteTopic(brokers []string, topicName string) (bool, error) {
	admin := kafka.NewClusterAdmin(brokers)

	defer admin.Close()

	isOK, err := admin.DeleteTopic(topicName)

	if err != nil {
		return isOK, xerrors.Wrapf(err, "failed to delete the topic with topicName :%s", topicName)
	}
	return isOK, err

}

func (c ClusterApi) DeleteTopic(topicName string) (bool, error) {
	defer c.AdminApi.Close()
	isOK, err := c.AdminApi.DeleteTopic(topicName)

	if err != nil {
		return isOK, xerrors.Wrapf(err, "failed to delete the topic with topicName :%s", topicName)
	}
	return isOK, err
}

// delete some kafka topics
// DEPRECATED: this function should be replaced with the method  DeletTopics
func DeletTopics(brokers, topicsName []string) (bool, error) {

	var deleteErrors error
	var deleteList []bool
	topicListLenth := len(topicsName)

	if topicListLenth == 0 {
		return false, xerrors.Wrapf(xerrors.New("topicName list is null"), "failed to delete the topics with topiclist:%v", topicsName)
	}
	admin := kafka.NewClusterAdmin(brokers)
	defer admin.Close()

	for _, v := range topicsName {
		isOK, err := admin.DeleteTopic(v)
		if isOK {
			deleteList = append(deleteList, true)
		}
		if err != nil {
			deleteErrors = xerrors.Wrapf(err, "failed to delete the topic with: %v\n", v)
		}

	}

	if len(deleteList) != topicListLenth {
		return false, deleteErrors
	}

	return true, nil
}

func (c ClusterApi) DeletTopics(topicsName []string) (bool, error) {
	var deleteErrors error
	var deleteList []bool
	topicListLenth := len(topicsName)

	if topicListLenth == 0 {
		return false, xerrors.Wrapf(xerrors.New("topicName list is null"), "failed to delete the topics with topiclist:%v", topicsName)
	}

	defer c.AdminApi.Close()

	for _, v := range topicsName {
		isOK, err := c.AdminApi.DeleteTopic(v)
		if isOK {
			deleteList = append(deleteList, true)
		}
		if err != nil {
			deleteErrors = xerrors.Wrapf(err, "failed to delete the topic with: %v\n", v)
		}

	}

	if len(deleteList) != topicListLenth {
		return false, deleteErrors
	}

	return true, nil

}
