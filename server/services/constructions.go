package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/services/reverse"
	"ForumPublica/server/services/reverse/jobruns"

	"github.com/jinzhu/gorm"
)

var PER_PAGE int64 = 20

func ConstructionsList(userId int64, page int64) models.CnList {
	cns := make([]models.Construction, 0)
	var total int64

	scope := db.DB.Where("user_id = ?", userId)
	scope.Model(&models.Construction{}).Count(&total)
	scope.Order("id desc").Limit(PER_PAGE).Offset((page - 1) * PER_PAGE).Find(&cns)

	result := models.CnList{Page: page, Total: total}
	result.Records = make([]models.CnRecord, 0)
	for _, r := range cns {
		temp := models.CnRecord{
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

func ConstructionGet(userId int64, cnId int64) (models.CnRecord, error) {
	cn := models.Construction{}
	errSel := db.DB.Preload("Bpos.Expenses").Preload("Runs").Preload("Bpos", bposOrder).Where("id = ? and user_id = ?", cnId, userId).First(&cn).Error

	var result models.CnRecord

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

func loadCn(result *models.CnRecord, cn models.Construction) {
	result.Model = cn

	result.Blueprints = make(models.CnBlueprints, 0)
	for _, r := range cn.Bpos {
		defaultME, _ := static.DefaultMeTe(r.TypeId)

		result.Blueprints = append(
			result.Blueprints,
			models.CnBlueprint{
				Model:         r,
				IsT2:          static.IsT2BPO(r.TypeId),
				DefaultME:     defaultME,
				CopyTime:      0, //todo
				InventCnt:     0, //todo
				WholeCopyTime: 0, //todo
				Expenses:      r.Expenses,
			},
		)
	}

	result.Components = reverse.Assembly(&cn)
	result.Materials = jobruns.RunsToMaterials(result.Components)
}
