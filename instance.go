package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type fgfInstance struct {
	LastFetch  time.Time `json:"lastFetch"`
	Database   []font    `json:"database"`
	FGFVersion string    `json:"fgfVersion"`
	Filename   string    `json:"filename"`
	fontsDir   string
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

func (instance *fgfInstance) loadFromFile() error {
	filepath := path.Join(os.TempDir(), instance.Filename)

	// log.Println("Loading", filepath)

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, instance)
}

func (instance *fgfInstance) saveToFile() error {
	filepath := path.Join(os.TempDir(), instance.Filename)

	// log.Println("Saving ", filepath)

	data, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, data, 0644)
}

func (instance *fgfInstance) sync() error {
	err := instance.loadFromFile()
	if err != nil {
		if err = instance.saveToFile(); err != nil {
			return err
		}
		return instance.sync()
	}

	// log.Println("Syncing")

	if !instance.LastFetch.Add(6 * time.Hour).After(time.Now()) {
		err = instance.fetchFonts()
		if err != nil {
			return err
		}

		err = instance.saveToFile()
		if err != nil {
			return err
		}
	}

	return nil
}

func newFGFInstance(opts ...fgfOptions) (*fgfInstance, error) {
	ins := &fgfInstance{
		Filename:   "fgf.json",
		FGFVersion: "0.0.2",
		fontsDir:   "/fonts",
	}

	for _, opt := range opts {
		opt.apply(ins)
	}

	return ins, nil
}
