package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

type CharIdName struct {
	Id   int64
	Name string
}

func CharsByUserId(userId int64) []CharIdName {
	chars := make([]models.Character, 0)
	errDb := db.DB.Where("user_id=?", userId).Order("name").Find(&chars).Error
	if errDb != nil {
		return []CharIdName{}
	}

	result := make([]CharIdName, 0)
	for _, r := range chars {
		result = append(result, CharIdName{Id: r.Id, Name: r.Name})
	}

	return result
}
