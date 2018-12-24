package models

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/utils"
	"fmt"
	"strings"
)

//Job job model
type Job struct {
	ID             int64 `gorm:"column:id;primary_key"`
	EsiCharacterID int64 `gorm:"column:esi_character_id"`
	ActivityID     int32 `gorm:"column:activity_id"`

	BlueprintID         int64   `gorm:"column:blueprint_id"`
	BlueprintLocationID int64   `gorm:"column:blueprint_location_id"`
	BlueprintTypeID     int32   `gorm:"column:blueprint_type_id"`
	Cost                int64   `gorm:"column:cost"`
	Duration            int64   `gorm:"column:duration"`
	EndDate             string  `gorm:"column:end_date"`
	FacilityID          int64   `gorm:"column:facility_id"`
	InstallerID         int64   `gorm:"column:installer_id"`
	LicensedRuns        int32   `gorm:"column:licensed_runs"`
	OutputLocationID    int64   `gorm:"column:output_location_id"`
	Probability         float64 `gorm:"column:probability"`
	ProductTypeID       int32   `gorm:"column:product_type_id"`
	Runs                int32   `gorm:"column:runs"`
	StartDate           string  `gorm:"column:start_date"`
	StationID           int64   `gorm:"column:station_id"`
	Status              string  `gorm:"column:status"`

	Character Character `gorm:"foreignkey:EsiCharacterId"`
}

//TableName job model table name
func (j *Job) TableName() string {
	return "esi_jobs"
}

//EndDateH format end date
func (j *Job) EndDateH() string {
	parts := strings.Split(j.EndDate, "T")
	second := strings.Replace(parts[1], "Z", "", 1)
	return fmt.Sprintf("%s %s", parts[0], second)
}

//Rest time to job end
func (j *Job) Rest() string {
	return utils.FormatToHM(j.EndDate)
}

//BlueprintName job blueprint name
func (j *Job) BlueprintName() string {
	return static.Types[j.BlueprintTypeID].Name
}

//ProductName job product name
func (j *Job) ProductName() string {
	return static.Types[j.ProductTypeID].Name
}
