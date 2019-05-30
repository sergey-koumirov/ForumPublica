package esi

import (
	"fmt"
	"time"
)

//MarketsOrder market order
type MarketsOrder struct {
	Duration     int64   `json:"duration"`
	IsBuyOrder   bool    `json:"is_buy_order"`
	Issued       string  `json:"issued"`
	LocationID   int64   `json:"location_id"`
	MinVolume    int64   `json:"min_volume"`
	OrderID      int64   `json:"order_id"`
	Price        float64 `json:"price"`
	Range        string  `json:"range"`
	SystemID     int64   `json:"system_id"`
	TypeID       int64   `json:"type_id"`
	VolumeRemain int64   `json:"volume_remain"`
	VolumeTotal  int64   `json:"volume_total"`
}

//MarketsOrdersArray market orders
type MarketsOrdersArray []MarketsOrder

func (slice MarketsOrdersArray) Len() int {
	return len(slice)
}
func (slice MarketsOrdersArray) Less(i, j int) bool {
	return slice[i].Price < slice[j].Price
}
func (slice MarketsOrdersArray) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//MarketsOrdersResponse response
type MarketsOrdersResponse struct {
	R       MarketsOrdersArray
	Pages   int64
	Expires time.Time
}

//MarketsOrders ESI /markets/{region_id}/orders/
func (esi *ESI) MarketsOrders(regionID int64, typeID int64, orderType string, page int64) (*MarketsOrdersResponse, error) {
	url := fmt.Sprintf("%s/markets/%d/orders/?page=%d", ESIRootURL, regionID, page)

	if typeID > 0 {
		url = url + fmt.Sprintf("&type_id=%d", typeID)
	}
	if orderType != "" {
		url = url + fmt.Sprintf("&order_type=%s", orderType)
	}
	records := make(MarketsOrdersArray, 0)
	expires, pages, err := get(url, &records)
	if err != nil {
		return nil, err
	}

	result := MarketsOrdersResponse{}
	result.R = records
	result.Pages = pages
	result.Expires = expires

	return &result, nil
}
