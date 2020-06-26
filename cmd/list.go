package cmd

import (
	"fmt"
	"strings"

	"gokafka/api"
	"gokafka/controller"
	_ "gokafka/modules"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.AddCommand(listClusterCmd)

	listCmd.AddCommand(listTopicCmd)
	listCmd.AddCommand(listConsumerGroupCmd)
	listTopicCmd.PersistentFlags().StringVar(&topicName, "topic-list", "", "指定topic名称,以逗号','分割.[t1,t2]")

}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list the kafka some info (cluster,topic,broker)",
	Long:    `this command can list the kafka info for cluster,broker,topic`,
}
// list the kafka cluster info 
var listClusterCmd = &cobra.Command{
	Use: "cluster",
	Short: "list the cluster info",
	Run: func(cmd *cobra.Command,args []string) {
		for _, c := range api.Cluster {
			fmt.Printf("cluster:%v version:%v connector_brokers:%v\n",c.Name,c.Version,c.Brokers)
		}

	},
}

// describe the loginfo with a topic name or all topic
var listTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "list the topic",
	Long:  `This command can list the kafka cluster someone topic base info.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" && broker == "" {
			fmt.Println("请指定kafka集群或broker地址:[--cluster or --broker]")
			cmd.Help()
			return
		}
		var _broker []string
		// 从配置文件中根据集群名获取指定的broker列表
		for _, v := range api.Cluster {
			if v.Name == cluster {
				_broker = v.Brokers
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
		}
		// topicName is a string
		// should a topicList
		if topicName == "" {
			controller.ListTopic(_broker, []string{})
		}
		topicList := strings.Split(topicName, ",")
		controller.ListTopic(_broker, topicList)

	},
}

var listConsumerGroupCmd = &cobra.Command{
	Use:   "consumer-g",
	Short: "list the consumer group",
	Long:  "This command can list the kafka cluster consumer group.",
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" && broker == "" {
			fmt.Println("请指定kafka集群或broker地址:[--cluster or --broker]")
			cmd.Help()
			return
		}
		var _broker []string
		// 从配置文件中根据集群名获取指定的broker列表
		for _, v := range api.Cluster {
			if v.Name == cluster {
				_broker = v.Brokers
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
		}

		controller.ListConsumerGroup(_broker)
	},
}
