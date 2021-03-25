package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
)

//DeviationsList list
func DeviationsList(user models.User) ([]models.DvRecord, []models.DvRecord) {

	var typeIDs []int32

	if user.ID > 0 {
		db.DB.Model(&models.MarketItem{}).Where("user_id = ?", user.ID).Pluck("type_id", &typeIDs)
	}

	records := make([]models.Deviation, 0)

	db.DB.Where("k>0").Order("k desc").Find(&records)

	resultOver := make([]models.DvRecord, 0)
	resultUnder := make([]models.DvRecord, 0)
	for _, r := range records {
		//t := static.Types[r.ID]
		p := static.ProductByBpoID(r.ID)
		g := static.Groups[p.GroupID]
		temp := models.DvRecord{
			Description:   fmt.Sprintf("%10d | %-38s | %-60s | %6.2f", p.ID, g.Name, p.Name, r.K),
			TypeName:      p.Name,
			K:             r.K,
			HasMarketItem: utils.FindInt32(typeIDs, p.ID) > -1,
		}
		if r.K < 1 {
			resultUnder = append(resultUnder, temp)
		} else {
			resultOver = append(resultOver, temp)
		}

	}
	return resultOver, resultUnder
}
