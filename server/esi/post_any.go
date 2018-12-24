package esi

//PostSimpleRequest just request
func PostSimpleRequest(url string, payload string, result interface{}) error {
	return post(url, payload, result)
}
