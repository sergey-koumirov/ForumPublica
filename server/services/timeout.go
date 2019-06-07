package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
)

//SaveTimeout save
func SaveTimeout(key string) {
	nowStr := utils.NowUTCStr()
	timeout := models.Timeout{Key: key, Dt: nowStr}
	errDb := db.DB.Save(&timeout).Error
	if errDb != nil {
		fmt.Println(errDb)
	}
}

//GetTimeout get
func GetTimeout(key string, minutes int64) string {
	timeout := models.Timeout{Key: key}
	errDb := db.DB.Where("skey = ?", key).First(&timeout).Error
	if errDb != nil {
		fmt.Println(errDb)
		return ""
	}

	return utils.FormatToHMPlus(timeout.Dt, minutes)
}
