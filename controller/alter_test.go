/*
Copyright 2020 BGBiao Ltd. All rights reserved.
@File   : alter_test.go
@Time   : 2021/03/02 15:16:04
@Update : 2021/03/02 15:16:04
@Author : BGBiao
@Version: 1.0
@Contact: weichaungxxb@qq.com
@Desc   : None
*/
package controller

import (
	"fmt"
	"testing"
)

func TestAlterTopicPartitionsNum(t *testing.T) {
	isok, err := AlterTopicPartitionsNum([]string{"172.16.32.22:9092"}, "heleitest", 3)

	fmt.Println(isok, err)
}

func TestAlterTopicConfigs(t *testing.T) {
	isok, err := AlterTopicConfigs([]string{"172.16.32.22:9092"}, "heleitest", map[string]string{"retention.ms": "43200000", "unclean.leader.election.enable": "true"})
	fmt.Println(isok, err)
}

func TestDeleteTopic(t *testing.T) {
	isok, err := DeleteTopic([]string{"172.16.32.22:9092"}, "heleitest")
	fmt.Println(isok, err)
}
