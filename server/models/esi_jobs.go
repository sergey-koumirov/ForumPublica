package models

type Job struct {
	Id             int64  `xorm:"job_id pk"`
	EsiCharacterId string `xorm:"esi_character_id"`
	ActivityId     int32  `xorm:"activity_id"`

	BlueprintId         int64   `xorm:"blueprint_id"`
	BlueprintLocationId int64   `xorm:"blueprint_location_id"`
	BlueprintTypeId     int32   `xorm:"blueprint_type_id"`
	Cost                int64   `xorm:"cost"`
	Duration            int64   `xorm:"duration"`
	EndDate             string  `xorm:"end_date"`
	FacilityId          int64   `xorm:"facility_id"`
	InstallerId         int64   `xorm:"installer_id"`
	LicensedRuns        int32   `xorm:"licensed_runs"`
	OutputLocationId    int64   `xorm:"output_location_id"`
	Probability         float64 `xorm:"probability"`
	ProductTypeId       int32   `xorm:"product_type_id"`
	Runs                int32   `xorm:"runs"`
	StartDate           string  `xorm:"start_date"`
	StationId           int64   `xorm:"station_id"`
	Status              string  `xorm:"status"`
}

func (j *Job) TableName() string {
	return "esi_jobs"
}
