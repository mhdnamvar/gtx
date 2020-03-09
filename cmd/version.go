package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show GTX version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GTX version 1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	serveCmd.AddCommand(versionCmd)
}
