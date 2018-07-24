package models

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
)

type Character struct {
	Id   int64  `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`

	UserId int64 `gorm:"column:user_id"`

	VerExpiresOn          string `gorm:"column:ver_expires_on"`
	VerScopes             string `gorm:"column:ver_scopes"`
	VerTokenType          string `gorm:"column:ver_token_type"`
	VerCharacterOwnerHash string `gorm:"column:ver_character_owner_hash"`

	TokAccessToken  string `gorm:"column:tok_access_token"`
	TokTokenType    string `gorm:"column:tok_token_type"`
	TokExpiresIn    int64  `gorm:"column:tok_expires_in"`
	TokRefreshToken string `gorm:"column:tok_refresh_token"`

	Jobs []Job `gorm:"foreignkey:EsiCharacterId"`
}

func (c *Character) TableName() string {
	return "esi_characters"
}

func (char *Character) GetESI() esi.ESI {
	db.DB.First(char, char.Id)

	result := esi.ESI{
		AccessToken:  char.TokAccessToken,
		ExpiresOn:    char.VerExpiresOn,
		RefreshToken: char.TokRefreshToken,
	}

	if result.IsAccessTokenExpired() {
		result.RefreshAccessToken()
		char.TokAccessToken = result.AccessToken
		char.VerExpiresOn = result.ExpiresOn
		db.DB.Update(char)
	}

	return result
}
