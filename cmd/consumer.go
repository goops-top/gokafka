package cmd

import (
	"fmt"

	"gokafka/api"
	"gokafka/controller"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(consumerMsgCmd)

	consumerMsgCmd.PersistentFlags().StringVar(&topicName, "topic", "", "指定topic名称")

}

var consumerMsgCmd = &cobra.Command{
	Use:   "consumer",
	Short: "consumer  a topic message data with specified kafka-cluster.",
	Long:  `This command can consumer the message  in a kafka-cluster.`,
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
			}
		}
		if len(_broker) == 0 {
			_broker = []string{broker}
		}

		// 指定topic进行消费消息
		controller.ConsumerMsgTopics(_broker, []string{topicName})
	},
}
