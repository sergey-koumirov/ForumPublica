package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
)

func RefreshJobs(userId int64) {
	var chars []models.Character
	db.DB.Where("user_id = ?", userId).Find(&chars)

	for _, char := range chars {
		api := char.GetESI()
		jobs, err := api.CharactersIndustryJobs(char.Id)

		fmt.Println("------")
		fmt.Printf("%+v\n", jobs, err)
		fmt.Println("------")

	}
}
