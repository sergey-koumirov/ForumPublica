package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type VarsValues struct {
	MODE         string
	PORT         string
	SITE         string
	SSOClientID  string
	SSOSecretKey string
	DBC          string
	SESSION_KEY  string
	SDE          string
}

var Vars *VarsValues = nil

func LoadVars() {
	data, rErr := ioutil.ReadFile("server/vars.json")
	if rErr != nil {
		fmt.Println("ReadFile", rErr)
		return
	}

	temp := VarsValues{}
	uErr := json.Unmarshal([]byte(data), &temp)
	if uErr != nil {
		fmt.Println("Unmarshal", uErr)
		return
	}
	Vars = &temp

}
