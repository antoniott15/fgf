package main

import (
	"encoding/json"
	"net/http"
)

const googleFontsAPIToken = "AIzaSyB4f0aKLlwbOrJZBevjQ-Ywnt1Z5DVKAcE"

const googleFontsAPIEndpoint = "https://www.googleapis.com/webfonts/v1/webfonts?key=" + googleFontsAPIToken

func fetchFonts() (*googleFontsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, googleFontsAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errGFBadResponse
	}

	fonts := new(googleFontsResponse)
	if err = json.NewDecoder(res.Body).Decode(fonts); err != nil {
		return nil, err
	}

	return fonts, nil
}
