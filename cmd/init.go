package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initConfigCmd)
}

var defaultConf string = `
app: gokafka
spec:
  clusters:
  - name: test-kafka
    version: V2_5_0_0
    brokers:
    - 172.29.203.62:9092
    - 172.29.203.106:9092
    - 172.29.203.86:9092
  - name: dev-kafka
    version: V1_0_0_0
    brokers:
    - 172.16.32.22:9092
    - 172.16.32.23:9092
    - 172.16.32.24:9092
`

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "init the gokafka some default config.",
	Long:  `generating a gokafka config with default kafka-cluster in "~/.goops-kafka".`,
	Run: func(cmd *cobra.Command, args []string) {
		data :=  []byte(defaultConf)
		if err := ioutil.WriteFile(fmt.Sprintf("%s/.goops-kafka",os.Getenv("HOME")),data,0644); err != nil {
			fmt.Printf("init the config failed with :%v\n",err)
		} else {
			fmt.Println("gokafka config init ok.")
	}
	},
}
