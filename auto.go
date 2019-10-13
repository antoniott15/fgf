package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func executeAutoMode(workspace string, instance *fgfInstance) error {
	fonts := make([]string, len(instance.Database))
	for i, f := range instance.Database {
		fonts[i] = f.Family
	}

	promptSearchType := promptui.Select{
		Searcher: func(input string, index int) bool {
			if matchNoop(strings.ToLower(input), strings.ToLower(fonts[index])) {
				return true
			}
			return false
		},
		StartInSearchMode: true,
		Size:              10,
		Label:             "Search and select your font family to install",
		Items:             fonts,
	}

	searchIndex, _, err := promptSearchType.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	return installFont(workspace, instance, instance.Database[searchIndex])
}
