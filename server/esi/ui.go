package esi

import (
	"errors"
	"fmt"
)

//OpenWindowMarketDetails open window
func (esi *ESI) OpenWindowMarketDetails(typeID int64) error {

	url := fmt.Sprintf("%s/ui/openwindow/marketdetails/?type_id=%d", ESIRootURL, typeID)

	result := Error{}

	_, _, err := auth("POST", esi.AccessToken, url, &result)
	if err != nil {
		return err
	}
	if result.Error != "" {
		return errors.New(result.Error)
	}
	return nil
}
