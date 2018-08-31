package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
	"math"
)

type CnRecord struct {
	Model models.Construction
	Title string
}

type CnList struct {
	Records    []CnRecord
	Page       int64
	TotalPages int64
}

var PER_PAGE int64 = 20

func ConstructionsList(userId int64, page int64) CnList {
	cns := make([]models.Construction, 0)
	var total int64

	scope := db.DB.Where("user_id = ?", userId)
	scope.Model(&models.Construction{}).Count(&total)
	scope.Order("id desc").Limit(PER_PAGE).Offset(page * PER_PAGE).Find(&cns)

	result := CnList{Page: page, TotalPages: int64(math.Ceil(float64(total) / float64(PER_PAGE)))}
	result.Records = make([]CnRecord, 0)
	for _, r := range cns {
		temp := CnRecord{
			Model: r,
			Title: "N/A",
		}
		result.Records = append(result.Records, temp)
	}

	fmt.Printf("%+v\n", result)

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

func ConstructionGet(userId int64, cnId int64) (models.Construction, error) {
	cn := models.Construction{}
	errSel := db.DB.Where("id = ? and user_id = ?", cnId, userId).First(&cn).Error

	if errSel != nil {
		return cn, errSel
	}

	return cn, nil
}

func ConstructionSaveBonus(userId int64, cnId int64, params map[string]string) {
	construction := models.Construction{Id: cnId}
	errDb := db.DB.Where("user_id=?", userId).Find(&construction).Error

	if errDb == nil {
		construction.CitadelType = params["CitadelType"]
		construction.RigFactor = params["RigFactor"]
		construction.SpaceType = params["SpaceType"]
		db.DB.Save(construction)
	}
}
