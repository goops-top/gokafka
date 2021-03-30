package cmd

import (
	"fmt"
	"strings"

	"gokafka/api"
	"gokafka/controller"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(describeCmd)

	// describe 子命令下增加loginfo 和topic 的子命令
	describeCmd.AddCommand(descLogInfoCmd)
	describeCmd.AddCommand(descTopicCmd)

	describeCmd.AddCommand(descBrokerCmd)
	describeCmd.AddCommand(descConsumerInfoCmd)
	describeCmd.AddCommand(descConsumerGroupOffSetCmd)

	// describeCmd.PersistentFlags().StringVar(&topicName, "topic-list", "", "指定topic名称,以逗号','分割.[t1,t2]")
	descLogInfoCmd.PersistentFlags().StringVar(&topicName, "topic-list", "", "指定topic名称,以逗号','分割.[t1,t2]")
	descTopicCmd.PersistentFlags().StringVar(&topicName, "topic-list", "", "指定topic名称,以逗号','分割.[t1,t2]")

	descConsumerInfoCmd.PersistentFlags().StringVar(&groupNames, "consumer-group-list", "", "指定消费组列表,以逗号','分割.[cg1,cg2]")

	descConsumerGroupOffSetCmd.PersistentFlags().StringVar(&topicName, "topic", "", "指定topic名称")
	descConsumerGroupOffSetCmd.PersistentFlags().StringVar(&group, "group", "", "指定consumer-group名称")

}

var describeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "describe the kafka some info (cluster,topic,broker,loginfo)",
	Long:    `this command can describe the kafka info for cluster,broker,topic,loginfo`,
}

// describe the loginfo with a topic name or all topic
var descLogInfoCmd = &cobra.Command{
	Use:   "loginfo",
	Short: "describe the loginfo",
	Long:  `This command can describe the kafka cluster loginfo or someone topic loginfo`,
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
				// 构造集群相关的基本函数
				clusterInfo.Brokers = _broker
				clusterInfo.Sasl = v.Sasl
				clusterInfo.SaslType = v.SaslType
				clusterInfo.SaslUser = v.SaslUser
				clusterInfo.SaslPassword = v.SaslPassword
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
		}

		// 使用初始化后的配置创建对应的api
		ctx := controller.NewClusterContext(*clusterInfo)
		adminApi, _ := controller.NewClusterApi(ctx, "admin")

		if topicName == "" {
			adminApi.DescribeTopicLog()

		} else {
			topicList := strings.Split(topicName, ",")
			adminApi.DescribeTopicListLog(topicList)
		}

	},
}

// describe topic info
var descTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "describe the topic info",
	Long:  `This command can describe the kafka cluster topic infos`,
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
				// 构造集群相关的基本函数
				clusterInfo.Brokers = _broker
				clusterInfo.Sasl = v.Sasl
				clusterInfo.SaslType = v.SaslType
				clusterInfo.SaslUser = v.SaslUser
				clusterInfo.SaslPassword = v.SaslPassword
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
		}
		// 使用初始化后的配置创建对应的api
		ctx := controller.NewClusterContext(*clusterInfo)
		adminApi, _ := controller.NewClusterApi(ctx, "admin")
		if topicName == "" {
			adminApi.DescribeTopic([]string{})

		} else {
			topicList := strings.Split(topicName, ",")
			adminApi.DescribeTopic(topicList)
		}
	},
}

// describe brokers
var descBrokerCmd = &cobra.Command{
	Use:   "broker",
	Short: "describe the brokers in specified cluster",
	Long:  "This command can describe the kafka cluster broker info (controller id)",
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
				// 构造集群相关的基本函数
				clusterInfo.Brokers = _broker
				clusterInfo.Sasl = v.Sasl
				clusterInfo.SaslType = v.SaslType
				clusterInfo.SaslUser = v.SaslUser
				clusterInfo.SaslPassword = v.SaslPassword
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
		}
		// 使用初始化后的配置创建对应的api
		ctx := controller.NewClusterContext(*clusterInfo)
		adminApi, _ := controller.NewClusterApi(ctx, "admin")
		// describe the broker info
		adminApi.DescribeBroker()
	},
}

// describe consumer-group
var descConsumerInfoCmd = &cobra.Command{
	Use:   "consumer-g",
	Short: "describe the consumer-group some info",
	Long:  "This command can describe the kafka consumer-group info with specified consumer-group",
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
				// 构造集群相关的基本函数
				clusterInfo.Brokers = _broker
				clusterInfo.Sasl = v.Sasl
				clusterInfo.SaslType = v.SaslType
				clusterInfo.SaslUser = v.SaslUser
				clusterInfo.SaslPassword = v.SaslPassword
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
		}

		if groupNames == "" {
			fmt.Println("请指定消费者组查看消费者组详情:--consumer-group-list")
			cmd.Help()
		} else {
			groupList := strings.Split(groupNames, ",")
			// 使用初始化后的配置创建对应的api
			ctx := controller.NewClusterContext(*clusterInfo)
			adminApi, _ := controller.NewClusterApi(ctx, "admin")
			adminApi.DescribeConsumerGroup(groupList)
		}

	},
}

var descConsumerGroupOffSetCmd = &cobra.Command{
	Use:   "consumer-group-offset",
	Short: "describe the consumer group's consumer offset with specified topic.",
	Long:  "This command can describe the consumer group offset with specified topic",
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" && broker == "" {
			fmt.Println("请指定kafka集群或broker地址:[--cluster or --broker]")
			cmd.Help()
			return
		}

		if group == "" && topicName == "" {
			fmt.Println("请指定消费组和topic查看消费的offset情况:[--group and --topic]")
			cmd.Help()
			return
		}
		var _broker []string
		// 从配置文件中根据集群名获取指定的broker列表
		for _, v := range api.Cluster {
			if v.Name == cluster {
				_broker = v.Brokers
				// 构造集群相关的基本函数
				clusterInfo.Brokers = _broker
				clusterInfo.Sasl = v.Sasl
				clusterInfo.SaslType = v.SaslType
				clusterInfo.SaslUser = v.SaslUser
				clusterInfo.SaslPassword = v.SaslPassword
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
			clusterInfo.Brokers = _broker
		}
		ctx := controller.NewClusterContext(*clusterInfo)
		adminApi, _ := controller.NewClusterApi(ctx, "admin")
		adminApi.ListConsumerGroupOffSet(group, topicName)

	},
}
