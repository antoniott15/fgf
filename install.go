package main

import (
	"fmt"
	"path"
)

func installFont(workspace string, instance *fgfInstance, font font) error {
	fmt.Printf("Installing %s...\n", font.Family)

	fontsDir := path.Join(workspace, instance.fontsDir)
	err := downloadFontInDir(fontsDir, font) //instance.Database[searchIndex])
	if err != nil {
		return err
	}

	err = installFontsOnFlutterProject(workspace, font) //instance.Database[searchIndex])
	if err != nil {
		return err
	}

	fmt.Println("Done")
	return nil
}
