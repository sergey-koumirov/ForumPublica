package utils

import (
	"fmt"
	"math"
	"time"
)

func StrToMinut(str string) int64 {
	now_str := time.Now().UTC().Format("2006-01-02 15:04:05")
	last, _ := time.Parse("2006-01-02T15:04:05Z", str)
	now, _ := time.Parse("2006-01-02 15:04:05", now_str)

	return int64(math.Floor(now.Sub(last).Minutes()))
}

func FormatToHM(str string) string {

	minutes := float64(StrToMinut(str))
	if minutes > 0 {
		minutes = 0
	} else {
		minutes = math.Abs(minutes)
	}

	days := math.Floor(minutes / (24 * 60))

	if days > 0 {
		return fmt.Sprintf("%2.0fd %02.0f:%02.0f", days, math.Mod(math.Floor(minutes/60), 24), math.Mod(minutes, 60))
	} else {
		return fmt.Sprintf("%02.0f:%02.0f", math.Floor(minutes/60), math.Mod(minutes, 60))
	}

}
