package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

func ConstructionBlueprintDelete(userId int64, cnId int64, bpId int64) {
	construction := models.Construction{Id: cnId}
	errDb1 := db.DB.Where("user_id=?", userId).Find(&construction).Error
	if errDb1 != nil {
		return
	}

	blueprint := models.ConstructionBpo{Id: bpId}
	errDb2 := db.DB.Where("construction_id=?", cnId).Find(&blueprint).Error
	if errDb2 != nil {
		return
	}

	blueprint.Delete()

}
