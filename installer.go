package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func extractExtensionFromFileLink(fileLink string) string {
	ext := "ttf"
	cks := strings.Split(fileLink, ".")
	if len(cks) == 0 {
		return ext
	}

	pExt := cks[len(cks)-1]
	if len(pExt) != 3 {
		return ext
	}

	return pExt
}

func downloadFontInDir(dir string, font font) error {
	baseDir := path.Join(dir, strings.ReplaceAll(font.Family, " ", "-"))
	_ = os.MkdirAll(baseDir, 0700)

	for _, variant := range font.Variants {
		fileLink := font.Files[string(variant)]
		ext := extractExtensionFromFileLink(fileLink)
		fileFontName := fmt.Sprintf("%s-%s.%s", font.Family, variant, ext)
		fileFontName = strings.ReplaceAll(fileFontName, " ", "-")
		fmt.Println("Downloading:", fileFontName)

		res, err := http.Get(fileLink)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return errGFBadResponse
		}

		if !strings.Contains(res.Header.Get("Content-Type"), "font") {
			return errGFBadResponse
		}

		filepath := path.Join(baseDir, fileFontName)

		out, err := os.Create(filepath)
		if err != nil {
			return err
		}

		defer out.Close()

		_, err = io.Copy(out, res.Body)
		if err != nil {
			return err
		}
	}
	return nil
}
