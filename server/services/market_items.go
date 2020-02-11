package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
)

//MarketItemsList list
func MarketItemsList(userID int64, page int64) models.MiList {

	marketItems, total := loadMarketItems(userID, page)

	marketDataMap := loadMarketData(marketItems)
	fmt.Println(marketDataMap)

	result := models.MiList{
		Page:       page,
		Total:      total,
		PerPage:    MIPerPage,
		Characters: CharsByUserID(userID),
	}

	result.Records = make([]models.MiRecord, 0)

	for _, r := range marketItems {

		locations := loadLocations(r)

		stores := loadStores(r)

		md, _ := marketDataMap[r.ID]

		temp := models.MiRecord{
			ModelID:     r.ID,
			TypeID:      r.TypeID,
			TypeName:    static.Types[r.TypeID].Name,
			MyPrice:     md.MyLowestPrice,
			LowestPrice: md.SellLowestPrice,
			Locations:   locations,
			Stores:      stores,
		}
		result.Records = append(result.Records, temp)
	}

	return result
}

func loadMarketData(marketItems []models.MarketItem) map[int64]models.MarketData {
	result := make(map[int64]models.MarketData)

	if len(marketItems) > 0 {

		ids := make([]int64, len(marketItems))
		for i, el := range marketItems {
			ids[i] = el.ID
		}

		sql := "select x.* from(" +
			"select d.*, ROW_NUMBER() OVER w AS 'row_number' from fp_market_data d where d.market_item_id in (?) " +
			"window w as (partition by d.market_item_id order by dt desc)) x where x.row_number=1"

		rows, _ := db.DB.Model(&models.MarketData{}).Raw(sql, ids).Rows()
		defer rows.Close()

		records := make([]models.MarketData, 0)

		for rows.Next() {
			temp := models.MarketData{}
			db.DB.ScanRows(rows, &temp)
			records = append(records, temp)
		}

		for _, record := range records {
			result[record.MarketItemID] = record
		}
	}

	return result
}

func loadStores(r models.MarketItem) []models.MiStore {
	stores := make([]models.MiStore, 0)
	for _, s := range r.Stores {
		stores = append(
			stores,
			models.MiStore{
				ID:            s.ID,
				Type:          s.LocationType,
				Name:          LocationName(s.LocationID, s.LocationType),
				CharacterName: s.Character.Name,
				Qty:           s.StoreQty,
			},
		)
	}
	return stores
}

func loadLocations(r models.MarketItem) []models.MiLocation {
	locations := make([]models.MiLocation, 0)

	for _, l := range r.Locations {
		locations = append(
			locations,
			models.MiLocation{
				ID:            l.ID,
				Type:          l.LocationType,
				Name:          LocationName(l.LocationID, l.LocationType),
				CharacterName: l.Character.Name,
			},
		)
	}
	return locations
}

func loadMarketItems(userID int64, page int64) ([]models.MarketItem, int64) {
	marketItems := make([]models.MarketItem, 0)
	var total int64

	scope := db.DB.Where("user_id = ?", userID)
	scope.Model(&models.MarketItem{}).Count(&total)
	scope.Preload("Locations.Character").
		Preload("Stores.Character").
		Order("id desc").
		Limit(MIPerPage).
		Offset((page - 1) * MIPerPage).
		Find(&marketItems)
	return marketItems, total
}

//MarketItemsCreate create
func MarketItemsCreate(userID int64, params map[string]int32) {
	typeID, ex := params["TypeID"]
	if ex {
		temp := models.MarketItem{
			UserID: userID,
			TypeID: typeID,
		}
		db.DB.Create(&temp)
	}
}

//MarketItemsDelete delete
func MarketItemsDelete(userID int64, miID int64) {
	mi := models.MarketItem{}
	errSel := db.DB.Where("id = ? and user_id = ?", miID, userID).First(&mi).Error
	if errSel != nil {
		return
	}
	mi.Delete()
}
