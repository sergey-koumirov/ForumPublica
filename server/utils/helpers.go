package utils

func TimeoutClass(mStr string) string {
	if mStr == "00:00" {
		return "orange"
	} else {
		return "gray"
	}

}
