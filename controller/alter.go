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
