package esi

import (
	"fmt"
	"net/url"
	"time"
)

//ESI main object
type ESI struct {
	AccessToken  string
	RefreshToken string
	ExpiresOn    string
}

//IsAccessTokenExpired is access token expired?
func (obj *ESI) IsAccessTokenExpired() bool {
	outdatedAt, _ := time.Parse("2006-01-02T15:04:05", obj.ExpiresOn)
	now := time.Now().UTC()
	return outdatedAt.Sub(now).Seconds() < 60
}

//RefreshAccessToken Refresh Access Token
func (obj *ESI) RefreshAccessToken() error {
	token, errAuth := OAuthToken(url.Values{"grant_type": {"refresh_token"}, "refresh_token": {obj.RefreshToken}})
	if errAuth != nil {
		fmt.Println("RefreshAccessToken/OAuthToken", errAuth)
		return errAuth
	}
	info, errVer := OAuthVerify(token.AccessToken)
	if errVer != nil {
		fmt.Println("RefreshAccessToken/OAuthVerify", errVer)
		return errVer
	}
	obj.AccessToken = token.AccessToken
	obj.ExpiresOn = info.ExpiresOn
	return nil
}

//Get get main object using given tokens
func Get(accessToken string, refreshToken string, expires string) ESI {
	result := ESI{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresOn:    expires,
	}

	if result.IsAccessTokenExpired() {
		result.RefreshAccessToken()
	}

	return result
}
