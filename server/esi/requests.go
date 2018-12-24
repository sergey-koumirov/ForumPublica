package esi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//TRY numbers of tries if request fail
var TRY = 5

func get(url string, result interface{}) error {
	return request("GET", url, "", result)
}

func post(url string, payload string, result interface{}) error {
	return request("POST", url, payload, result)
}

func request(method string, url string, payload string, result interface{}) error {
	client := &http.Client{}

	req, errReq := http.NewRequest(method, url, bytes.NewBufferString(payload))
	if errReq != nil {
		fmt.Println(errReq)
		return errReq
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Accept", "application/json")

	var (
		resp  *http.Response
		errDo error
	)
	i := 0
	for {
		if i >= TRY {
			break
		}
		i++

		resp, errDo = client.Do(req)
		if errDo != nil {
			return errDo
		}
		defer resp.Body.Close()

		if resp.StatusCode == 500 {
			fmt.Println(i, resp.StatusCode, url)
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}

	}

	bodyBytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return errRead
	}

	//fmt.Println( url )
	//fmt.Println( string(bodyBytes) )
	//fmt.Println( "Status: ", resp.StatusCode )

	if resp.StatusCode == 200 {
		errUn := json.Unmarshal(bodyBytes, result)
		if errUn != nil {
			return errUn
		}
	} else {
		esiError := Error{}
		errUn := json.Unmarshal(bodyBytes, &esiError)
		if errUn != nil {
			return errUn
		}
		return errors.New(esiError.Error)
	}
	return nil
}

func auth(method string, accessToken string, url string, result interface{}) (time.Time, int64, error) {
	return authRequest(method, accessToken, url, "", result)
}

func authRequest(method string, accessToken string, url string, payload string, result interface{}) (time.Time, int64, error) {
	expires := time.Now()
	pages := int64(1)

	client := &http.Client{}
	req, errReq := http.NewRequest(method, url, bytes.NewBufferString(payload))
	if errReq != nil {
		fmt.Println(errReq)
		return expires, pages, errReq
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Accept", "application/json")

	var (
		resp  *http.Response
		errDo error
	)
	i := 0
	for {
		if i >= TRY {
			return expires, pages, errors.New("Try limit exceeded")
		}
		i++

		resp, errDo = client.Do(req)
		if errDo != nil {
			fmt.Println("errDo", errDo)
		}
		defer resp.Body.Close()

		if errDo != nil {
			return expires, pages, errDo
		}
		if resp.StatusCode == 500 {
			fmt.Println(i, resp.StatusCode, url)
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}

	if len(resp.Header["Expires"]) > 0 {
		expires, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", resp.Header["Expires"][0])
	}

	if len(resp.Header["X-Pages"]) > 0 {
		pages, _ = strconv.ParseInt(resp.Header["X-Pages"][0], 10, 64)
	}

	bodyBytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return expires, pages, errRead
	}

	//fmt.Println( url )
	//fmt.Println(string(bodyBytes))
	//fmt.Println( "Status: ", resp.StatusCode )

	if resp.StatusCode == 200 {
		errUn := json.Unmarshal(bodyBytes, result)
		if errUn != nil {
			return expires, pages, errUn
		}
	} else if resp.StatusCode == 204 {
		return expires, pages, nil
	} else {
		esiError := Error{}
		errUn := json.Unmarshal(bodyBytes, &esiError)
		if errUn != nil {
			return expires, pages, errUn
		}
		return expires, pages, errors.New(esiError.Error)
	}
	return expires, pages, nil
}
