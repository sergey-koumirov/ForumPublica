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

	miIDs := make([]int64, len(marketItems))
	for i, el := range marketItems {
		miIDs[i] = el.ID
	}

	loadMarketVolumes(miIDs)
	marketDataMap, mdHist := loadMarketData(miIDs)
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
		bottomPrice := float64(0)
		bpoID, bpoEx := static.BpoIDByTypeID[r.TypeID]
		if bpoEx {
			bpo := static.Blueprints[bpoID]
			unitPrice = UnitPrice(bpo)
			bottomPrice = getBottomPrice(bpoID)
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
			LowestHist:  mdHist[r.ID],
			UnitPrice:   unitPrice,
			BottomPrice: bottomPrice,
			Locations:   locations,
			Stores:      stores,
		}
		result.Records = append(result.Records, temp)
	}

	return result
}

var marketVolumesSql = `           
select d.market_item_id,
 	   d.dt,
	   s.vol,
	   s.is_my
  from fp_market_data d,
       fp_market_screenshots s
  where d.id = s.market_data_id
	and d.dt > ?
	and d.market_item_id in (?)
  order by d.market_item_id, d.dt, s.price, s.id`

func loadMarketVolumes(miIDs []int64) {

	minus90d := time.Now().AddDate(0, 0, -90).Format("2006-01-02")

	rows, errRaw := db.DB.Raw(marketVolumesSql, minus90d, miIDs).Rows()
	defer rows.Close()
	if errRaw != nil {
		fmt.Println("loadMarketVolumes", errRaw)
	}

	records := make(map[int64][]models.MiMarketVolume)
	for rows.Next() {
		temp := models.MiMarketVolume{}
		db.DB.ScanRows(rows, &temp)
		records[temp.MarketItemID] = append(records[temp.MarketItemID], temp)
	}

	by_date := make(map[int64][][]models.MiMarketVolume)
	for k, vv := range records {
		index := 0
		for i := 1; i < len(vv); i++ {
			if vv[i-1].Dt != vv[i].Dt {
				by_date[k] = append(by_date[k], vv[index:i])
				index = i
			} else if i == len(vv)-1 {
				by_date[k] = append(by_date[k], vv[index:i+1])
			}
		}
	}

	compacted := make(map[int64][][]models.MiMarketVolume)
	for k, vv := range by_date {
		for _, vd := range vv {
			temp := make([]models.MiMarketVolume, 0)
			tempVol := int64(0)
			for i := 1; i < len(vd); i++ {
				tempVol = tempVol + vd[i-1].Vol
				if vd[i-1].IsMy != vd[i].IsMy {
					tempCompact := models.MiMarketVolume{
						MarketItemID: k,
						Dt:           vd[i-1].Dt,
						Vol:          tempVol,
						IsMy:         vd[i-1].IsMy,
					}
					temp = append(temp, tempCompact)
				} else if i == len(vd)-1 {
					tempCompact := models.MiMarketVolume{
						MarketItemID: k,
						Dt:           vd[i].Dt,
						Vol:          tempVol + vd[i].Vol,
						IsMy:         vd[i].IsMy,
					}
					temp = append(temp, tempCompact)
				}
			}
			compacted[k] = append(compacted[k], temp)
		}
	}

	for k, vv := range compacted {
		fmt.Println(k)
		for _, v := range vv {
			fmt.Printf("  %+v\n", v)
		}
	}

}

var bottomPriceSql = `
select sum(e.exvalue) / b.qty as price
  from fp_constructions c,
       fp_construction_bpos b,
       fp_construction_expenses e
  where c.id = b.construction_id
    and b.id = e.construction_bpo_id
    and b.type_id = ?
  group by c.id, b.id
  order by c.id desc
  limit 1`

func getBottomPrice(typeID int32) float64 {
	result := float64(0)
	rows, _ := db.DB.Raw(bottomPriceSql, typeID).Rows()
	rows.Next()
	rows.Scan(&result)
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

var sqlMdLast = `
select x.* 
  from(
    select d.*, ROW_NUMBER() OVER w AS 'row_number' 
      from fp_market_data d 
      where d.market_item_id in (?) 
      window w as (partition by d.market_item_id order by dt desc)
  ) x 
  where x.row_number=1`

var sqlMdHist = `
select mi.id, d.dt, d.sell_lowest_price as price
  from fp_market_data d,
       fp_market_items mi
  where mi.id = d.market_item_id
    and d.sell_lowest_price>0
    and mi.id in (?)
  order by mi.id, mi.type_id, d.dt, d.sell_lowest_price`

func loadMarketData(miIDs []int64) (map[int64]models.MarketData, map[int64][]models.MiHist) {
	result := make(map[int64]models.MarketData)
	hist := make(map[int64][]models.MiHist)

	if len(miIDs) > 0 {
		rows, errRaw := db.DB.Raw(sqlMdLast, miIDs).Rows()
		defer rows.Close()
		if errRaw != nil {
			fmt.Println("loadMarketData:", errRaw)
		}

		records := make([]models.MarketData, 0)
		for rows.Next() {
			temp := models.MarketData{}
			db.DB.ScanRows(rows, &temp)
			records = append(records, temp)
		}

		for _, record := range records {
			result[record.MarketItemID] = record
		}

		rows, errHist := db.DB.Raw(sqlMdHist, miIDs).Rows()
		defer rows.Close()
		if errHist != nil {
			fmt.Println("loadMarketData.errHist:", errHist)
		}

		recordsHist := make([]models.MiHist, 0)
		for rows.Next() {
			temp := models.MiHist{}
			db.DB.ScanRows(rows, &temp)
			recordsHist = append(recordsHist, temp)
		}

		for _, record := range recordsHist {
			hist[record.ID] = append(hist[record.ID], record)
		}

	}

	return result, hist
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
