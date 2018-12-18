package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

func ConstructionRunAdd(userId int64, cnId int64, params map[string]int64) {
	construction := models.Construction{Id: cnId}
	errDb := db.DB.Where("user_id=?", userId).Find(&construction).Error

	if errDb != nil {
		return
	}

	_, defaultTE := static.DefaultMeTe(int32(params["BlueprintId"]))

	new := models.ConstructionBpoRun{
		ConstructionId: construction.Id,
		TypeId:         int32(params["BlueprintId"]),
		ME:             int32(params["me"]),
		TE:             defaultTE,
		Qty:            params["qty"],
		Repeats:        int32(params["repeats"]),
	}

	db.DB.Create(&new)

}

func ConstructionRunDelete(userId int64, cnId int64, id int64) {
	construction := models.Construction{Id: cnId}
	errDb1 := db.DB.Where("user_id=?", userId).Find(&construction).Error
	if errDb1 != nil {
		return
	}

	run := models.ConstructionBpoRun{Id: id}
	errDb2 := db.DB.Where("construction_id =?", cnId).Find(&run).Error
	if errDb2 != nil {
		return
	}
	run.Delete()
}
