package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

type CnBlueprint struct {
	Model         models.ConstructionBpo
	IsT2          bool
	CopyTime      int64
	InventCnt     int64
	WholeCopyTime int64
}

type CnBlueprints []CnBlueprint

func ConstructionBluprintAdd(userId int64, cnId int64, params map[string]int32) {
	construction := models.Construction{Id: cnId}
	errDb := db.DB.Where("user_id=?", userId).Find(&construction).Error

	if errDb != nil {
		return
	}

	new := models.ConstructionBpo{
		ConstructionId: construction.Id,
		Kind:           "goal",
		TypeId:         params["BlueprintId"],
		ME:             10,
		TE:             20,
		Qty:            1,
	}

	db.DB.Create(&new)

}

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

func ConstructionBlueprintUpdate(userId int64, cnId int64, bpId int64, params map[string]int32) {
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

	db.DB.Model(&blueprint).Updates(params)

}
