package cmd

import (
	"fmt"
	"encoding/json"
	"os"

	"gokafka/utils"

	"github.com/spf13/cobra"
)

var (
	output		string
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().StringVar(&output, "out", "", "指定输出格式[--out json]")
}

func parserVersion() {
    v := utils.Get()
	if output == "json" {
		versionMarshal,err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	
		fmt.Println(string(versionMarshal))
	} else {
		fmt.Printf("Version: %s\nGitBranch: %s\nCommitId: %s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s\n",v.Version,v.GitBranch,v.GitCommit,v.BuildDate,v.GoVersion,v.Platform)
	}



}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version info of goops-kafka tool",
	Long:  `This command can be used get the version number of goops-kafka tool`,
	Run: func(cmd *cobra.Command, args []string) {
		parserVersion()
	},
}
