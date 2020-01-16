package tasks

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"ForumPublica/server/utils"
	"fmt"
	"time"
)

//LoadMarketData updates prices using ESI API
func LoadMarketData(user models.User) error {
	mls := make([]models.MarketLocation, 0)
	db.DB.Preload("MarketItem").Preload("Character").Find(&mls)

	if len(mls) > 0 {
		updatePublicMarketStructures(mls[0].Character)

		dt := time.Now().UTC().Format("2006-01-02 15:04:05")

		orders := make(map[int64]esi.MarketsOrdersArray)

		loadOrdersFromRegions(&mls, orders)
		loadOrdersFromStructures(&mls, orders)

		marketItems := make([]models.MarketItem, 0)
		db.DB.Preload("Locations.Character").Find(&marketItems)

		charOrdersCache := make(map[int64][]int64)

		for _, mi := range marketItems {

			preloadCharOrders(charOrdersCache, mi)

			createMarketData(mi, dt, orders[mi.ID])
		}

		for miid, oo := range orders {
			fmt.Printf("%d:\n", miid)
			for _, o := range oo {
				fmt.Printf("    %+v\n", o)
			}
		}

		//todo region-[items]

		//todo load public market citadels

		//todo load closed_order_volume qty

	}

	return nil
}

// ########################################################################################################################

type mapOfArrays map[int64][]int32

func preloadCharOrders(charOrdersCache map[int64][]int64, mi models.MarketItem) {
	temp := charOrdersCache[mi.ID]
	chars := make(map[int64]models.Character)
	for _, location := range mi.Locations {
		chars[location.Character.ID] = *location.Character
	}

	for cid, char := range chars {
		api, _ := char.GetESI()
		r, _ := api.CharactersOrders(cid)

		for _, order := range r.R {
			temp = append(temp, order.OrderID)
			// todo cache by char
		}

	}

	charOrdersCache[mi.ID] = temp
}

func createMarketData(mi models.MarketItem, dt string, orders esi.MarketsOrdersArray) {

	var (
		sellVol         int64
		sellLowestPrice float64
		buyVol          int64
		buyHighestPrice float64
	)

	for _, order := range orders {
		if !order.IsBuyOrder {
			sellVol = sellVol + order.VolumeRemain
			if sellLowestPrice > order.Price || sellLowestPrice == 0 {
				sellLowestPrice = order.Price
			}
		} else {
			buyVol = buyVol + order.VolumeRemain
			if buyHighestPrice < order.Price {
				buyHighestPrice = order.Price
			}
		}
	}

	data := models.MarketData{
		MarketItemID:    mi.ID,
		Dt:              dt,
		SellVol:         sellVol,
		SellLowestPrice: sellLowestPrice,
		BuyVol:          buyVol,
		BuyHighestPrice: buyHighestPrice,
		LowerVol:        0,
		MyVol:           0,
		GreaterVol:      0,
	}
	db.DB.Create(&data)
}

func loadOrdersFromStructures(locations *[]models.MarketLocation, result map[int64]esi.MarketsOrdersArray) {
	tis := getTypesInStructures(locations)
	character := (*locations)[0].Character
	for sid := range tis {
		api, _ := character.GetESI()
		orders, _ := api.MarketsStructuresAll(sid)
		for _, order := range orders {
			for _, l := range *locations {
				if order.TypeID == l.MarketItem.TypeID && (l.LocationType == "structure" && order.LocationID == l.LocationID) {
					temp := result[l.MarketItem.ID]
					temp = append(temp, order)
					result[l.MarketItem.ID] = temp
				}
			}
		}
	}
}

func loadOrdersFromRegions(locations *[]models.MarketLocation, result map[int64]esi.MarketsOrdersArray) {
	tir := getTypesInRegions(locations)

	for rid, tids := range tir {
		for _, tid := range tids {
			api := esi.ESI{}
			orders, err := api.MarketsOrdersAll(rid, tid, "all")
			if err != nil {
				fmt.Println("loadOrdersFromRegions: ", err)
			} else {
				for _, order := range orders {
					for _, l := range *locations {
						if order.TypeID == l.MarketItem.TypeID && (l.LocationType == "solar_system" && order.SystemID == l.LocationID ||
							l.LocationType == "station" && order.LocationID == l.LocationID) {
							temp := result[l.MarketItem.ID]
							temp = append(temp, order)
							result[l.MarketItem.ID] = temp
						}
					}
				}
			}
		}
	}
}

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
