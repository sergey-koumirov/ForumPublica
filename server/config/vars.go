package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//VarsValues var values from file
type VarsValues struct {
	MODE         string
	PORT         string
	SITE         string
	SSOClientID  string
	SSOSecretKey string
	DBC          string
	SessionKey   string
	SDE          string
	SSLPath      string
}

//Vars global variable
var Vars *VarsValues

//LoadVars load vars
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
