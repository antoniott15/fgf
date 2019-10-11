package main

type result struct {
	Name   string
	Family string
	Index  int
}

type searchResults map[int]result

func findFontByFamilyName(database []font, familyName string) (searchResults, error) {
	results := make(searchResults)
	for i, font := range database {
		// if strings.Contains(font.Family, familyName) {
		if matchNoop(familyName, font.Family) {
			results[i] = result{
				Name:   familyName,
				Family: font.Family,
				Index:  i,
			}
		}
	}
	return results, nil
}

func findFontByCategory(database []font, category string) (searchResults, error) {
	results := make(searchResults)
	for i, font := range database {
		// if strings.Contains(font.Category, category) {
		if matchNoop(category, font.Category) {
			results[i] = result{
				Name:   category,
				Family: font.Category,
				Index:  i,
			}
		}
	}
	
	return results, nil
}
