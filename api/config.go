/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:config.go
*Author:Xuebiao Xu
*Date:2020年05月07日
*Description:
*
================================================================*/
package api

// Broker地址寻址流程
// 优先从-b 参数中读取broker列表
// 其次从配置~/.goops-kafka 中读取BROKER_LIST
// 最后从BROKER_LIST环境变量中读取

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	_ "reflect"
)

type ClusterInfo struct {
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Brokers []string `json:"brokers"`
}

// var Cluster []*ClusterInfo
var Cluster []ClusterInfo

var AppName, AppVersion string

type VipConfig struct {
	ConfObj viper.Viper
}

// 初始化kafka集群配置
/*
func init() {
    conf := NewConfig()
    conf.ParserConfig("")
}
*/

func NewConfig() *VipConfig {
	return &VipConfig{
		ConfObj: *viper.New(),
	}
}

// configFile is a abs path config
func (config *VipConfig) ParserConfig(configFile string) {
	if configFile == "" {
		// os.UserHomeDir()
		config.ConfObj.AddConfigPath(os.Getenv("HOME"))
		config.ConfObj.SetConfigName(".goops-kafka")
	} else {
		config.ConfObj.AddConfigPath(filepath.Dir(configFile))
		config.ConfObj.SetConfigName(filepath.Base(configFile))
	}
	// 设置配置文件为yaml格式
	config.ConfObj.SetConfigType("yaml")
	if err := config.ConfObj.ReadInConfig(); err != nil {
		log.Warnln(err)
		InitConfig()
		return
	}
	AppName = config.ConfObj.GetString("app")
	AppVersion = config.ConfObj.GetString("version")
	// config.Get() 返回一个interface{}
	// 但由于结构本身是array，返回的是[]interface{}
	clusters := config.ConfObj.Get("spec.clusters").([]interface{})
	for i := range clusters {
		c := clusters[i].(map[interface{}]interface{})
		//fmt.Printf("%v %T\n",c["brokers"],c["brokers"])
		var brokers []string
		for _, broker := range c["brokers"].([]interface{}) {
			brokers = append(brokers, broker.(string))
		}
		/*
		   clusterInfo := &ClusterInfo{
		       Name: c["name"].(string),
		       Brokers: brokers,
		   }
		*/
		clusterInfo := ClusterInfo{
			Name: c["name"].(string),
			// Version: c["version"].(string),
			Brokers: brokers,
		}
		Cluster = append(Cluster, clusterInfo)
	}
}

var defaultConf string = `
app: gokafka
spec:
  clusters:
  - name: test-kafka
    version: V2_5_0_0
    brokers:
    - 10.0.0.1:9092
    - 10.0.0.2:9092
    - 10.0.0.3:9092
  - name: dev-kafka
    version: V1_0_0_0
    brokers:
    - 192.168.0.22:9092
    - 192.168.0.23:9092
    - 192.168.0.24:9092
`

func InitConfig() bool {
	data := []byte(defaultConf)
	filePath := fmt.Sprintf("%s/.goops-kafka", os.Getenv("HOME"))
	os.Create(filePath)
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		log.Errorf("init the config failed with :%v\n", err)
		return false
	} else {
		return true
	}
}
