package models

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
)

type Character struct {
	Id   int64  `xorm:"pk"`
	Name string `xorm:"name"`

	UserId int64 `xorm:"user_id"`

	VerExpiresOn          string `xorm:"ver_expires_on"`
	VerScopes             string `xorm:"ver_scopes"`
	VerTokenType          string `xorm:"ver_token_type"`
	VerCharacterOwnerHash string `xorm:"ver_character_owner_hash"`

	TokAccessToken  string `xorm:"tok_access_token"`
	TokTokenType    string `xorm:"tok_token_type"`
	TokExpiresIn    int64  `xorm:"tok_expires_in"`
	TokRefreshToken string `xorm:"tok_refresh_token"`
}

func (c *Character) TableName() string {
	return "esi_characters"
}

func (char *Character) GetESI() esi.ESI {
	db.DB.Get(char)

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
