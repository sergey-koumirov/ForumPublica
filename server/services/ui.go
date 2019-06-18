package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
)

//UIOpenMarket open window request
func UIOpenMarket(userID int64, params map[string]int64) {

	var char models.Character
	errSel := db.DB.Where("user_id =? and id = ?", userID, params["CharacterId"]).First(&char).Error
	if errSel != nil {
		return
	}

	api, _ := char.GetESI()
	errEsi := api.OpenWindowMarketDetails(params["TypeId"])
	if errEsi != nil {
		fmt.Println("UIOpenMarket errEsi", errEsi)
		return
	}

}
