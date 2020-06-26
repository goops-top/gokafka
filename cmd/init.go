package cmd

import (
	"fmt"

	"gokafka/api"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initConfigCmd)
}

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "init the gokafka some default config.",
	Long:  `generating a gokafka config with default kafka-cluster in "~/.goops-kafka".`,
	Run: func(cmd *cobra.Command, args []string) {
		if api.InitConfig() {
			fmt.Println("init the config is ok.")
		}
	},
}
