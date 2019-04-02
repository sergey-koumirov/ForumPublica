package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/services/reverse"
	"ForumPublica/server/services/reverse/jobruns"

	"github.com/jinzhu/gorm"
)

//ConstructionsList list
func ConstructionsList(userID int64, page int64) models.CnList {
	cns := make([]models.Construction, 0)
	var total int64

	scope := db.DB.Where("user_id = ?", userID)
	scope.Model(&models.Construction{}).Count(&total)
	scope.Order("id desc").Limit(PerPage).Offset((page - 1) * PerPage).Find(&cns)

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

//ConstructionCreate create
func ConstructionCreate(userID int64) models.Construction {
	new := models.Construction{
		UserID:      userID,
		Name:        "",
		CitadelType: "",
		RigFactor:   "",
		SpaceType:   "",
	}

	db.DB.Create(&new)
	return new
}

//ConstructionDelete delete
func ConstructionDelete(userID int64, cnID int64) {
	cn := models.Construction{}
	errSel := db.DB.Where("id = ? and user_id = ?", cnID, userID).First(&cn).Error
	if errSel != nil {
		return
	}
	cn.Delete()
}

func bposOrder(db *gorm.DB) *gorm.DB {
	return db.Order("fp_construction_bpos.id asc")
}

//ConstructionGet get
func ConstructionGet(userID int64, cnID int64) (models.CnRecord, error) {
	cn := models.Construction{}
	errSel := db.DB.Preload("Bpos.Expenses").Preload("Runs").Preload("Bpos", bposOrder).Where("id = ? and user_id = ?", cnID, userID).First(&cn).Error

	var result models.CnRecord

	if errSel != nil {
		return result, errSel
	}

	loadCn(&result, cn)

	return result, nil
}

//ConstructionSaveBonus save citadel bonuses
func ConstructionSaveBonus(userID int64, cnID int64, params map[string]string) {
	construction := models.Construction{ID: cnID}
	errDb := db.DB.Where("user_id=?", userID).Find(&construction).Error

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
		defaultME, _ := static.DefaultMeTe(r.TypeID)

		result.Blueprints = append(
			result.Blueprints,
			models.CnBlueprint{
				Model:         r,
				IsT2:          static.IsT2BPO(r.TypeID),
				DefaultME:     defaultME,
				CopyTime:      int32( float64(static.T1CopyTime(r.TypeID)) * (1.0 - 5.0 * 5.0 / 100.0) ),
				InventTime:    int32( float64(static.InventTime(r.TypeID)) * (1.0 - 3.0 * 5.0 / 100.0) ),
				InventCnt:     static.InventCount(r.TypeID, r.Qty),
				Expenses:      r.Expenses,
			},
		)
	}

	result.Components = reverse.Assembly(&cn)
	result.Materials = jobruns.RunsToMaterials(result.Components)

	typeIds := make([]int32, len(result.Materials))
	for i, m := range result.Materials{
		typeIds[i] = m.Model.ID
	}
	AppraisalUpdatePrices(typeIds)
	for i, m := range result.Materials{
		result.Materials[i].Price = GetDefaultPrice(m.Model.ID)
	}

}
