package services

import (
	"ForumPublica/sde/static"
	"fmt"
)

func ImageUrl(typeID int32) string {
	code := "icon"
	_, ex := static.Blueprints[typeID]
	if ex {
		code = "bp"
	}
	return fmt.Sprintf("https://images.evetech.net/types/%d/%s", typeID, code)
}
