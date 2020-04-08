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
func LoadMarketData() error {
	fmt.Println("LoadMarketData started", time.Now().Format("2006-01-02 15:04:05"))
	loadMarketData(0, false)
	fmt.Println("LoadMarketData finished", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}

func LoadMarketDataForUser(userID int64) {
	loadMarketData(userID, true)
}

// ########################################################################################################################

func loadMarketData(userID int64, technical bool) {
	mrkLocs := make([]models.MarketLocation, 0)
	scopeLocs := db.DB.Preload("MarketItem").Preload("Character")
	if userID > 0 {
		scopeLocs = scopeLocs.Where("fp_market_locations.market_item_id in (select fp_market_items.id from fp_market_items where fp_market_items.user_id=?)", userID)
	}
	scopeLocs.Find(&mrkLocs)

	if len(mrkLocs) > 0 {
		updatePublicMarketStructures(mrkLocs[0].Character)

		dt := time.Now().UTC().Format("2006-01-02 15:04:05")

		orders := make(map[int64]esi.MarketsOrdersArray)

		loadOrdersFromRegions(&mrkLocs, orders)
		loadOrdersFromStructures(&mrkLocs, orders)

		marketItems := make([]models.MarketItem, 0)
		scopeMI := db.DB.Preload("Locations.Character").Preload("Stores.Character")
		if userID > 0 {
			scopeMI = scopeMI.Where("fp_market_items.user_id=?", userID)
		}
		scopeMI.Find(&marketItems)

		charOrdersCache := preloadCharOrders(marketItems)
		charItemsCache := preloadCharItems(marketItems)

		for _, mi := range marketItems {
			createMarketData(mi, dt, orders[mi.ID], charOrdersCache, technical)
			updateMarketStores(mi, charItemsCache)
			updateExpirations(mi, charOrdersCache)
		}
	}
}

func updateExpirations(mi models.MarketItem, charOrdersCache map[int64]esi.CharacterMarketOrders) {

	for _, loc := range mi.Locations {

		orders := charOrdersCache[loc.EsiCharacterID]

		earliest := ""

		for _, order := range orders {
			temp := models.Location{ID: order.LocationID}
			db.DB.Find(&temp)
			if order.TypeID == mi.TypeID && (loc.LocationType == "solar_system" && temp.SolarSystemID == loc.LocationID || order.LocationID == loc.LocationID) {
				test := order.Issued.AddDate(0, 0, order.Duration).Format("2006-01-02 15:04:05")
				if earliest == "" || test < earliest {
					earliest = test
				}
			}
		}
		db.DB.Model(&loc).Update("expiration", earliest)
	}

}

type itemsByChar map[int64]esi.ItemsByLocation

func preloadCharItems(mis []models.MarketItem) itemsByChar {
	chars := make(map[int64]models.Character)
	for _, mi := range mis {
		for _, store := range mi.Stores {
			chars[store.EsiCharacterID] = *store.Character
		}
	}

	result := make(itemsByChar)

	for _, char := range chars {
		api, _ := char.GetESI()
		d, err := api.CharactersAssetsTypeIdsByLocationID(char.ID)
		if err != nil {
			fmt.Println("preloadCharItems: ", err, char)
		} else {
			result[char.ID] = d
		}
	}

	return result
}

func updateMarketStores(mi models.MarketItem, items itemsByChar) {
	for _, store := range mi.Stores {
		charItems, exChar := items[store.EsiCharacterID]
		if exChar {
			locItems, exLoc := charItems[store.LocationID]
			if exLoc {
				qty, exType := locItems[mi.TypeID]
				if exType {
					store.StoreQty = qty
				} else {
					store.StoreQty = 0
				}
				err := db.DB.Save(&store).Error
				if err != nil {
					fmt.Println("updateMarketStores: ", err)
				}
			}
		}
	}
}

type mapOfArrays map[int64][]int32

func preloadCharOrders(mis []models.MarketItem) map[int64]esi.CharacterMarketOrders {
	charOrdersCache := make(map[int64]esi.CharacterMarketOrders)

	chars := make(map[int64]models.Character)
	for _, mi := range mis {
		for _, location := range mi.Locations {
			chars[location.Character.ID] = *location.Character
		}
	}

	for cid, char := range chars {
		api, _ := char.GetESI()
		r, _ := api.CharactersOrders(cid)
		charOrdersCache[char.ID] = make(esi.CharacterMarketOrders, len(r.R))
		for i, order := range r.R {
			charOrdersCache[char.ID][i] = order
		}
	}

	return charOrdersCache
}

func createMarketData(mi models.MarketItem, dt string, orders esi.MarketsOrdersArray, charOrdersCache map[int64]esi.CharacterMarketOrders, technical bool) {

	var (
		sellVol         int64
		sellLowestPrice float64
		buyVol          int64
		buyHighestPrice float64
		myVol           int64
		myLowestPrice   float64
	)

	charIds := make([]int64, 0)
	for _, location := range mi.Locations {
		if utils.FindInt64(charIds, location.Character.ID) == -1 {
			charIds = append(charIds, location.Character.ID)
		}
	}

	for _, order := range orders {
		if !order.IsBuyOrder {
			sellVol = sellVol + order.VolumeRemain
			if sellLowestPrice > order.Price || sellLowestPrice == 0 {
				sellLowestPrice = order.Price
			}

			isMy := false
			for _, cid := range charIds {
				for _, r := range charOrdersCache[cid] {
					if r.OrderID == order.OrderID {
						isMy = true
						break
					}
				}
			}
			if isMy {
				myVol = myVol + order.VolumeRemain
				if myLowestPrice > order.Price || myLowestPrice == 0 {
					myLowestPrice = order.Price
				}
			}

		} else {
			buyVol = buyVol + order.VolumeRemain
			if buyHighestPrice < order.Price {
				buyHighestPrice = order.Price
			}
		}
	}

	dataPoint := models.MarketData{
		MarketItemID:    mi.ID,
		Dt:              dt,
		SellVol:         sellVol,
		SellLowestPrice: sellLowestPrice,
		BuyVol:          buyVol,
		BuyHighestPrice: buyHighestPrice,
		MyVol:           myVol,
		MyLowestPrice:   myLowestPrice,
		Technical:       technical,
	}
	errDb := db.DB.Create(&dataPoint).Error
	if errDb != nil {
		fmt.Println("createMarketData: ", errDb)
	}

	if !technical {
		for _, order := range orders {
			if !order.IsBuyOrder {

				isMy := false
				for _, cid := range charIds {
					for _, r := range charOrdersCache[cid] {
						if r.OrderID == order.OrderID {
							isMy = true
							break
						}
					}
				}

				screenPoint := models.MarketScreenshot{
					MarketDataID: dataPoint.ID,
					Vol:          order.VolumeRemain,
					Price:        order.Price,
					IsMy:         isMy,
				}
				db.DB.Create(&screenPoint)
			}
		}
	}

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
							added := false
							for _, t := range temp {
								if t.OrderID == order.OrderID {
									added = true
									break
								}
							}
							if !added {
								temp = append(temp, order)
								result[l.MarketItem.ID] = temp
							}

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
