package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
)

//ConstructionsList list
func DeviationsList() []models.DvRecord {
	records := make([]models.Deviation, 0)
	db.DB.Where("k>0").Order("k desc").Find(&records)
	result := make([]models.DvRecord, 0)
	for _, r := range records {
		t := static.Types[r.ID]
		p := static.ProductByBpoID(r.ID)
		g := static.Groups[p.GroupID]
		temp := models.DvRecord{
			Description: fmt.Sprintf("%10d | %-38s | %-60s | %6.2f", r.ID, g.Name, t.Name, r.K),
			K:           r.K,
		}
		result = append(result, temp)
	}
	return result
}
