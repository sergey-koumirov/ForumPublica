package esi

import (
	"errors"
	"fmt"
)

func (esi *ESI) OpenWindowMarketDetails(typeId int64) error {

	url := fmt.Sprintf("%s/ui/openwindow/marketdetails/?type_id=%d", ESI_ROOT_URL, typeId)

	result := EsiError{}

	_, _, err := auth("POST", esi.AccessToken, url, &result)
	if err != nil {
		return err
	}
	if result.Error != "" {
		return errors.New(result.Error)
	}
	return nil
}
