package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows gtx version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gtx version 1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	runCmd.AddCommand(versionCmd)
}
