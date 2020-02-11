package esi

import (
	"fmt"
	"time"
)

type ESIAsset struct {
	IsBlueprintCopy bool   `json:"is_blueprint_copy"`
	IsSingleton     bool   `json:"is_singleton"`
	ItemId          int64  `json:"item_id"`
	LocationFlag    string `json:"location_flag"`
	LocationId      int64  `json:"location_id"`
	LocationType    string `json:"location_type"`
	TypeId          int32  `json:"type_id"`
	Quantity        int64  `json:"quantity"`
}

type CharactersAssetsResponse struct {
	R       []ESIAsset
	Expires time.Time
	Pages   int64
}

func (esi *ESI) CharactersAssets(characterId int64, page int64) (*CharactersAssetsResponse, error) {
	url := fmt.Sprintf("%s/characters/%d/assets/?page=%d", ESIRootURL, characterId, page)
	records := make([]ESIAsset, 0)

	expires, pages, err := auth("GET", esi.AccessToken, url, &records)
	if err != nil {
		return nil, err
	}

	result := CharactersAssetsResponse{}
	result.R = records
	result.Pages = pages
	result.Expires = expires

	return &result, nil
}

//CharactersAssetsAll all assets
func (esi *ESI) CharactersAssetsAll(characterId int64) ([]ESIAsset, error) {

	result := make([]ESIAsset, 0)
	response, err1 := esi.CharactersAssets(characterId, 1)
	if err1 != nil {
		return result, err1
	}

	result = append(result, (response.R)...)
	for i := int64(2); i <= response.Pages; i++ {
		response, err2 := esi.CharactersAssets(characterId, i)
		if err2 != nil {
			return result, err2
		}
		result = append(result, (response.R)...)
	}

	// sort.Sort(result)
	return result, nil
}

type QtyByType map[int32]int64
type ItemsByLocation map[int64]QtyByType

func (esi *ESI) CharactersAssetsTypeIdsByLocationID(characterId int64) (ItemsByLocation, error) {
	result := make(ItemsByLocation)

	allAssets, err1 := esi.CharactersAssetsAll(characterId)
	if err1 != nil {
		return result, err1
	}

	childParent := make(map[int64]ESIAsset, 0)
	for _, a := range allAssets {
		childParent[a.ItemId] = a
	}

	for _, v := range childParent {
		pid := v.LocationId
		for {
			parent, exists := childParent[pid]
			if !exists {
				if pid == v.LocationId || v.LocationFlag == "Hangar" || v.LocationFlag == "Unlocked" {
					_, exLoc := result[pid]
					if !exLoc {
						result[pid] = make(QtyByType)
					}
					result[pid][v.TypeId] = result[pid][v.TypeId] + v.Quantity
				}
				break
			} else {
				pid = parent.LocationId
			}
		}
	}

	return result, nil

}
