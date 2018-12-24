package services

import (
	sdem "ForumPublica/sde/models"
	"ForumPublica/sde/static"
	"regexp"
	"sort"
	"strings"
)

//SearchItems count
var SearchItems = 10

//SearchResult array
type SearchResult []sdem.ZipType

//SearchSorter for sorting
type SearchSorter struct {
	Array SearchResult
	Term  string
}

func (s SearchSorter) Len() int {
	return len(s.Array)
}
func (s SearchSorter) Swap(i, j int) {
	s.Array[i], s.Array[j] = s.Array[j], s.Array[i]
}
func (s SearchSorter) Less(i, j int) bool {
	ni := strings.Index(strings.ToLower(s.Array[i].Name), s.Term)
	nj := strings.Index(strings.ToLower(s.Array[j].Name), s.Term)
	return ni < nj || (ni == nj && s.Array[i].Name < s.Array[j].Name)
}

//SearchItemType serahc
func SearchItemType(term string, filter string) SearchResult {

	result := SearchSorter{Term: strings.ToLower(term), Array: make(SearchResult, 0)}

	var hasTerm *regexp.Regexp
	hasTerm = regexp.MustCompile("(?i)" + term)

	for _, v := range static.Types {

		isMatch := hasTerm.MatchString(v.Name)
		if filter == "blueprint" {
			_, bpoEx := static.Blueprints[v.ID]
			isMatch = bpoEx && isMatch
		}

		if isMatch {
			result.Array = append(result.Array, v)
		}
	}

	rlen := len(result.Array)
	if rlen > 0 {
		sort.Sort(result)
		if rlen > SearchItems {
			rlen = SearchItems
		}
	}

	return result.Array[:rlen]
}
