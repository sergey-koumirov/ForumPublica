package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
	"time"
)

var (
	//Prices global variable for prices
	Prices map[string]map[int32]models.Price
)

//InitPrices init prices
func InitPrices() {
	Prices = map[string]map[int32]models.Price{
		"appraisal": map[int32]models.Price{},
	}
	LoadPricesFromDB("appraisal")
}

//GetDefaultPrice get default price
func GetDefaultPrice(id int32) float64 {
	return Prices["appraisal"][id].SellPrice
}

//LoadPricesFromDB load prices
func LoadPricesFromDB(source string) {
	prices := []models.Price{}
	db.DB.Where("source=?", source).Find(&prices)
	for _, p := range prices {
		Prices[source][p.TypeID] = p
	}
}

//ExpiredPrices filter out recently updated type_ids
func ExpiredPrices(source string, typeIDs []int32) []int32 {
	result := []int32{}

	prices := []models.Price{}
	err := db.DB.Where("type_id in (?) and source = ?", typeIDs, source).Find(&prices).Error
	if err != nil {
		fmt.Println(err)
		return result
	}

	excludedIds := map[int32]int32{}
	for _, p := range prices {
		if utils.DbStrToMinut(p.Dt) < 30 {
			excludedIds[p.TypeID] = 1
		}
	}

	for _, id := range typeIDs {
		_, ex := excludedIds[id]
		if !ex {
			result = append(result, id)
		}
	}

	return result
}

//UpsertPrice insert or update price
func UpsertPrice(typeID int32, source string, buy float64, sell float64, volume int64) {
	temp := models.Price{
		TypeID:    typeID,
		Source:    source,
		BuyPrice:  buy,
		SellPrice: sell,
		Dt:        time.Now().Format("2006-01-02 15:04:05"),
		MarketVol: volume,
	}

	var price models.Price
	err := db.DB.Where("type_id = ? and source = ?", typeID, source).First(&price).Error
	if err != nil {
		err1 := db.DB.Create(&temp).Error
		if err1 != nil {
			fmt.Println(err1)
		}
	} else {
		temp.ID = price.ID
		err2 := db.DB.Save(&temp).Error
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}
