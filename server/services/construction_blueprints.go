package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//ConstructionBluprintAdd add blueprint to construction
func ConstructionBluprintAdd(userID int64, cnID int64, params map[string]int32) {
	construction := models.Construction{ID: cnID}
	errDb := db.DB.Where("user_id=?", userID).Find(&construction).Error

	if errDb != nil {
		return
	}

	defaultME, defaultTE := static.DefaultMeTe(params["BlueprintId"])

	new := models.ConstructionBpo{
		ConstructionID: construction.ID,
		Kind:           "goal",
		TypeID:         params["BlueprintId"],
		ME:             defaultME,
		TE:             defaultTE,
		Qty:            1,
	}
	db.DB.Create(&new)
}

//ConstructionBlueprintDelete delete
func ConstructionBlueprintDelete(userID int64, cnID int64, bpID int64) {
	construction := models.Construction{ID: cnID}
	errDb1 := db.DB.Where("user_id=?", userID).Find(&construction).Error
	if errDb1 != nil {
		return
	}

	blueprint := models.ConstructionBpo{ID: bpID}
	errDb2 := db.DB.Where("construction_id=?", cnID).Find(&blueprint).Error
	if errDb2 != nil {
		return
	}

	blueprint.Delete()

}

//ConstructionBlueprintUpdate update
func ConstructionBlueprintUpdate(userID int64, cnID int64, bpID int64, params map[string]int32) {
	construction := models.Construction{ID: cnID}
	errDb1 := db.DB.Where("user_id=?", userID).Find(&construction).Error
	if errDb1 != nil {
		return
	}

	blueprint := models.ConstructionBpo{ID: bpID}
	errDb2 := db.DB.Where("construction_id=?", cnID).Find(&blueprint).Error
	if errDb2 != nil {
		return
	}

	db.DB.Model(&blueprint).Updates(params)

}
