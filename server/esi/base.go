package esi

import (
	"ForumPublica/server/config"
	"bytes"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strings"
)

var USER_AGENT = "Forum Publica 0.1"
var OAUTH_TOKEN_URL string = "https://login.eveonline.com/oauth/token"
var OAUTH_VERIFY_URL string = "https://login.eveonline.com/oauth/verify"
var OAUTH_AUTHORIZE_URL string = "https://login.eveonline.com/oauth/authorize"
var ESI_ROOT_URL string = "https://esi.tech.ccp.is/latest"
var SCOPES = []string{
	"esi-skills.read_skills.v1",
	"esi-search.search_structures.v1",
	"esi-universe.read_structures.v1",
	"esi-markets.structure_markets.v1",
	"esi-assets.read_assets.v1",
	"esi-industry.read_character_jobs.v1",
}

type OAuthTokenJson struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type OAuthVerifyJson struct {
	CharacterID        int64  `json:"CharacterID"`
	CharacterName      string `json:"CharacterName"`
	ExpiresOn          string `json:"ExpiresOn"`
	Scopes             string `json:"Scopes"`
	TokenType          string `json:"TokenType"`
	CharacterOwnerHash string `json:"CharacterOwnerHash"`
}

type EsiError struct {
	Error string `json:"error"`
}

func CallbackURL() (string, string) {
	b := make([]byte, 16)
	rand.Read(b)
	state := b64.URLEncoding.EncodeToString(b)

	url := fmt.Sprintf(
		"%s?response_type=code&redirect_uri=http%%3A%%2F%%2F%s%%3A%s%%2Fprobleme_callback&realm=ESI&client_id=%s&scope=%s&state=%s",
		OAUTH_AUTHORIZE_URL,
		config.Vars.SITE,
		config.Vars.PORT,
		config.Vars.SSOClientID,
		strings.Join(SCOPES, "%20"),
		state,
	)

	return url, state
}

func OAuthToken(data u.Values) (OAuthTokenJson, error) {
	ssoString := b64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", config.Vars.SSOClientID, config.Vars.SSOSecretKey)),
	)
	auth := fmt.Sprintf("Basic %s", ssoString)
	postData := []byte(data.Encode())

	client := &http.Client{}
	req, err1 := http.NewRequest(
		"POST",
		OAUTH_TOKEN_URL,
		bytes.NewReader(postData),
	)

	if err1 != nil {
		fmt.Println(err1)
		return OAuthTokenJson{}, err1
	}

	req.Header.Add("User-Agent", USER_AGENT)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err2 := client.Do(req)
	defer resp.Body.Close()

	if err2 != nil {
		fmt.Println(err2)
		return OAuthTokenJson{}, err2
	}

	bodyBytes, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err3)
		return OAuthTokenJson{}, err3
	}

	result := OAuthTokenJson{}
	err4 := json.Unmarshal(bodyBytes, &result)
	if err4 != nil {
		fmt.Println(err4)
		return OAuthTokenJson{}, err4
	}

	return result, nil
}

func OAuthVerify(accessToken string) (OAuthVerifyJson, error) {
	auth := fmt.Sprintf("Bearer %s", accessToken)

	req, err1 := http.NewRequest(
		"GET",
		OAUTH_VERIFY_URL,
		bytes.NewReader([]byte{}),
	)
	if err1 != nil {
		fmt.Println(err1)
		return OAuthVerifyJson{}, err1
	}

	req.Header.Add("User-Agent", USER_AGENT)
	req.Header.Add("Authorization", auth)

	client := &http.Client{}
	resp, err2 := client.Do(req)
	defer resp.Body.Close()
	if err2 != nil {
		fmt.Println(err2)
		return OAuthVerifyJson{}, err2
	}

	bodyBytes, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err3)
		return OAuthVerifyJson{}, err3
	}

	result := OAuthVerifyJson{}
	err4 := json.Unmarshal(bodyBytes, &result)
	if err4 != nil {
		fmt.Println(err4)
		return OAuthVerifyJson{}, err4
	}

	return result, nil
}
