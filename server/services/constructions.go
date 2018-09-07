package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"

	"github.com/jinzhu/gorm"
)

type CnBlueprint struct {
	Model models.ConstructionBpo
}

type CnBlueprints []CnBlueprint

type CnRecord struct {
	Model      models.Construction
	Title      string
	Blueprints CnBlueprints
}

type CnList struct {
	Records []CnRecord
	Page    int64
	Total   int64
}

var PER_PAGE int64 = 20

func ConstructionsList(userId int64, page int64) CnList {
	cns := make([]models.Construction, 0)
	var total int64

	scope := db.DB.Where("user_id = ?", userId)
	scope.Model(&models.Construction{}).Count(&total)
	scope.Order("id desc").Limit(PER_PAGE).Offset((page - 1) * PER_PAGE).Find(&cns)

	result := CnList{Page: page, Total: total}
	result.Records = make([]CnRecord, 0)
	for _, r := range cns {
		temp := CnRecord{
			Model: r,
			Title: "N/A",
		}
		result.Records = append(result.Records, temp)
	}

	return result
}

func ConstructionCreate(userId int64) models.Construction {
	new := models.Construction{
		UserId:      userId,
		Name:        "",
		CitadelType: "",
		RigFactor:   "",
		SpaceType:   "",
	}

	db.DB.Create(&new)
	return new
}

func ConstructionDelete(userId int64, cnId int64) {
	cn := models.Construction{}
	errSel := db.DB.Where("id = ? and user_id = ?", cnId, userId).First(&cn).Error
	if errSel != nil {
		return
	}
	cn.Delete()
}

func bposOrder(db *gorm.DB) *gorm.DB {
	return db.Order("fp_construction_bpos.id asc")
}

func ConstructionGet(userId int64, cnId int64) (CnRecord, error) {
	cn := models.Construction{}
	errSel := db.DB.Preload("Bpos", bposOrder).Where("id = ? and user_id = ?", cnId, userId).First(&cn).Error

	var result CnRecord

	if errSel != nil {
		return result, errSel
	}

	loadCn(&result, cn)

	return result, nil
}

func ConstructionSaveBonus(userId int64, cnId int64, params map[string]string) {
	construction := models.Construction{Id: cnId}
	errDb := db.DB.Where("user_id=?", userId).Find(&construction).Error

	if errDb != nil {
		return
	}

	construction.CitadelType = params["CitadelType"]
	construction.RigFactor = params["RigFactor"]
	construction.SpaceType = params["SpaceType"]
	db.DB.Save(construction)
}

func ConstructionAddBluprint(userId int64, cnId int64, params map[string]int32) {
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

func loadCn(result *CnRecord, cn models.Construction) {
	result.Model = cn
	result.Blueprints = make(CnBlueprints, 0)
	for _, r := range cn.Bpos {
		result.Blueprints = append(
			result.Blueprints,
			CnBlueprint{
				Model: r,
			},
		)
	}
}
