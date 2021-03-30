package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"gokafka/api"
	"gokafka/controller"

	"github.com/spf13/cobra"
)

var (
	partitions string
	configs    string
)

// alter:
// alter topic --topiclist addPart --partitions --configs
// alter broker --configs
// notics: 增加点返回输出即可

func init() {
	rootCmd.AddCommand(alterCmd)

	alterCmd.AddCommand(alterTopicCmd)
	// 必须参数
	alterTopicCmd.PersistentFlags().StringVar(&topicName, "topic-list", "", "指定topic名称 以逗号','分割.[t1,t2]")

	// 可选参数
	// PersistentFlags()和Flags()
	alterTopicCmd.AddCommand(alterTopicPartitionsCmd)
	alterTopicPartitionsCmd.Flags().StringVar(&partitions, "partNum", "", "指定topic想要扩到的分区数量")

	alterTopicCmd.AddCommand(alterTopicConfigsCmd)
	alterTopicConfigsCmd.Flags().StringVar(&configs, "configs", "", "指定topic想要更新的config参数 以逗号','分割.[retention.ms:8888,unclean.leader.election:true]")

}

var alterCmd = &cobra.Command{
	Use:     "alter",
	Aliases: []string{"alter"},
	Short:   "alter the kafka some metadata (topic,broker)",
}

var alterTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "alter the topic configs and partitions with specified kafka-cluster.",
	Long:  `This command can update  the topic metadata  in a kafka-cluster.`,
}

var alterTopicPartitionsCmd = &cobra.Command{
	Use:   "addPart",
	Short: "alter the topic configs and partitions with specified kafka-cluster.",
	Long:  `This command can update  the topic metadata  in a kafka-cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" && broker == "" {
			fmt.Println("请指定kafka集群或broker地址:[--cluster or --broker]")
			cmd.Help()
			return
		}

		if topicName == "" {
			fmt.Println("请指定需要消费的topic名称:[--topic]")
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

		newPart, _ := strconv.Atoi(partitions)
		topicList := strings.Split(topicName, ",")
		for _, topic := range topicList {
			ctx := controller.NewClusterContext(*clusterInfo)
			adminApi, _ := controller.NewClusterApi(ctx, "admin")
			isok, err := adminApi.AlterTopicPartitionsNum(topic, int32(newPart))
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Println(isok)
		}
	},
}

var alterTopicConfigsCmd = &cobra.Command{
	Use:   "configs",
	Short: "alter the topic configs and partitions with specified kafka-cluster.",
	Long:  `This command can update  the topic metadata  in a kafka-cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" && broker == "" {
			fmt.Println("请指定kafka集群或broker地址:[--cluster or --broker]")
			cmd.Help()
			return
		}

		if topicName == "" {
			fmt.Println("请指定需要消费的topic名称:[--topic]")
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

		topicList := strings.Split(topicName, ",")
		// 构造一个map[string]string{} 类型的配置文件
		topicConfig := make(map[string]string)
		if len(configs) != 0 {
			configEntry := strings.Split(configs, ",")
			for _, v := range configEntry {
				k := strings.Split(v, ":")[0]
				v := strings.Split(v, ":")[1]
				topicConfig[k] = v
			}
		} else {
			fmt.Println("请使用--configs参数指定需要更新topic的具体参数，负责会将全部topic级别参数覆盖.")
			cmd.Help()
			return
		}

		for _, topic := range topicList {
			ctx := controller.NewClusterContext(*clusterInfo)
			adminApi, _ := controller.NewClusterApi(ctx, "admin")
			isok, err := adminApi.AlterTopicConfigs(topic, topicConfig)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Println(isok)
		}
	},
}
