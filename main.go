/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:main.go
*Author:Xuebiao Xu
*Date:2020年05月07日
*Description:
*
================================================================*/
package main

import (
	"fmt"
	"gokafka/cmd"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.ClientID = "test-xxb"
	config.ChannelBufferSize = 256
	config.Version = sarama.V2_5_0_0

	if confErr := config.Validate(); confErr != nil {
		fmt.Println("kafka config 解析错误，check please.")
	}
	cmd.Execute()
}
