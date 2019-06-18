package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//MarketItemsList list
func MarketItemsList(userID int64, page int64) models.MiList {
	cns := make([]models.MarketItem, 0)
	var total int64

	scope := db.DB.Where("user_id = ?", userID)
	scope.Model(&models.MarketItem{}).Count(&total)
	scope.Preload("Locations.Character").Order("id desc").Limit(MIPerPage).Offset((page - 1) * MIPerPage).Find(&cns)

	result := models.MiList{
		Page:       page,
		Total:      total,
		PerPage:    MIPerPage,
		Characters: CharsByUserID(userID),
	}
	result.Records = make([]models.MiRecord, 0)
	for _, r := range cns {

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

		temp := models.MiRecord{
			Model:     r,
			TypeName:  static.Types[r.TypeID].Name,
			Locations: locations,
		}
		result.Records = append(result.Records, temp)
	}

	return result
}

//MarketItemsCreate create
func MarketItemsCreate(userID int64, params map[string]int32) {
	typeID, ex := params["TypeID"]
	if ex {
		new := models.MarketItem{
			UserID: userID,
			TypeID: typeID,
		}
		db.DB.Create(&new)
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
