package models

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"fmt"
)

//Character character model
type Character struct {
	ID   int64  `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`

	UserID int64 `gorm:"column:user_id"`

	VerExpiresOn          string `gorm:"column:ver_expires_on"`
	VerScopes             string `gorm:"column:ver_scopes"`
	VerTokenType          string `gorm:"column:ver_token_type"`
	VerCharacterOwnerHash string `gorm:"column:ver_character_owner_hash"`

	TokAccessToken  string `gorm:"column:tok_access_token"`
	TokTokenType    string `gorm:"column:tok_token_type"`
	TokExpiresIn    int64  `gorm:"column:tok_expires_in"`
	TokRefreshToken string `gorm:"column:tok_refresh_token"`

	Jobs   []Job   `gorm:"foreignkey:EsiCharacterID"`
	Skills []Skill `gorm:"foreignkey:EsiCharacterID"`
}

//TableName character table name
func (char *Character) TableName() string {
	return "esi_characters"
}

//GetESI get ESI object for char
func (char *Character) GetESI() (esi.ESI, error) {
	errGet := db.DB.First(&char, char.ID).Error
	if errGet != nil {
		fmt.Println("errGet:", errGet, char, char.ID)
		return esi.ESI{}, errGet
	}

	result := esi.ESI{
		AccessToken:  char.TokAccessToken,
		ExpiresOn:    char.VerExpiresOn,
		RefreshToken: char.TokRefreshToken,
	}

	if result.IsAccessTokenExpired() {
		result.RefreshAccessToken()
		char.TokAccessToken = result.AccessToken
		char.VerExpiresOn = result.ExpiresOn
		errUpd := db.DB.Model(&char).Updates(&char).Error
		if errUpd != nil {
			fmt.Println("errUpd:", errUpd)
		}
	}

	return result, nil
}
