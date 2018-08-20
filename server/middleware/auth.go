package middleware

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	USER_ID = "userId"
	STATE   = "state"
	USER    = "user"
)

func SetUser(c *gin.Context) {

	session := sessions.Default(c)
	value := session.Get(USER_ID)
	if value == nil {
		return
	}

	var (
		userId int64 = value.(int64)
		user         = models.User{Id: userId}
	)
	errDb := db.DB.First(&user, user.Id).Error
	if errDb != nil {
		return
	}

	c.Set(USER, user)

}

func Auth(c *gin.Context) {

	value, _ := c.Get(USER)

	if value == nil {
		url, state := esi.CallbackURL()

		session := sessions.Default(c)
		session.Set(STATE, state)
		session.Save()

		c.Redirect(http.StatusTemporaryRedirect, url)
		c.Abort()
		return
	}

}

func Add(c *gin.Context) {
	url, state := esi.CallbackURL()

	session := sessions.Default(c)
	session.Set(STATE, state)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, url)
	c.Abort()
}

func AuthCallback(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get(STATE)
	if value == nil {
		return
	}

	saved := value.(string)
	state := c.Query(STATE)
	if saved != state {
		return
	}

	code := c.Query("code")
	token, _ := esi.OAuthToken(url.Values{"grant_type": {"authorization_code"}, "code": {code}})
	info, _ := esi.OAuthVerify(token.AccessToken)

	char := models.Character{Id: info.CharacterID}
	errChar := db.DB.First(&char, char.Id).Error
	charEx := errChar == nil

	// fmt.Println(char, errChar, charEx)

	var user models.User
	setUser, _ := c.Get(USER)

	if !charEx && setUser == nil {
		user = models.User{Role: "U"}
		db.DB.Create(&user)
	}

	if charEx && setUser == nil {
		user.Id = char.UserId
		db.DB.First(&user, user.Id)
	} else if charEx && setUser != nil {
		user = setUser.(models.User)

		if user.Id != char.UserId {
			char.UserId = user.Id
			db.DB.Model(&char).Update("user_id", user.Id)
		}
	} else {
		user = setUser.(models.User)
	}

	if charEx {
		char.VerExpiresOn = info.ExpiresOn
		char.VerScopes = info.Scopes
		char.VerTokenType = info.TokenType
		char.VerCharacterOwnerHash = info.CharacterOwnerHash
		char.TokAccessToken = token.AccessToken
		char.TokTokenType = token.TokenType
		char.TokExpiresIn = token.ExpiresIn
		char.TokRefreshToken = token.RefreshToken
		errUpd := db.DB.Model(&char).Updates(char).Error
		if errUpd != nil {
			fmt.Println("errUpd", errUpd)
		}
	} else {
		newChar := models.Character{
			Id:                    info.CharacterID,
			Name:                  info.CharacterName,
			UserId:                user.Id,
			VerExpiresOn:          info.ExpiresOn,
			VerScopes:             info.Scopes,
			VerTokenType:          info.TokenType,
			VerCharacterOwnerHash: info.CharacterOwnerHash,
			TokAccessToken:        token.AccessToken,
			TokTokenType:          token.TokenType,
			TokExpiresIn:          token.ExpiresIn,
			TokRefreshToken:       token.RefreshToken,
		}
		db.DB.Create(&newChar)
	}

	session.Set(USER_ID, user.Id)
	session.Save()

	if setUser == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/app/index")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/app/chars")
	}

}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(STATE, nil)
	session.Set(USER_ID, nil)
	session.Save()
	c.Set(USER, nil)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
