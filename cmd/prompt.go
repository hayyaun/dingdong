package cmd

import (
	"fmt"
	"os"
	"slices"

	"github.com/hayyaun/dingdong/game"
	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetSelect(pc promptContent) string {
	var result string
	var err error

	index := -1
	for index < 0 {
		prompt := promptui.Select{
			Label: pc.label,
			Items: game.Modes,
		}

		index, result, err = prompt.Run()

		if !slices.Contains(game.Modes, result) {
			index = -1 // just in case
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
