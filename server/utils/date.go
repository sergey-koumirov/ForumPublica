package utils

import "time"

//NowUTCStr now() as string
func NowUTCStr() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}
