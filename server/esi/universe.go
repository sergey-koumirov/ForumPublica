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
	Name          string         `json:"name"`
	SolarSystemID int64          `json:"solar_system_id"`
	Position      PositionRecord `json:"position"`
}

//UniverseStructure get structure info
func (esi *ESI) UniverseStructure(structureID int64) (UniverseStructuresRecord, error) {
	url := fmt.Sprintf("%s/universe/structures/%d/", ESIRootURL, structureID)

	result := UniverseStructuresRecord{}
	_, _, err := auth("GET", esi.AccessToken, url, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}

//PositionRecord model
type PositionRecord struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

//StationRecord model
type StationRecord struct {
	MaxDockableShipVolume    float64        `json:"max_dockable_ship_volume"`
	Name                     string         `json:"name"`
	OfficeRentalCost         float64        `json:"office_rental_cost"`
	Owner                    int64          `json:"owner"`
	Position                 PositionRecord `json:"position"`
	RaceID                   int64          `json:"race_id"`
	ReprocessingEfficiency   float64        `json:"reprocessing_efficiency"`
	ReprocessingStationsTake float64        `json:"reprocessing_stations_take"`
	Services                 []string       `json:"services"`
	StationID                int64          `json:"station_id"`
	SystemID                 int64          `json:"system_id"`
	TypeID                   int64          `json:"type_id"`
}

//UniverseStations get station info
func (esi *ESI) UniverseStations(stationID int64) (StationRecord, error) {
	url := fmt.Sprintf("%s/universe/stations/%d/", ESIRootURL, stationID)

	result := StationRecord{}
	_, _, err := get(url, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}

//UniverseStructures get public structures ids (market, manufacturing_basic)
func (esi *ESI) UniverseStructures(filter string) ([]int64, error) {
	url := fmt.Sprintf("%s/universe/structures/", ESIRootURL)
	if filter != "" {
		url = url + "?filter=" + filter
	}

	result := make([]int64, 0)
	_, _, err := get(url, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
