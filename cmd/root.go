package cmd

import (
	"fmt"
	"os"

	"github.com/hayyaun/dingdong/game"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dingdong",
	Short: "DingDong is a fun game for CLI",
	Long:  "Lovely game made for command-line for anyone, even server admins!",

	Run: func(cmd *cobra.Command, args []string) {
		game.Play()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
