package middleware

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
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
	ex, errDb := db.DB.Get(&user)
	if !ex || errDb != nil {
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
	charEx, _ := db.DB.Get(&char)

	var user models.User
	setUser, _ := c.Get(USER)

	if !charEx && setUser == nil {
		user = models.User{Role: "U"}
		db.DB.Insert(&user)
	} else if charEx && setUser == nil {
		user.Id = char.UserId
		db.DB.Get(&user)
	} else if charEx && setUser != nil {
		user = setUser.(models.User)
		if user.Id != char.UserId {
			char.UserId = user.Id
			db.DB.Update(&char)
		}

	} else {
		user = setUser.(models.User)
	}

	if !charEx {
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
		db.DB.Insert(&newChar)
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
