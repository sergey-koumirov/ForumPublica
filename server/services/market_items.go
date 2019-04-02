package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//MarketItemList list
func MarketItemsList(userID int64, page int64) models.MiList {
	cns := make([]models.MarketItem, 0)
	var total int64

	scope := db.DB.Where("user_id = ?", userID)
	scope.Model(&models.MarketItem{}).Count(&total)
	scope.Order("id desc").Limit(PerPage).Offset((page - 1) * PerPage).Find(&cns)

	result := models.MiList{Page: page, Total: total}
	result.Records = make([]models.MiRecord, 0)
	for _, r := range cns {
		temp := models.MiRecord{
			Model: r,
		}
		result.Records = append(result.Records, temp)
	}

	return result
}

//MarketItemsCreate create
func MarketItemsCreate(userID int64) models.MarketItem {
	new := models.MarketItem{
		UserID: userID,
	}

	db.DB.Create(&new)
	return new
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
