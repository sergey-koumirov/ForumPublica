package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
	"time"
)

//MarketItemsList list
func MarketItemsList(userID int64, page int64) models.MiList {

	marketItems, total := loadMarketItems(userID, page)

	marketDataMap := loadMarketData(marketItems)

	d90 := loadD90(marketItems)

	result := models.MiList{
		Page:       page,
		Total:      total,
		PerPage:    MIPerPage,
		Characters: CharsByUserID(userID),
	}

	result.Records = make([]models.MiRecord, 0)

	for _, r := range marketItems {

		locations := loadLocations(r)

		stores, storeVol := loadStores(r)

		md, _ := marketDataMap[r.ID]

		unitPrice := float64(0)
		bpoID, bpoEx := static.BpoIDByTypeID[r.TypeID]
		if bpoEx {
			bpo := static.Blueprints[bpoID]
			unitPrice = UnitPrice(bpo)
		}

		temp := models.MiRecord{
			ModelID:     r.ID,
			TypeID:      r.TypeID,
			TypeName:    static.Types[r.TypeID].Name,
			MyPrice:     md.MyLowestPrice,
			MyVol:       md.MyVol,
			StoreVol:    storeVol,
			D90Vol:      d90[r.ID].Total,
			D90Data:     d90[r.ID].R,
			LowestPrice: md.SellLowestPrice,
			UnitPrice:   unitPrice,
			Locations:   locations,
			Stores:      stores,
		}
		result.Records = append(result.Records, temp)
	}

	return result
}

var loadD90Sql = `
select mio.id, d.date as d, ifnull(x.q,0) as q
  from fp_market_items mio
         inner join sys_dates d
         left join (
			  select mi.id,
					 substring(t.dt, 1, 10) as d,
					 sum(quantity) as q
				from esi_transactions t
					   inner join esi_locations l on t.location_id = l.id
					   inner join fp_market_items mi on mi.id in (
						select mix.id
							from fp_market_items mix,
								fp_market_locations milx
							where mix.type_id = t.type_id
							and milx.market_item_id = mix.id
							and milx.esi_character_id = t.esi_character_id
							and (
								milx.location_type = 'system' and l.solar_system_id = milx.location_id
							or
							milx.location_type in ('station','structure') and t.location_id = milx.location_id
							)
						)
				where t.dt > '%s'
				  and t.esi_character_id in (?)
				  and t.type_id in (?)
				  and t.is_buy = 0
				group by mi.id, substring(t.dt, 1, 10)
        ) x on x.d = d.date and mio.id = x.id 
  where d.date between '%s' and '%s'
    and mio.id in (?)
  order by mio.id, d.date`

func loadD90(marketItems []models.MarketItem) map[int64]models.Tr90dSummary {

	miIds := make([]int64, 0)
	charIdsMap := make(map[int64]int32)
	typeIdsMap := make(map[int32]int32)
	for _, mi := range marketItems {
		miIds = append(miIds, mi.ID)
		typeIdsMap[mi.TypeID] = 1
		for _, loc := range mi.Locations {
			charIdsMap[loc.EsiCharacterID] = 1
		}
	}

	charIds := make([]int64, 0)
	for k := range charIdsMap {
		charIds = append(charIds, k)
	}

	typeIds := make([]int32, 0)
	for k := range typeIdsMap {
		typeIds = append(typeIds, k)
	}

	minus90dFull := time.Now().AddDate(0, 0, -90).Format("2006-01-02 15:04:05")
	minus90d := time.Now().AddDate(0, 0, -90).Format("2006-01-02")
	minus0d := time.Now().Format("2006-01-02")

	rawSql := fmt.Sprintf(loadD90Sql, minus90dFull, minus90d, minus0d)

	rows, errRaw := db.DB.Raw(rawSql, charIds, typeIds, miIds).Rows()
	defer rows.Close()

	if errRaw != nil {
		fmt.Println("loadD90", errRaw)
	}

	records := make(map[int64]models.Tr90dSummary)
	for rows.Next() {
		temp := models.Tr90d{}
		db.DB.ScanRows(rows, &temp)

		v, _ := records[temp.Id]
		v.R = append(v.R, temp)
		v.Total = v.Total + temp.Q
		records[temp.Id] = v
	}

	return records
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

		rows, _ := db.DB.Raw(sql, ids).Rows()
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

func loadStores(r models.MarketItem) ([]models.MiStore, int64) {
	stores := make([]models.MiStore, 0)
	storeVol := int64(0)
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
		storeVol = storeVol + s.StoreQty
	}
	return stores, storeVol
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
