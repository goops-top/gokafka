package cmd

import (
	"gokafka/api"
	"gokafka/cmd/templates"

	"github.com/spf13/cobra"
)

var (
	// kafka 相关参数
	cluster    string // 指定集群名称
	broker     string // 指定broker地址"ip:9092"
	groupNames string // 指定分组列表"group1,group2"
	group      string // 指定分组名称"消费组"
	msgData    string // 指定生产者的生产消息
	version    string // 查看版本信息

	// 配置文件
	cfgFile string
	// cmd主命令
	rootCmd = &cobra.Command{
		Use:   "gokafka some tools",
		Short: "goops-kafka: A common no jvm kafka operate tools",
		// Long:  `goops-kafka: A kafka tools that has no jvm with golang.`,
		Long: templates.LongDesc(`
    goops-kafka: A kafka tools with golang that can operate the kafka for describe,create,update and so on.

    Note: Without the jvm,so you must be specify the [--broker or --cluster and the Value must be in --config entry.]
    `),
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 持久标记允许参数作用到全局
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "指定配置文件(default is $HOME/.goops-kafka)")
	rootCmd.PersistentFlags().StringVar(&broker, "broker", "", "指定broker地址")
	rootCmd.PersistentFlags().StringVar(&cluster, "cluster", "", "指定集群")

	cobra.OnInitialize(initConfig)

}

func initConfig() {
	conf := api.NewConfig()
	if cfgFile != "" {
		conf.ParserConfig(cfgFile)
	} else {
		conf.ParserConfig("")
	}

}
