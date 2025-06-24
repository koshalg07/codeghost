package main

import (
	"codeghost/core/agent"
	"fmt"
	"github.com/spf13/cobra"
)

var diffPath string
var logPath string

func main() {
	// Initialize Cobra command
	var rootCmd = &cobra.Command{
		Use:   "codeghost",
		Short: "A brief description of your application",
	}

	var analyzeCmd = &cobra.Command{
		Use:   "analyze",
		Short: "Analyze diff and log file",
		Run: func(cmd *cobra.Command, args []string) {
			report := agent.Analyze(diffPath, logPath)
			fmt.Println(report)
		},
	}

	analyzeCmd.Flags().StringVarP(&diffPath, "diff", "d", "", "Path to the diff file")
	analyzeCmd.Flags().StringVarP(&logPath, "log", "l", "", "Path to the log file")
	analyzeCmd.MarkFlagRequired("diff")
	analyzeCmd.MarkFlagRequired("log")

	rootCmd.AddCommand(analyzeCmd)
	rootCmd.Execute()
}
