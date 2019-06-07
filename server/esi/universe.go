package esi

import (
	"fmt"
	"strconv"
	"strings"
)

//UniverseNameRecord model
type UniverseNameRecord struct {
	Category string `json:"category"`
	ID       int64  `json:"id"`
	Name     string `json:"name"`
}

//UniverseNames id to names
func (esi *ESI) UniverseNames(ids []int64) ([]UniverseNameRecord, error) {
	url := fmt.Sprintf("%s/universe/names", ESIRootURL)

	qIds := make([]string, len(ids))
	for i, id := range ids {
		qIds[i] = strconv.FormatInt(id, 10)
	}

	result := make([]UniverseNameRecord, 0)
	_, _, err := post(url, fmt.Sprintf("[%s]", strings.Join(qIds, ",")), &result)

	if err != nil {
		return result, err
	}

	return result, nil
}

//UniverseStructuresRecord model
type UniverseStructuresRecord struct {
	Name          string `json:"name"`
	SolarSystemID int64  `json:"solar_system_id"`
}

//UniverseStructures get structure info
func (esi *ESI) UniverseStructures(structureID int64) (UniverseStructuresRecord, error) {
	url := fmt.Sprintf("%s/universe/structures/%d/", ESIRootURL, structureID)

	result := UniverseStructuresRecord{}
	_, _, err := auth("GET", esi.AccessToken, url, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}
