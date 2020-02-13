package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//TransactionsList list
func TransactionsList(userID int64, page int64) models.TrList {

	transactions, total := loadTransactions(userID, page)

	result := models.TrList{
		Page:    page,
		Total:   total,
		PerPage: TrPerPage,
	}

	result.Records = make([]models.TrRecord, 0)

	for _, r := range transactions {

		if r.Location == nil {
			var test int64
			db.DB.Model(models.Location{}).Where("id = ?", r.LocationID).Count(&test)
			if test == 0 {
				api, _ := r.Character.GetESI()
				temp := AddLocation(api, r.LocationID, "", 0, 0)
				r.Location = &temp
			}
		}

		temp := models.TrRecord{
			ModelID:       r.ID,
			TypeID:        r.TypeID,
			TypeName:      static.Types[r.TypeID].Name,
			Dt:            r.Dt,
			CharacterName: r.Character.Name,
			Quantity:      r.Quantity,
			Price:         r.UnitPrice,
			IsBuy:         r.IsBuy,
			ClientName:    r.ClientName.Name,
			LocationName:  r.Location.Name,
			ImageURL:      ImageUrl(r.TypeID),
			InSummary:     false,
		}
		result.Records = append(result.Records, temp)
	}

	return result
}

func loadTransactions(userID int64, page int64) ([]models.Transaction, int64) {
	transactions := make([]models.Transaction, 0)
	var total int64

	charIDs := CharIDsByUserID(userID)

	scope := db.DB.Where("esi_character_id in (?)", charIDs)

	scope.Model(&models.Transaction{}).Count(&total)

	scope.Preload("Location").
		Preload("ClientName").
		Preload("Character").
		Order("id desc").
		Limit(TrPerPage).
		Offset((page - 1) * TrPerPage).
		Find(&transactions)
	return transactions, total
}
