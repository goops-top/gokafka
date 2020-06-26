package cmd

import (
	"fmt"
	"gokafka/api"
	"gokafka/controller"

	"github.com/spf13/cobra"
)

var (
	topicName     string
	partNum       int32
	replicaFactor int16
	topicConfig   map[string]string
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVar(&topicName, "topic", "", "指定topic名称")
	createCmd.PersistentFlags().Int32Var(&partNum, "partition", 3, "指定分区数量[默认分区:3]")
	createCmd.PersistentFlags().Int16Var(&replicaFactor, "replicas", 3, "指定副本数量[默认副本:3]")
	// createCmd.PersistentFlags().StringToStringVar(&topicConfig, "config", map[string]string{"":""},"指定topic相关的配置")

}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create the kafka topic with some base params in specify kafka-cluster.",
	Long:  `this command can use some params create a topic in a kafka-cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" && broker == "" {
			fmt.Println("请指定kafka集群或broker地址:[--cluster or --broker]")
			cmd.Help()
			return
		}

		if topicName == "" {
			fmt.Println("请指定需要创建的topic名称:[--topic]")
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

		// 创建topic
		controller.CreateTopic(_broker, topicName, partNum, replicaFactor, topicConfig)
	},
}
