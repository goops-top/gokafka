package cmd

import (
	"fmt"

	"gokafka/api"
	"gokafka/controller"

	"github.com/spf13/cobra"
)

var (
	offset        string
	consumerGroup string
)

func init() {
	rootCmd.AddCommand(consumerMsgCmd)

	consumerMsgCmd.PersistentFlags().StringVar(&topicName, "topic", "", "指定topic名称")

	// 可选参数
	// PersistentFlags()和Flags()
	consumerMsgCmd.Flags().StringVar(&consumerGroup, "group", "", "指定消费者组")
	consumerMsgCmd.Flags().StringVar(&offset, "offset", "", "指定消费的位置[latest|earliest]")

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

		ctx := controller.NewClusterContext(*clusterInfo)
		consumerApi, _ := controller.NewClusterApi(ctx, "consumer")
		// 指定topic进行消费消息
		consumerApi.ConsumerMsgTopics([]string{topicName}, consumerGroup, offset)
	},
}
