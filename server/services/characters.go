package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

//CharIDName model
type CharIDName struct {
	ID   int64
	Name string
}

//CharsByUserID chars by user id
func CharsByUserID(userID int64) []CharIDName {
	chars := make([]models.Character, 0)
	errDb := db.DB.Where("user_id=?", userID).Order("name").Find(&chars).Error
	if errDb != nil {
		return []CharIDName{}
	}

	result := make([]CharIDName, 0)
	for _, r := range chars {
		result = append(result, CharIDName{ID: r.ID, Name: r.Name})
	}

	return result
}
