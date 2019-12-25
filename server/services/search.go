package services

import (
	sdem "ForumPublica/sde/models"
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
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

//SearchLocationRecord result record type
type SearchLocationRecord struct {
	ID   int64
	Name string
	Type string
}

//SearchLocationResult array
type SearchLocationResult []SearchLocationRecord

//SearchLocationSorter for sorting
type SearchLocationSorter struct {
	Array SearchLocationResult
	Term  string
}

func (s SearchLocationSorter) Len() int {
	return len(s.Array)
}
func (s SearchLocationSorter) Swap(i, j int) {
	s.Array[i], s.Array[j] = s.Array[j], s.Array[i]
}
func (s SearchLocationSorter) Less(i, j int) bool {
	ni := strings.Index(strings.ToLower(s.Array[i].Name), s.Term)
	nj := strings.Index(strings.ToLower(s.Array[j].Name), s.Term)
	return ni < nj || (ni == nj && s.Array[i].Name < s.Array[j].Name)
}

//SearchLocation search
func SearchLocation(userID int64, charID int64, term string, filter string) (SearchLocationResult, error) {

	result := SearchLocationSorter{Term: strings.ToLower(term), Array: make(SearchLocationResult, 0)}

	char := models.Character{ID: charID}
	errDb := db.DB.Where("user_id=?", userID).Find(&char).Error
	if errDb != nil {
		fmt.Println("SearchLocation.errDb:", errDb)
		return result.Array, errDb
	}

	api, errESI := char.GetESI()
	if errESI != nil {
		fmt.Println("SearchLocation.errESI:", errESI)
		return result.Array, errESI
	}

	var filters []string
	if filter == "stations" {
		filters = []string{"structure", "station"}
	} else {
		filters = []string{"solar_system", "structure", "station"}
	}

	ids, apiErr := api.CharactersSearch(charID, filters, term, false)
	if apiErr != nil {
		fmt.Println("SearchLocation: [api.CharactersSearch]", apiErr)
	}

	if len(ids["solar_system"]) > 0 {
		addSolarSystems(ids["solar_system"], &result)
	}

	if len(ids["station"]) > 0 {
		addStations(api, ids["station"], &result)
	}

	if len(ids["structure"]) > 0 {
		addStructures(api, ids["structure"], &result)
	}

	sort.Sort(result)

	return result.Array, nil
}

func addSolarSystems(ids []int64, result *SearchLocationSorter) {
	for _, id := range ids {
		result.Array = append(
			result.Array,
			SearchLocationRecord{
				ID:   id,
				Name: static.SolarSystems[id].Name,
				Type: "solar_system",
			},
		)
	}
}

func addStations(api esi.ESI, ids []int64, result *SearchLocationSorter) {
	exLocations := []models.Location{}
	err := db.DB.Where("id in (?)", ids).Find(&exLocations).Error
	if err != nil {
		fmt.Println("addStations:", err)
	}

	notFoundIds := make([]int64, 0)
	for _, id := range ids {
		notFound := true
		var founded models.Location
		for _, l := range exLocations {
			if id == l.ID {
				notFound = false
				founded = l
			}
		}
		if notFound {
			notFoundIds = append(notFoundIds, id)
		} else {
			result.Array = append(
				result.Array,
				SearchLocationRecord{
					ID:   founded.ID,
					Name: founded.Name,
					Type: "station",
				},
			)
		}
	}

	var newLocations []esi.UniverseNameRecord
	if len(notFoundIds) > 0 {
		var e error
		newLocations, e = api.UniverseNames(notFoundIds)
		if err != nil {
			fmt.Println("addStations: api.UniverseNames", e)
		} else {
			for _, l := range newLocations {
				api := esi.ESI{}
				AddLocation(api, l.ID, l.Name, 0, 0)
				result.Array = append(
					result.Array,
					SearchLocationRecord{
						ID:   l.ID,
						Name: l.Name,
						Type: "station",
					},
				)
			}
		}
	}

}

func addStructures(api esi.ESI, ids []int64, result *SearchLocationSorter) {
	exLocations := []models.Location{}
	err := db.DB.Where("id in (?)", ids).Find(&exLocations).Error
	if err != nil {
		fmt.Println("addStations:", err)
	}

	for _, id := range ids {
		notFound := true
		var founded models.Location
		for _, l := range exLocations {
			if id == l.ID {
				notFound = false
				founded = l
			}
		}
		if notFound || founded.LastCheckAt == "" || utils.StrToMinut(founded.LastCheckAt) > 24*60 {
			if notFound {
				founded = AddLocation(api, id, "", 0, 0)
			} else {
				db.DB.Model(&founded).Update("last_check_at", utils.NowUTCStr())
			}
		}

		result.Array = append(
			result.Array,
			SearchLocationRecord{
				ID:   founded.ID,
				Name: founded.Name,
				Type: "structure",
			},
		)
	}

}
