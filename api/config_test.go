/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:config_test.go
*Author:Xuebiao Xu
*Date:2020年05月07日
*Description:
*
================================================================*/
package api

import (
	"fmt"
	"os"
	"testing"
)

func TestParserConfig(t *testing.T) {
	fmt.Println(os.Getenv("HOME"))
	for i, cluster := range Cluster {
		if i == 0 {
			if cluster.Name == "default" {
				fmt.Println("unit test pass.")
			}
		}
	}
}
