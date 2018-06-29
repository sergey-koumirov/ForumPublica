package config

import (
	"fmt"
  "io/ioutil"
  "encoding/json"
)

type VarsValues struct{
	MODE string
  PORT string
  SITE string
  SSOClientID  string
  SSOSecretKey string
}


var Vars *VarsValues = nil

func LoadVars(){
  data, rErr := ioutil.ReadFile("server/vars.json")
  if rErr != nil {
    fmt.Println("ReadFile",rErr)
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
