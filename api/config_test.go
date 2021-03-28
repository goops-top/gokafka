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
	conf := NewConfig()
	conf.ParserConfig("")
	for i, cluster := range Cluster {
		fmt.Println(cluster)
		if i == 0 {
			if cluster.Name == "default" {
				fmt.Println("unit test pass.")
			}
		}
	}
}

func TestInitConfig(t *testing.T) {
	// init the default config
	InitConfig()
	// paser the config filed
	conf := NewConfig()
	conf.ParserConfig("")
	fmt.Println(Cluster)
}
