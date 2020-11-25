package main

import (
	"os"

	"github.com/DemoHn/Catilina/app"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "catilina",
	Short: "Catilina 是一个简单的记账服务",
	Long:  "Catiline 是一个简单的记账服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.StartServer(configFile)
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "配置文件")
}
