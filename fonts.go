package main

type variant string

type subset string

type font struct {
	// Kind         string            `json:"kind"`
	Family       string            `json:"family"`
	Category     string            `json:"category"`
	Variants     []variant         `json:"variants"`
	Subsets      []subset          `json:"subsets"`
	Version      string            `json:"version"`
	LastModified string            `json:"lastModified"`
	Files        map[string]string `json:"files"`
}

type googleFontsResponse struct {
	Kind  string `json:"kind"`
	Items []font `json:"items"`
}

func getFontByFamilyName(fonts []font, familyName string) (font, error) {
	for _, f := range fonts {
		if f.Family == familyName {
			return f, nil
		}
	}

	return font{}, errFamilyFontNotFound
}
