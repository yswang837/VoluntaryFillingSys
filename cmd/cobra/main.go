package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "一个短的帮助信息",
	Long:  `一个长的帮助信息`,
}

func init() {
	rootCmd.AddCommand(genMysqlTableCmd, completionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
