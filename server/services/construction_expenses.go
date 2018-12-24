package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//ExpenseParams model
type ExpenseParams struct {
	Description string
	ExValue     float64
	BpoID       int64
}

//ConstructionExpenseAdd add expense to construction
func ConstructionExpenseAdd(userID int64, cnID int64, params ExpenseParams) {

	construction := models.Construction{ID: cnID}
	errDb := db.DB.Where("user_id=?", userID).Find(&construction).Error
	if errDb != nil {
		return
	}

	bpo := models.ConstructionBpo{ID: params.BpoID}
	errDb = db.DB.Where("construction_id=?", cnID).Find(&bpo).Error
	if errDb != nil {
		return
	}

	new := models.ConstructionExpense{
		ConstructionBpoID: params.BpoID,
		Description:       params.Description,
		ExValue:           params.ExValue,
	}

	db.DB.Create(&new)

}

//ConstructionExpenseDelete delete
func ConstructionExpenseDelete(userID int64, cnID int64, id int64) {
	construction := models.Construction{ID: cnID}
	errDb1 := db.DB.Where("user_id=?", userID).Find(&construction).Error
	if errDb1 != nil {
		return
	}

	expense := models.ConstructionExpense{ID: id}
	errDb2 := db.DB.Where("construction_bpo_id in (select b.id from fp_construction_bpos b where b.construction_id =?)", cnID).Find(&expense).Error
	if errDb2 != nil {
		return
	}
	expense.Delete()
}
