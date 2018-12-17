package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

type ExpenseParams struct {
	Description string
	ExValue     float64
	BpoId       int64
}

//ConstructionExpenseAdd add expense to construction
func ConstructionExpenseAdd(userId int64, cnId int64, params ExpenseParams) {

	construction := models.Construction{Id: cnId}
	errDb := db.DB.Where("user_id=?", userId).Find(&construction).Error
	if errDb != nil {
		return
	}

	bpo := models.ConstructionBpo{Id: params.BpoId}
	errDb = db.DB.Where("construction_id=?", cnId).Find(&bpo).Error
	if errDb != nil {
		return
	}

	new := models.ConstructionExpense{
		ConstructionBpoId: params.BpoId,
		Description:       params.Description,
		ExValue:           params.ExValue,
	}

	db.DB.Create(&new)

}

// func ConstructionExpenseDelete(userId int64, cnId int64, bpId int64) {
// 	construction := models.Construction{Id: cnId}
// 	errDb1 := db.DB.Where("user_id=?", userId).Find(&construction).Error
// 	if errDb1 != nil {
// 		return
// 	}

// 	blueprint := models.ConstructionBpo{Id: bpId}
// 	errDb2 := db.DB.Where("construction_id=?", cnId).Find(&blueprint).Error
// 	if errDb2 != nil {
// 		return
// 	}

// 	blueprint.Delete()

// }
