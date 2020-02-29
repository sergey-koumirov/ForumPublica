package utils

import (
	"ForumPublica/sde/static"
	"encoding/json"
	"html/template"
)

//TimeoutClass func
func TimeoutClass(mStr string) string {
	if mStr == "00:00" {
		return "orange"
	} else {
		return "gray"
	}

}

//Marshal func
func Marshal(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

//TypeName func
func TypeName(id int32) string {
	return static.Types[id].Name
}

func FormatFloat(in float64) string {
	return RenderFloat("# ###.##", in)
}

func FormatInt(in int64) string {
	return RenderFloat("# ###.", float64(in))
}
