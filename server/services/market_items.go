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
	scope.Order("id desc").Limit(MIPerPage).Offset((page - 1) * MIPerPage).Find(&cns)

	result := models.MiList{
		Page:       page,
		Total:      total,
		PerPage:    MIPerPage,
		Characters: CharsByUserID(userID),
	}
	result.Records = make([]models.MiRecord, 0)
	for _, r := range cns {
		temp := models.MiRecord{
			Model:    r,
			TypeName: static.Types[r.TypeID].Name,
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
