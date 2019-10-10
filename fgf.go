package main

import (
	"errors"
	"fmt"

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

	promptSearchType := promptui.Select{
		Searcher: func(input string, index int) bool {
			return true
		},
		StartInSearchMode: true,
		Size:              18,
		Label:             "Search By...",
		Items:             []string{"Family Name", "Category"},
	}

	_, searchType, err := promptSearchType.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}
	fmt.Println(searchType)
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("invalid input length")
		}
		return nil
	}

	promptSearchValue := promptui.Prompt{
		Label:    "Font Name",
		Validate: validate,
	}

	value, err := promptSearchValue.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}

	fmt.Println(value)

}
