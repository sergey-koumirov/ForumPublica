package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
)

//AddLocation add location to DB
func AddLocation(api esi.ESI, id int64, name string, solarSystemID int64, regionID int64) models.Location {

	if id >= models.StationMinID && id <= models.StationMaxID && (solarSystemID == 0 || regionID == 0) {
		data, err := api.UniverseStations(id)
		if err == nil {
			name = data.Name
			solarSystemID = data.SystemID
			regionID = static.SolarSystems[data.SystemID].RegionID
		} else {
			fmt.Println("AddLocation: ", err)
			return models.Location{}
		}
	}

	if id > models.StationMaxID && (name == "" || solarSystemID == 0 || regionID == 0) {
		data, err := api.UniverseStructure(id)
		if err == nil {
			name = data.Name
			solarSystemID = data.SolarSystemID
			regionID = static.SolarSystems[data.SolarSystemID].RegionID
		} else {
			fmt.Println("AddLocation: ", err)
			return models.Location{}
		}
	}

	newLocation := models.Location{
		ID:            id,
		Name:          name,
		SolarSystemID: solarSystemID,
		RegionID:      regionID,
		LastCheckAt:   utils.NowUTCStr(),
	}
	db.DB.Create(&newLocation)

	return newLocation
}
