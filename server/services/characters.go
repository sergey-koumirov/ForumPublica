package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//CharsByUserID chars by user id
func CharsByUserID(userID int64) []models.CharIDName {
	chars := make([]models.Character, 0)
	errDb := db.DB.Where("user_id=?", userID).Order("name").Find(&chars).Error
	if errDb != nil {
		return []models.CharIDName{}
	}

	result := make([]models.CharIDName, 0)
	for _, r := range chars {
		result = append(result, models.CharIDName{ID: r.ID, Name: r.Name})
	}

	return result
}
