package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
)

//LocationParams location params
type LocationParams struct {
	LocationID        int64
	LocationType      string
	StoreLocationID   int64
	StoreLocationType string
	CharacterID       int64
}

//MarketItemsLocationsCreate create
func MarketItemsLocationsCreate(userID int64, marketItemID int64, params LocationParams) {

	fmt.Println("MarketItemsLocationsCreate", marketItemID)

	mi := models.MarketItem{}
	errDb1 := db.DB.Model(&models.MarketItem{}).Where("id = ? and user_id = ?", marketItemID, userID).Find(&mi).Error
	if errDb1 != nil {
		fmt.Println("MarketItemsLocationsCreate: errDb1 ", errDb1)
		return
	}

	new := models.MarketLocation{
		MarketItemID:      marketItemID,
		LocationID:        params.LocationID,
		LocationType:      params.LocationType,
		StoreLocationID:   params.StoreLocationID,
		StoreLocationType: params.StoreLocationType,
		EsiCharacterID:    params.CharacterID,
	}
	db.DB.Create(&new)

}

//MarketItemsLocationsDelete delete
func MarketItemsLocationsDelete(userID int64, miID int64, lID int64) {
	mi := models.MarketItem{}
	errSel := db.DB.Where("id = ? and user_id = ?", miID, userID).First(&mi).Error
	if errSel != nil {
		return
	}

	l := models.MarketLocation{}
	errL := db.DB.Where("id = ? and market_item_id = ?", lID, miID).First(&l).Error
	if errL != nil {
		return
	}
	l.Delete()
}

//LocationName get location name
func LocationName(id int64, t string) string {
	if t == "solar_system" {
		return static.SolarSystems[id].Name
	} else if t == "station" || t == "structure" {
		n := models.Location{ID: id}
		db.DB.Find(&n)
		return n.Name
	}
	return ""
}