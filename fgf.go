package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	instance, err := newFGFInstance()
	if err != nil {
		panic(err)
	}

	if err = instance.sync(); err != nil {
		panic(err)
	}

	fonts := []string{}
	for _, f := range instance.Database {
		fonts = append(fonts, f.Family)
	}

	searcher := func(input string, index int) bool {

		if matchNoop(strings.ToLower(input), strings.ToLower(fonts[index])) {
			return true
		}

		return false
	}

	promptSearchType := promptui.Select{
		Searcher:          searcher,
		StartInSearchMode: true,
		Size:              10,
		Label:             "Select your font family to install",
		Items:             fonts,
	}

	_, searchType, err := promptSearchType.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}

	fmt.Println(searchType)

}
