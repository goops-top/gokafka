package cmd

import (
	"fmt"
	"strings"

	"gokafka/api"
	"gokafka/controller"

	"github.com/spf13/cobra"
)

var (
	topicName     string
	partNum       int32
	replicaFactor int16
	configEntries string
	retentionTime int32
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVar(&topicName, "topic", "", "指定topic名称")
	createCmd.PersistentFlags().Int32Var(&partNum, "partition", 3, "指定分区数量[默认分区:3]")
	createCmd.PersistentFlags().Int16Var(&replicaFactor, "replicas", 3, "指定副本数量[默认副本:3]")
	// createCmd.PersistentFlags().StringToStringVar(&topicConfig, "config", map[string]string{"":""},"指定topic相关的配置")
	createCmd.PersistentFlags().Int32Var(&retentionTime, "retention", 48, "指定保留的时间，单位:h(小时)")
	createCmd.PersistentFlags().StringVar(&configEntries, "topicConfig", "", "指定topic的配置参数进行创建，配置项以逗号分割(retention.ms:172800000,unclean.leader.election.enable:true)")

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
		topicConfig := make(map[string]string)
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

		if len(configEntries) != 0 {
			configEntry := strings.Split(configEntries, ",")
			for _, v := range configEntry {
				k := strings.Split(v, ":")[0]
				v := strings.Split(v, ":")[1]
				topicConfig[k] = v
			}
		}
		// topicConfig = map[string]string{"retention.ms": fmt.Sprintf("%v", retentionTime*60*60*1000)}
		topicConfig["retention.ms"] = fmt.Sprintf("%v", retentionTime*60*60*1000)

		// 创建topic
		ctx := controller.NewClusterContext(*clusterInfo)
		adminApi, _ := controller.NewClusterApi(ctx, "admin")
		adminApi.CreateTopic(topicName, partNum, replicaFactor, topicConfig)
	},
}
