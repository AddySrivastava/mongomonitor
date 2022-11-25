package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mongomonitor",
	Short: "Periodically monitor and store the MongoDB deployment",
	Long:  `This CLI will run jobs to pull logs from your MongoDB deployment and store them either in a remote	location or in S3, which can be used to visualize and find some anomalies over time.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
