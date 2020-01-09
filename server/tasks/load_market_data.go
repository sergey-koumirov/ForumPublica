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

		fmt.Println(tir)

		//todo region-[items]

		//todo load public market citadels

		//todo char-citadel

		//todo load closed_order_volume qty

	}

	return nil
}

type typesInRegions map[int64][]int32

func getTypesInRegions(locations *[]models.MarketLocation) typesInRegions {
	result := make(typesInRegions)

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

		temp, _ := result[regionID]

		if utils.FindInt32(temp, location.MarketItem.TypeID) == -1 {
			temp = append(temp, location.MarketItem.TypeID)
		}

		result[regionID] = temp
	}

	return result
}

func updatePublicMarketStructures(character models.Character) {
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
