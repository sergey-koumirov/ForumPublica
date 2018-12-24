package esi

import (
	"fmt"
	"time"
)

//IndustryJob model
type IndustryJob struct {
	ActivityID          int32   `json:"activity_id"`
	BlueprintID         int64   `json:"blueprint_id"`
	BlueprintLocationID int64   `json:"blueprint_location_id"`
	BlueprintTypeID     int32   `json:"blueprint_type_id"`
	Cost                float64 `json:"cost"`
	Duration            int64   `json:"duration"`
	EndDate             string  `json:"end_date"`
	FacilityID          int64   `json:"facility_id"`
	InstallerID         int64   `json:"installer_id"`
	JobID               int64   `json:"job_id"`
	LicensedRuns        int32   `json:"licensed_runs"`
	OutputLocationID    int64   `json:"output_location_id"`
	ProductTypeID       int32   `json:"product_type_id"`
	Runs                int32   `json:"runs"`
	StartDate           string  `json:"start_date"`
	StationID           int64   `json:"station_id"`
	Status              string  `json:"status"`
}

//IndustryJobs array
type IndustryJobs []IndustryJob

//IndustryJobsResponse response
type IndustryJobsResponse struct {
	R       IndustryJobs
	Expires time.Time
}

//CharactersIndustryJobs character jobs
func (esi *ESI) CharactersIndustryJobs(characterID int64) (*IndustryJobsResponse, error) {
	url := fmt.Sprintf("%s/characters/%d/industry/jobs/", ESIRootURL, characterID)
	result := IndustryJobsResponse{}
	result.R = make(IndustryJobs, 0)

	t, _, err := auth("GET", esi.AccessToken, url, &result.R)
	if err != nil {
		return nil, err
	}
	result.Expires = t

	return &result, nil
}
