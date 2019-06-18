package services

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

//CharJobs model
type CharJobs struct {
	Char       models.Character
	MnfJobs    int32
	MaxMnfJobs int32
	ScnJobs    int32
	MaxScnJobs int32
}

//JOBS jobs timeout name
var JOBS = "RefreshJobs"

//CharJobsList job list for client
func CharJobsList(userID int64) []CharJobs {
	var chars []models.Character

	errSel := db.DB.Preload("Jobs", jobsOrder).
		Preload("Skills").
		Where("user_id=?", userID).
		Order("name").
		Find(&chars).Error
	if errSel != nil {
		fmt.Println("errSel", errSel)
	}

	result := make([]CharJobs, len(chars))

	mnfIds := []int32{1}
	scnIds := []int32{3, 4, 5, 8}

	for i, char := range chars {
		result[i] = CharJobs{
			Char:       char,
			MnfJobs:    jobsCnt(char.Jobs, mnfIds),
			MaxMnfJobs: maxJobs(char.Skills, 3387, 24625),
			ScnJobs:    jobsCnt(char.Jobs, scnIds),
			MaxScnJobs: maxJobs(char.Skills, 3406, 24624),
		}
	}

	return result

}

func jobsCnt(jobs []models.Job, ids []int32) int32 {
	result := int32(0)
	for _, job := range jobs {
		for _, id := range ids {
			if job.ActivityID == id {
				result++
			}
		}
	}
	return result
}

func maxJobs(skills []models.Skill, id1 int32, id2 int32) int32 {
	result := int32(1)
	for _, skill := range skills {
		if skill.SkillID == id1 || skill.SkillID == id2 {
			result = result + skill.Level
		}
	}
	return result
}

func jobsOrder(db *gorm.DB) *gorm.DB {
	return db.Order("esi_jobs.end_date asc")
}

//RefreshJobs refresh jobs
func RefreshJobs(userID int64) {
	var chars []models.Character
	db.DB.Where("user_id = ?", userID).Find(&chars)

	sql := "delete from esi_jobs where esi_character_id = ?"

	for _, char := range chars {
		db.DB.Exec(sql, char.ID)

		api, _ := char.GetESI()
		jobs, errEsi := api.CharactersIndustryJobs(char.ID)
		if errEsi != nil {
			fmt.Println("errEsi", errEsi)
		} else {
			for _, job := range jobs.R {
				var temp models.Job = models.Job{}
				temp.ID = job.JobID
				temp.EsiCharacterID = char.ID
				temp.ActivityID = job.ActivityID
				temp.BlueprintID = job.BlueprintID
				temp.BlueprintLocationID = job.BlueprintLocationID
				temp.BlueprintTypeID = job.BlueprintTypeID
				temp.Cost = int64(job.Cost)
				temp.Duration = job.Duration
				temp.EndDate = job.EndDate
				temp.FacilityID = job.FacilityID
				temp.InstallerID = job.InstallerID
				temp.LicensedRuns = job.LicensedRuns
				temp.OutputLocationID = job.OutputLocationID
				temp.ProductTypeID = job.ProductTypeID
				temp.Runs = job.Runs
				temp.StartDate = job.StartDate
				temp.StationID = job.StationID
				temp.Status = job.Status
				errIns := db.DB.Create(&temp).Error
				if errIns != nil {
					fmt.Println("errIns", errIns)
				}
			}
		}

	}

	SaveTimeout(JOBS)
}
