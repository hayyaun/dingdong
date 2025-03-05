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
		fmt.Println("\nWelcome to DingDong Game!")
		fmt.Println("\nPlease select difficulty of the game.")
		fmt.Println()
		game.Mode = promptGetSelect(promptContent{
			"Please select an option.",
			"How hard should it be?",
		})
		game.Play()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
