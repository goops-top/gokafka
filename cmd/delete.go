package cmd

import (
	"fmt"
	"gokafka/api"
	"gokafka/controller"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)

	// 必须参数
	deleteCmd.PersistentFlags().StringVar(&topicName, "topic-list", "", "指定topic名称 以逗号','分割.[t1,t2]")

}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"delete"},
	Short:   "delete a topic from specified kafka cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if cluster == "" && broker == "" {
			fmt.Println("请指定kafka集群或broker地址:[--cluster or --broker]")
			cmd.Help()
			return
		}

		if topicName == "" {
			fmt.Println("请指定需要删除的topic名称:[--topic-list]")
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

		topicList := strings.Split(topicName, ",")
		isok, err := controller.DeletTopics(_broker, topicList)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(isok)
	},
}
