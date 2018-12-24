package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//ConstructionRunAdd add
func ConstructionRunAdd(userID int64, cnID int64, params map[string]int64) {
	construction := models.Construction{ID: cnID}
	errDb := db.DB.Where("user_id=?", userID).Find(&construction).Error

	if errDb != nil {
		return
	}

	_, defaultTE := static.DefaultMeTe(int32(params["BlueprintId"]))

	new := models.ConstructionBpoRun{
		ConstructionID: construction.ID,
		TypeID:         int32(params["BlueprintId"]),
		ME:             int32(params["me"]),
		TE:             defaultTE,
		Qty:            params["qty"],
		Repeats:        int32(params["repeats"]),
	}

	db.DB.Create(&new)
}

//ConstructionRunDelete delete
func ConstructionRunDelete(userID int64, cnID int64, id int64) {
	construction := models.Construction{ID: cnID}
	errDb1 := db.DB.Where("user_id=?", userID).Find(&construction).Error
	if errDb1 != nil {
		return
	}

	run := models.ConstructionBpoRun{ID: id}
	errDb2 := db.DB.Where("construction_id =?", cnID).Find(&run).Error
	if errDb2 != nil {
		return
	}
	run.Delete()
}
