package tasks

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"ForumPublica/server/utils"
	"fmt"
)

//LoadMarketData updates prices using ESI API
func LoadMarketData(user models.User) error {
	mls := make([]models.MarketLocation, 0)
	db.DB.Preload("MarketItem").Preload("Character").Find(&mls)

	if len(mls) > 0 {
		updatePublicMarketStructures(mls[0].Character)

		tir := getTypesInRegions(&mls)

		tis := getTypesInStructures(&mls)

		fmt.Printf("%+v\n", tir)
		fmt.Printf("%+v\n", tis)

		//todo region-[items]

		//todo load public market citadels

		//todo char-citadel

		//todo load closed_order_volume qty

	}

	return nil
}

// ########################################################################################################################

type mapOfArrays map[int64][]int32

func getTypesInStructures(locations *[]models.MarketLocation) mapOfArrays {
	result := make(mapOfArrays)

	for _, location := range *locations {

		if location.LocationType == "structure" {
			addToMapOfArrays(result, location.LocationID, location.MarketItem.TypeID)
		}

		if location.LocationType == "solar_system" {
			structures := make([]models.Location, 0)
			db.DB.Where("solar_system_id = ? and id > ?", location.LocationID, models.StationMaxID).Find(&structures)
			for _, s := range structures {
				addToMapOfArrays(result, s.ID, location.MarketItem.TypeID)
			}
		}

	}

	return result
}

func getTypesInRegions(locations *[]models.MarketLocation) mapOfArrays {
	result := make(mapOfArrays)

	for _, location := range *locations {
		regionID := int64(0)
		if location.LocationType == "solar_system" {
			regionID = static.SolarSystems[location.LocationID].RegionID
		}
		if location.LocationType == "station" || location.LocationType == "structure" {
			l := models.Location{ID: location.LocationID}
			db.DB.Find(&l)
			regionID = l.RegionID
		}

		addToMapOfArrays(result, regionID, location.MarketItem.TypeID)
	}

	return result
}

func updatePublicMarketStructures(character *models.Character) {
	api, _ := character.GetESI()
	d, _ := api.UniverseStructures("market")

	exIDs := make([]int64, 0)

	err := db.DB.Model(&models.Location{}).Where("id in (?)", d).Pluck("id", &exIDs).Error

	if err != nil {
		fmt.Println("updatePublicMarketStructures", err)
	} else {
		for _, id := range utils.DiffInt64(d, exIDs) {
			services.AddLocation(api, id, "", 0, 0)
		}
	}

}

func addToMapOfArrays(result mapOfArrays, lid int64, tid int32) {
	temp, _ := result[lid]
	if utils.FindInt32(temp, tid) == -1 {
		temp = append(temp, tid)
	}
	result[lid] = temp
}
