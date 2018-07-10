package esi

import (
    "time"
    "net/url"
    "fmt"
)

type ESI struct{
    AccessToken  string
    RefreshToken string
    ExpiresOn    string
}

func (obj *ESI) IsAccessTokenExpired() bool {
    outdatedAt, _ := time.Parse("2006-01-02T15:04:05", obj.ExpiresOn)
    now := time.Now().UTC()
    return outdatedAt.Sub(now).Seconds() < 60
}

func (obj *ESI) RefreshAccessToken() {
    token, errAuth := OAuthToken(url.Values{"grant_type": {"refresh_token"}, "refresh_token": {obj.RefreshToken}})
    if errAuth!=nil {
        fmt.Println(errAuth)
        return
    }
    info, errVer := OAuthVerify(token.AccessToken)
    if errVer!=nil {
        fmt.Println(errVer)
        return
    }
    obj.AccessToken = token.AccessToken
    obj.ExpiresOn = info.ExpiresOn
}

func Get(accessToken string, refreshToken  string, expires  string) ESI {
    result := ESI{
        AccessToken: accessToken,
        RefreshToken: refreshToken,
        ExpiresOn: expires,
    }

    if result.IsAccessTokenExpired() {
        result.RefreshAccessToken()
    }

    return result
}
