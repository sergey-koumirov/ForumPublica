package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
)

func RefreshJobs(userId int64) {
	var chars []models.Character
	db.DB.Where("user_id = ?", userId).Find(&chars)

	sql := "delete from esi_jobs where esi_character_id = ?"

	for _, char := range chars {
		db.DB.Exec(sql, char.Id)

		api := char.GetESI()
		jobs, errEsi := api.CharactersIndustryJobs(char.Id)
		if errEsi != nil {
			fmt.Println("errEsi", errEsi)
		}

		for _, job := range jobs.R {
			var temp models.Job = models.Job{}
			temp.Id = job.JobId
			temp.EsiCharacterId = char.Id
			temp.ActivityId = job.ActivityId
			temp.BlueprintId = job.BlueprintId
			temp.BlueprintLocationId = job.BlueprintLocationId
			temp.BlueprintTypeId = job.BlueprintTypeId
			temp.Cost = int64(job.Cost)
			temp.Duration = job.Duration
			temp.EndDate = job.EndDate
			temp.FacilityId = job.FacilityId
			temp.InstallerId = job.InstallerId
			temp.LicensedRuns = job.LicensedRuns
			temp.OutputLocationId = job.OutputLocationId
			temp.ProductTypeId = job.ProductTypeId
			temp.Runs = job.Runs
			temp.StartDate = job.StartDate
			temp.StationId = job.StationId
			temp.Status = job.Status
			errIns := db.DB.Create(&temp).Error
			if errIns != nil {
				fmt.Println("errIns", errIns)
			}
		}

	}
}
