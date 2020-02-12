package esi

import (
	"ForumPublica/server/config"
	"bytes"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strings"
)

var (
	//UserAgent app user agent
	UserAgent = "Forum Publica 0.1"

	//OAuthTokenURL url
	OAuthTokenURL = "https://login.eveonline.com/oauth/token"

	//OAuthVerifyURL url
	OAuthVerifyURL = "https://login.eveonline.com/oauth/verify"

	//OAuthAuthorizeURL url
	OAuthAuthorizeURL = "https://login.eveonline.com/oauth/authorize"

	//ESIRootURL url
	ESIRootURL = "https://esi.evetech.net/latest"

	//Scopes app scopes
	Scopes = []string{
		"esi-assets.read_assets.v1",
		"esi-industry.read_character_jobs.v1",
		"esi-markets.structure_markets.v1",
		"esi-search.search_structures.v1",
		"esi-skills.read_skills.v1",
		"esi-ui.open_window.v1",
		"esi-universe.read_structures.v1",
		"esi-markets.read_character_orders.v1",
		"esi-wallet.read_character_wallet.v1",
	}
)

//OAuthTokenJSON model
type OAuthTokenJSON struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

//OAuthVerifyJSON model
type OAuthVerifyJSON struct {
	CharacterID        int64  `json:"CharacterID"`
	CharacterName      string `json:"CharacterName"`
	ExpiresOn          string `json:"ExpiresOn"`
	Scopes             string `json:"Scopes"`
	TokenType          string `json:"TokenType"`
	CharacterOwnerHash string `json:"CharacterOwnerHash"`
}

//Error model
type Error struct {
	Error string `json:"error"`
}

//CallbackURL form callback URL
func CallbackURL() (string, string) {
	b := make([]byte, 16)
	rand.Read(b)
	state := b64.URLEncoding.EncodeToString(b)

	url := fmt.Sprintf(
		"%s?response_type=code&redirect_uri=https%%3A%%2F%%2F%s%%3A%s%%2Fprobleme_callback&realm=ESI&client_id=%s&scope=%s&state=%s",
		OAuthAuthorizeURL,
		config.Vars.SITE,
		//coniootrs. url
		config.Vars.PORT,
		config.Vars.SSOClientID,
		strings.Join(Scopes, "%20"),
		state,
	)

	return url, state
}

//OAuthToken get token
func OAuthToken(data u.Values) (OAuthTokenJSON, error) {
	ssoString := b64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", config.Vars.SSOClientID, config.Vars.SSOSecretKey)),
	)
	auth := fmt.Sprintf("Basic %s", ssoString)
	postData := []byte(data.Encode())

	client := &http.Client{}
	req, err1 := http.NewRequest(
		"POST",
		OAuthTokenURL,
		bytes.NewReader(postData),
	)

	if err1 != nil {
		fmt.Println("OAuthToken err1", err1)
		return OAuthTokenJSON{}, err1
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("OAuthToken err2", err2)
		return OAuthTokenJSON{}, err2
	}
	defer resp.Body.Close()

	bodyBytes, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println("OAuthToken err3", err3)
		return OAuthTokenJSON{}, err3
	}

	result := OAuthTokenJSON{}
	err4 := json.Unmarshal(bodyBytes, &result)
	if err4 != nil {
		fmt.Println("OAuthToken err4", err4)
		return OAuthTokenJSON{}, err4
	}

	if result.Error != "" {
		return OAuthTokenJSON{}, errors.New(result.ErrorDescription)
	}

	return result, nil
}

//OAuthVerify verify token
func OAuthVerify(accessToken string) (OAuthVerifyJSON, error) {
	auth := fmt.Sprintf("Bearer %s", accessToken)

	req,
		err1 := http.NewRequest(
		"GET",

		//OAuthVerifyURL url
		OAuthVerifyURL,

		//
		bytes.NewReader([]byte{}),
	)

	if err1 != nil {
		fmt.Println(err1)
		return OAuthVerifyJSON{}, err1
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", auth)

	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return OAuthVerifyJSON{}, err2
	}
	defer resp.Body.Close()

	bodyBytes, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err3)
		return OAuthVerifyJSON{}, err3
	}

	result := OAuthVerifyJSON{}
	err4 := json.Unmarshal(bodyBytes, &result)
	if err4 != nil {
		fmt.Println(err4)
		return OAuthVerifyJSON{}, err4
	}

	return result, nil
}
