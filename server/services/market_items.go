package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
	"sort"
	"time"
)

//MarketItemsList list
func MarketItemsList(userID int64, page int64) models.MiList {

	marketItems, total := loadMarketItems(userID, page)

	miIDs := make([]int64, len(marketItems))
	for i, el := range marketItems {
		miIDs[i] = el.ID
	}

	volumes := loadMarketVolumes(miIDs)
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

		d90Data := make([]models.AnyArr, 0)
		for _, t := range d90[r.ID].R {
			d90Data = append(d90Data, models.AnyArr{t.Q})
		}

		temp := models.MiRecord{
			ModelID:     r.ID,
			TypeID:      r.TypeID,
			TypeName:    static.Types[r.TypeID].Name,
			MyPrice:     md.MyLowestPrice,
			MyVol:       md.MyVol,
			StoreVol:    storeVol,
			D90Vol:      d90[r.ID].Total,
			D90Data:     d90Data,
			LowestPrice: md.SellLowestPrice,
			LowestHist:  mdHist[r.ID],
			UnitPrice:   unitPrice,
			BottomPrice: bottomPrice,
			Locations:   locations,
			Stores:      stores,
			VolumeHist:  volumes[r.ID],
		}

		warningsForMarketItem(&temp)

		result.Records = append(result.Records, temp)
	}

	return result
}

func warningsForMarketItem(mir *models.MiRecord) {
	result := make(map[string]bool)

	result["MarketQty"] = false
	if mir.MyVol == 0 {
		result["MarketQty"] = true
	}

	result["StoreQty"] = false
	if mir.MyVol+mir.StoreVol < mir.D90Vol/3 {
		result["StoreQty"] = true
	}

	result["LowestPrice"] = false
	if mir.MyPrice > mir.LowestPrice || mir.MyPrice == 0 {
		result["LowestPrice"] = true
	}

	for i, l := range mir.Locations {
		if l.Expiration != "" && utils.DbStrToMinut(l.Expiration) > -7*24*60 {
			mir.Locations[i].OrderExpired = true
		}
	}

	mir.Warnings = result
}

var marketVolumesSql = `
select market_item_id, dt, vol, is_my as m
  from vol30d 
  where market_item_id in (?)
    and dt > ?
  order by group_num`

func loadMarketVolumes(miIDs []int64) map[int64]models.AnyArr {

	start := time.Now().UTC()
	minus30d := start.AddDate(0, 0, -30).Format("2006-01-02")

	byDate := make(map[int64]map[string][]models.MiVolume, len(miIDs))
	rows, errRaw := db.DB.Raw(marketVolumesSql, miIDs, minus30d).Rows()
	defer rows.Close()
	if errRaw != nil {
		fmt.Println("loadMarketVolumes", errRaw)
	}

	for rows.Next() {
		temp := models.MiVolume{}
		db.DB.ScanRows(rows, &temp)
		_, ex := byDate[temp.MarketItemID]
		if !ex {
			byDate[temp.MarketItemID] = make(map[string][]models.MiVolume, 32)
		}
		values := byDate[temp.MarketItemID][temp.Dt]
		byDate[temp.MarketItemID][temp.Dt] = append(values, temp)
	}

	compacted := make(map[int64][][]models.MiVolume)

	for miID, dates := range byDate {
		compacted[miID] = make([][]models.MiVolume, 0)
		for _, volumes := range dates {
			compacted[miID] = append(compacted[miID], volumes)
		}
		sort.SliceStable(
			compacted[miID],
			func(i, j int) bool {
				return compacted[miID][i][0].Dt < compacted[miID][j][0].Dt
			},
		)
	}
	//fmt.Println("load:", time.Since(start))

	//align compacted arrays
	for k, vv := range compacted {
		firstSame := true
		for i := 1; i < len(vv); i++ {
			if len(vv[i-1]) > 0 && len(vv[i]) > 0 && vv[i-1][0].M != vv[i][0].M {
				firstSame = false
			}
		}
		if !firstSame {
			for i := 0; i < len(vv); i++ {
				if len(vv[i]) > 0 && vv[i][0].M == 0 {
					compacted[k][i] = append(
						[]models.MiVolume{models.MiVolume{
							MarketItemID: k,
							Dt:           vv[i][0].Dt,
							Vol:          0,
							M:            1,
						}},
						compacted[k][i]...,
					)
				}
			}
		}
	}

	result := make(map[int64]models.AnyArr)
	for k, vv := range compacted {

		maxLen := 0
		maxIndex := -1
		for i, v := range vv {
			if maxLen < len(v) {
				maxLen = len(v)
				maxIndex = i
			}
		}

		temp := make(models.AnyArr, 0)
		for _, values := range vv {

			flattened := make([]models.AnyArr, len(values))
			for i, el := range values {
				flattened[i] = models.AnyArr{el.Vol, el.M}
			}

			if len(values) < maxLen {
				for i := len(values); i < maxLen; i++ {
					flattened = append(flattened, models.AnyArr{0, vv[maxIndex][i].M})
				}
			}
			temp = append(temp, models.AnyArr{values[0].Dt, flattened})
		}

		result[k] = temp
	}
	return result

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
	defer rows.Close()
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
        and d.dt > '%s'
      window w as (partition by d.market_item_id order by dt desc)
  ) x 
  where x.row_number=1`

var sqlMdHist = `
select mi.id, d.dt as d, d.sell_lowest_price as p
  from fp_market_data d,
       fp_market_items mi
  where mi.id = d.market_item_id
    and d.sell_lowest_price>0
    and mi.id in (?)
    and d.dt > '%s'
  order by mi.id, mi.type_id, d.dt, d.sell_lowest_price`

func loadMarketData(miIDs []int64) (map[int64]models.MarketData, map[int64][]models.AnyArr) {
	result := make(map[int64]models.MarketData)
	hist := make(map[int64][]models.AnyArr)

	if len(miIDs) > 0 {
		minus90dFull := time.Now().AddDate(0, 0, -90).Format("2006-01-02 15:04:05")

		rawSqlLast := fmt.Sprintf(sqlMdLast, minus90dFull)
		rows, errRaw := db.DB.Raw(rawSqlLast, miIDs).Rows()
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

		minus30d := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
		rawSqlHist := fmt.Sprintf(sqlMdHist, minus30d)
		rows, errHist := db.DB.Raw(rawSqlHist, miIDs).Rows()
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
			//hist[record.ID] = append(hist[record.ID], record)
			hist[record.ID] = append(hist[record.ID], models.AnyArr{record.D, record.P})
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
				LocationID:    l.LocationID,
				Type:          l.LocationType,
				Name:          LocationName(l.LocationID, l.LocationType),
				CharacterName: l.Character.Name,
				Expiration:    l.Expiration,
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

	sort.SliceStable(
		marketItems,
		func(i, j int) bool {
			return static.Types[marketItems[i].TypeID].Name < static.Types[marketItems[j].TypeID].Name
		},
	)

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
