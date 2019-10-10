package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const pubSpecFilename = "pubspec.yaml"
const pubSpecFilenameAlternative = "pubspec.yml"

func installFontsOnFlutterProject(flutterProjectDir string, font font) error {
	// reading pubspec.yml
	file := path.Join(flutterProjectDir, pubSpecFilename)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	completePub := map[string]interface{}{}

	err = yaml.Unmarshal(data, &completePub)
	if err != nil {
		return err
	}

	flutter, ok := completePub["flutter"].(map[interface{}]interface{})
	if !ok {
		return errFlutterInvalidPubSpec
	}

	fmt.Printf("%+v\n", flutter)

	fonts, ok := flutter["fonts"].([]map[interface{}]interface{})
	if !ok {
		return errFlutterInvalidPubSpec
	}

	fmt.Printf("%+v\n", fonts)

	assets := []map[interface{}]interface{}{}
	for _, variant := range font.Variants {
		var weight int
		if variant == "regular" || variant == "italic" {
			weight = 400
		} else {
			fixedVariant := strings.ReplaceAll(string(variant), "regular", "")
			fixedVariant = strings.ReplaceAll(string(variant), "italic", "")
			weight, _ = strconv.Atoi(fixedVariant)
		}

		fileLink := font.Files[string(variant)]
		ext := extractExtensionFromFileLink(fileLink)
		fileFontName := fmt.Sprintf("%s-%s.%s", font.Family, variant, ext)
		fileFontName = strings.ReplaceAll(fileFontName, " ", "-")

		baseDir := path.Join(flutterProjectDir, "fonts", strings.ReplaceAll(font.Family, " ", "-"))

		assetFilename := path.Join(baseDir, fileFontName)
		asset := map[interface{}]interface{}{
			"asset":  assetFilename,
			"weight": weight,
		}

		if strings.Contains(string(variant), "italic") {
			asset["style"] = "italic"
		}
		assets = append(assets, asset)
	}

	fonts = append(fonts, map[interface{}]interface{}{
		"family": font.Family,
		"fonts":  assets,
	})

	flutter["fonts"] = fonts
	completePub["flutter"] = flutter

	fmt.Printf("%+v\n", completePub)

	return nil
}
