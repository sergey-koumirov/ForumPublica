package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
)

func UIOpenMarket(userId int64, params map[string]int64) {

	var char models.Character
	errSel := db.DB.Where("user_id =? and id = ?", userId, params["CharacterId"]).First(&char).Error
	if errSel != nil {
		return
	}

	api := char.GetESI()
	errEsi := api.OpenWindowMarketDetails(params["TypeId"])
	if errEsi != nil {
		fmt.Println("UIOpenMarket errEsi", errEsi)
		return
	}

}
