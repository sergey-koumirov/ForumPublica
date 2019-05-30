package esi

//PostSimpleRequest just request
func PostSimpleRequest(url string, payload string, result interface{}) error {
	_, _, err := post(url, payload, result)
	return err
}
