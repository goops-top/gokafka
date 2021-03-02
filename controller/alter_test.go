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
	isok, err := AlterTopicPartitionsNum([]string{"172.16.32.22:9092"}, "heleitest", 2)

	fmt.Println(isok, err)
}
