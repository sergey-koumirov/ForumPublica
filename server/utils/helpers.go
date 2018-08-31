package utils

import (
	"encoding/json"
	"html/template"
)

func TimeoutClass(mStr string) string {
	if mStr == "00:00" {
		return "orange"
	} else {
		return "gray"
	}

}

func Marshal(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}
