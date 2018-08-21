package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
)

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

func ConstructionGet(userId int64, cnId int64) (models.Construction, error) {
	cn := models.Construction{}
	errSel := db.DB.Where("id = ? and user_id = ?", cnId, userId).First(&cn).Error

	if errSel != nil {
		return cn, errSel
	}

	return cn, nil
}
