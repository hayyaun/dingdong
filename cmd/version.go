package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of DingDong",
	Long:  `All software has versions. This is DingDong's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ding Dong v0.1")
	},
}
