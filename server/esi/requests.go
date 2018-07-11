package esi

import (
    "net/http"
    "fmt"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "errors"
    "time"
    "strconv"
)

var TRY int = 5;

func get(url string, result interface{}) error {
  return request("GET", url, "", result)
}

func post(url string,  payload string, result interface{}) error {
  return request("POST", url, payload, result)
}

func request(method string, url string, payload string, result interface{}) error {
    client := &http.Client{}

    req, errReq := http.NewRequest(method, url, bytes.NewBufferString(payload))
    if errReq!=nil {
        fmt.Println(errReq)
        return errReq
    }

    req.Header.Add("User-Agent", USER_AGENT)
    req.Header.Add("Accept", "application/json")

    var(
        resp *http.Response
        errDo error
    )
    i := 0
    for {
        if i >= TRY { break }
        i++;

        resp, errDo = client.Do(req)
        defer resp.Body.Close()

        if errDo!=nil {
            return errDo
        }else{
            if resp.StatusCode == 500 {
                fmt.Println(i, resp.StatusCode, url)
                time.Sleep(1000 * time.Millisecond)
            }else{
                break
            }
        }
    }

    bodyBytes, errRead := ioutil.ReadAll(resp.Body)
    if errRead!=nil {
        return errRead
    }

    //fmt.Println( url )
    //fmt.Println( string(bodyBytes) )
    //fmt.Println( "Status: ", resp.StatusCode )

    if resp.StatusCode == 200 {
        errUn := json.Unmarshal(bodyBytes, result)
        if errUn!=nil {
            return errUn
        }
    }else{
        esiError := EsiError{}
        errUn := json.Unmarshal(bodyBytes, &esiError)
        if errUn!=nil {
            return errUn
        }
        return errors.New(esiError.Error)
    }
    return nil
}

func auth(method string, accessToken string, url string, result interface{}) (time.Time, int64, error) {
    expires := time.Now()
    pages := int64(1)

    client := &http.Client{}
    req, errReq := http.NewRequest(method, url, bytes.NewReader([]byte{}))
    if errReq!=nil {
        fmt.Println(errReq)
        return expires, pages, errReq
    }

    req.Header.Add("User-Agent", USER_AGENT)
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
    req.Header.Add("Accept", "application/json")

    var(
        resp *http.Response
        errDo error
    )
    i := 0
    for {
        if i >= TRY {
          return expires, pages, errors.New("Try limit exceeded")
        }
        i++;

        resp, errDo = client.Do(req)
        defer resp.Body.Close()

        if errDo!=nil {
            return expires, pages, errDo
        }else{
            if resp.StatusCode == 500 {
                fmt.Println(i, resp.StatusCode, url)
                time.Sleep(1000 * time.Millisecond)
            }else{
                break
            }
        }
    }

    if len(resp.Header["Expires"]) > 0 {
        expires, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 MST", resp.Header["Expires"][0])
    }

    if len(resp.Header["X-Pages"]) > 0 {
        pages, _ = strconv.ParseInt(resp.Header["X-Pages"][0], 10, 64)
    }

    bodyBytes, errRead := ioutil.ReadAll(resp.Body)
    if errRead!=nil {
        return expires, pages, errRead
    }

    //fmt.Println( url )
    //fmt.Println( string(bodyBytes) )
    //fmt.Println( "Status: ", resp.StatusCode )

    if resp.StatusCode == 200 {
        errUn := json.Unmarshal(bodyBytes, result)
        if errUn != nil {
            return expires, pages, errUn
        }
    }else if resp.StatusCode == 204 {
        return expires, pages, nil
    }else{
        esiError := EsiError{}
        errUn := json.Unmarshal(bodyBytes, &esiError)
        if errUn!=nil {
            return expires, pages, errUn
        }
        return expires, pages, errors.New(esiError.Error)
    }
    return expires, pages, nil
}