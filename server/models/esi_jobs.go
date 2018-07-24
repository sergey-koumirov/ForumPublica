package models

import (
	"ForumPublica/server/utils"
	"fmt"
	"strings"
)

type Job struct {
	Id             int64 `gorm:"column:id;primary_key"`
	EsiCharacterId int64 `gorm:"column:esi_character_id"`
	ActivityId     int32 `gorm:"column:activity_id"`

	BlueprintId         int64   `gorm:"column:blueprint_id"`
	BlueprintLocationId int64   `gorm:"column:blueprint_location_id"`
	BlueprintTypeId     int32   `gorm:"column:blueprint_type_id"`
	Cost                int64   `gorm:"column:cost"`
	Duration            int64   `gorm:"column:duration"`
	EndDate             string  `gorm:"column:end_date"`
	FacilityId          int64   `gorm:"column:facility_id"`
	InstallerId         int64   `gorm:"column:installer_id"`
	LicensedRuns        int32   `gorm:"column:licensed_runs"`
	OutputLocationId    int64   `gorm:"column:output_location_id"`
	Probability         float64 `gorm:"column:probability"`
	ProductTypeId       int32   `gorm:"column:product_type_id"`
	Runs                int32   `gorm:"column:runs"`
	StartDate           string  `gorm:"column:start_date"`
	StationId           int64   `gorm:"column:station_id"`
	Status              string  `gorm:"column:status"`

	Character Character `gorm:"foreignkey:EsiCharacterId"`
}

func (j *Job) TableName() string {
	return "esi_jobs"
}

func (j *Job) EndDateH() string {
	parts := strings.Split(j.EndDate, "T")
	second := strings.Replace(parts[1], "Z", "", 1)
	return fmt.Sprintf("%s %s", parts[0], second)
}

func (j *Job) Rest() string {
	return utils.FormatToHM(j.EndDate)
}
