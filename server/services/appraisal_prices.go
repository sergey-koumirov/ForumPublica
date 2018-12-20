package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/esi"
	"fmt"
	"strings"
)

//RawAppraisalPrice json description
type RawAppraisalPrice struct {
	Avg        float64 `json:"avg"`
	Median     float64 `json:"median"`
	Percentile float64 `json:"percentile"`
	Min        float64 `json:"min"`
	Max        float64 `json:"max"`
}

//RawAppraisalPrices json description
type RawAppraisalPrices struct {
	All  RawAppraisalPrice `json:"all"`
	Buy  RawAppraisalPrice `json:"buy"`
	Sell RawAppraisalPrice `json:"sell"`
}

//RawAppraisalItem json description
type RawAppraisalItem struct {
	TypeID int32              `json:"typeID"`
	Prices RawAppraisalPrices `json:"prices"`
}

//RawAppraisalData json description
type RawAppraisalData struct {
	Created int64              `json:"created"`
	Items   []RawAppraisalItem `json:"items"`
}

//RawAppraisalResponse json description
type RawAppraisalResponse struct {
	Data RawAppraisalData `json:"appraisal"`
}

// AppraisalUpdatePrices update prices using evepraisal.com
func AppraisalUpdatePrices(typeIds []int32) {

	expiredIds := ExpiredPrices("appraisal", typeIds)

	if len(expiredIds) == 0 {
		return
	}

	typeNames := make([]string, 0)
	for _, typeID := range expiredIds {
		typeNames = append(typeNames, static.Types[typeID].Name)
	}
	names := strings.Join(typeNames, "%0D%0A")

	var result RawAppraisalResponse
	//fmt.Println("Appraisal")
	err := esi.PostSimpleRequest(
		"https://evepraisal.com/appraisal.json?market=jita&persist=no&raw_textarea="+names,
		"",
		&result,
	)
	// fmt.Println(err)
	// fmt.Printf("%+v\n", result)
	if err == nil {
		for _, item := range result.Data.Items {
			fmt.Printf("%+v\n", item)
			UpsertPrice(
				item.TypeID,
				"appraisal",
				item.Prices.Buy.Max,
				item.Prices.Sell.Min,
				0,
			)
		}
	}

	LoadPricesFromDB("appraisal")
}
