package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"time"
)

type fgfInstance struct {
	LastFetch  time.Time `json:"lastFetcg"`
	Database   []font    `json:"database"`
	FGFVersion string    `json:"fgfVersion"`
}

func (instance *fgfInstance) fetchFonts() error {
	resp, err := fetchFonts()
	if err != nil {
		return err
	}

	if resp.Kind != "webfonts#webfontList" {
		return errGFBadResponse
	}

	instance.Database = resp.Items
	instance.LastFetch = time.Now()

	return nil
}

func (instance *fgfInstance) loadFromFile(filename string) error {
	dir, err := ioutil.TempDir(fgfTempDir, "temp")
	if err != nil {
		return err
	}

	filepath := path.Join(dir, filename)

	log.Println(filepath)

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, instance)
}
