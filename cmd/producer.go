package cmd

import (
	"fmt"

	"gokafka/api"
	"gokafka/controller"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(producerMsgCmd)

	producerMsgCmd.PersistentFlags().StringVar(&topicName, "topic", "", "指定topic名称")
	producerMsgCmd.PersistentFlags().StringVar(&msgData, "msg", "", "消息内容")

}

var producerMsgCmd = &cobra.Command{
	Use:   "producer",
	Short: "producer  a topic message data with specified kafka-cluster.",
	Long:  `This command can produce a message to specified topic  in a kafka-cluster.`,
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
		ctx := controller.NewClusterContext(*clusterInfo)
		producerApi, _ := controller.NewClusterApi(ctx, "producer")
		// 指定topic进行消费消息
		producerApi.ProducerMsgFromString(topicName, msgData)
	},
}
