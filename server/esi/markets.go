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
	TypeID       int32   `json:"type_id"`
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
func (esi *ESI) MarketsOrders(regionID int64, typeID int32, orderType string, page int64) (*MarketsOrdersResponse, error) {
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

//MarketsOrdersAll all orders
func (esi *ESI) MarketsOrdersAll(regionID int64, typeID int32, orderType string) (MarketsOrdersArray, error) {

	result := make(MarketsOrdersArray, 0)
	response, err1 := esi.MarketsOrders(regionID, typeID, orderType, 1)
	if err1 != nil {
		return result, err1
	}

	result = append(result, (response.R)...)
	for i := int64(2); i <= response.Pages; i++ {
		response, err2 := esi.MarketsOrders(regionID, typeID, orderType, i)
		if err2 != nil {
			return result, err2
		}
		result = append(result, (response.R)...)
	}

	// sort.Sort(result)
	return result, nil
}

//MarketsStructures ESI /markets/structures/{structure_id}/
func (esi *ESI) MarketsStructures(structureID int64, page int64) (*MarketsOrdersResponse, error) {
	url := fmt.Sprintf("%s/markets/structures/%d/?page=%d", ESIRootURL, structureID, page)

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

//MarketsStructuresAll all orders
func (esi *ESI) MarketsStructuresAll(structureID int64) (MarketsOrdersArray, error) {

	result := make(MarketsOrdersArray, 0)
	response, err1 := esi.MarketsStructures(structureID, 1)
	if err1 != nil {
		return result, err1
	}

	result = append(result, (response.R)...)
	for i := int64(2); i <= response.Pages; i++ {
		response, err2 := esi.MarketsStructures(structureID, i)
		if err2 != nil {
			return result, err2
		}
		result = append(result, (response.R)...)
	}

	// sort.Sort(result)
	return result, nil
}

//CharacterMarketOrder model
type CharacterMarketOrder struct {
	Duration int `json:"duration"`
	// Escrow        float64   `json:"escrow"`
	// IsBuyOrder    bool      `json:"is_buy_order"`
	IsCorporation bool      `json:"is_corporation"`
	Issued        time.Time `json:"issued"`
	LocationID    int64     `json:"location_id"`
	// MinVolume     int64     `json:"min_volume"`
	OrderID  int64   `json:"order_id"`
	Price    float64 `json:"price"`
	Range    string  `json:"range"`
	RegionID int64   `json:"region_id"`
	// State         string    `json:"state"`
	TypeID       int64 `json:"type_id"`
	VolumeRemain int64 `json:"volume_remain"`
	VolumeTotal  int64 `json:"volume_total"`
}

//CharacterMarketOrders array
type CharacterMarketOrders []CharacterMarketOrder

//CharactersOrdersResponse Characters Orders Response
type CharactersOrdersResponse struct {
	R       CharacterMarketOrders
	Expires time.Time
}

//CharactersOrders character open orders
func (esi *ESI) CharactersOrders(characterID int64) (*CharactersOrdersResponse, error) {
	url := fmt.Sprintf("%s/characters/%d/orders/", ESIRootURL, characterID)
	result := make(CharacterMarketOrders, 0)

	expires, _, err := auth("GET", esi.AccessToken, url, &result)
	if err != nil {
		return nil, err
	}
	return &CharactersOrdersResponse{R: result, Expires: expires}, nil
}
