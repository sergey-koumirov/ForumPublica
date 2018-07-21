package esi

import (
	"fmt"
	"time"
)

type ESIIndustryJob struct {
	ActivityId          int32   `json:"activity_id"`
	BlueprintId         int64   `json:"blueprint_id"`
	BlueprintLocationId int64   `json:"blueprint_location_id"`
	BlueprintTypeId     int32   `json:"blueprint_type_id"`
	Cost                float64 `json:"cost"`
	Duration            int64   `json:"duration"`
	EndDate             string  `json:"end_date"`
	FacilityId          int64   `json:"facility_id"`
	InstallerId         int64   `json:"installer_id"`
	jobId               int64   `json:"job_id"`
	LicensedRuns        int32   `json:"licensed_runs"`
	OutputLocationId    int64   `json:"output_location_id"`
	ProductTypeId       int32   `json:"product_type_id"`
	Runs                int32   `json:"runs"`
	StartDate           string  `json:"start_date"`
	StationId           int64   `json:"station_id"`
	Status              string  `json:"status"`
}

type ESIIndustryJobs []ESIIndustryJob

type IndustryJobsResponse struct {
	R       ESIIndustryJobs
	Expires time.Time
}

func (esi *ESI) CharactersIndustryJobs(characterId int64) (*IndustryJobsResponse, error) {
	url := fmt.Sprintf("%s/characters/%d/industry/jobs/", ESI_ROOT_URL, characterId)
	result := IndustryJobsResponse{}
	result.R = make(ESIIndustryJobs, 0)

	t, _, err := auth("GET", esi.AccessToken, url, &result.R)
	if err != nil {
		return nil, err
	}
	result.Expires = t

	return &result, nil
}
