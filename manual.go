package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

const commandAdd = "add"
const commandDownload = "download"

func executeManualMode(workspace string, instance *fgfInstance, command string, params ...string) error {
	if command == commandAdd {
		if len(params) != 1 {
			return errNumberOfParamsNotValid
		}
		fontFamily := params[0]

		result, err := findFontByFamilyName(instance.Database, fontFamily)
		if err != nil {
			return err
		}

		if len(result) == 0 {
			fmt.Println("font family not found, try with other")
			return nil // TODO: Check that
		}

		if len(result) == 1 {
			fmt.Printf("Font family found: %s\n", result[0].Family)
			fontToInstall := instance.Database[result[0].Index]
			return installFont(workspace, instance, fontToInstall)
		}
		fontsToSelect := []font{}

		for _, res := range result {
			f, err := getFontByFamilyName(instance.Database, res.Family)
			if err != nil {
				return err
			}
			fontsToSelect = append(fontsToSelect, f)
		}

		fontsFamilyNames := make([]string, len(fontsToSelect))
		for i := range fontsToSelect {
			fontsFamilyNames[i] = fontsToSelect[i].Family
		}
		// selecting...
		promptSelectFont := promptui.Select{
			StartInSearchMode: true,
			Size:              10,
			Label:             "Select Your Font",
			Items:             fontsFamilyNames,
		}

		indexSelected, _, err := promptSelectFont.Run()
		if err != nil {
			return err
		}

		return installFont(workspace, instance, fontsToSelect[indexSelected])

	} else if command == commandDownload {
		if len(params) != 1 {
			return errNumberOfParamsNotValid
		}
		return errNotImplemented

	} else {
		return errInvalidCommand
	}

	return nil
}
