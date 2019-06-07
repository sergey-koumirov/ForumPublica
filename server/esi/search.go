package esi

import (
	"fmt"
	"strings"
)

//CharactersSearchResult model
type CharactersSearchResult map[string][]int64

//CharactersSearch get names by ids
func (esi *ESI) CharactersSearch(charID int64, categories []string, term string, strict bool) (CharactersSearchResult, error) {

	url := fmt.Sprintf(
		"%s/characters/%d/search/?categories=%s&search=%s&%t",
		ESIRootURL,
		charID,
		strings.Join(categories, ","),
		term,
		strict,
	)
	var result CharactersSearchResult
	_, _, err := auth("GET", esi.AccessToken, url, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
