package utils

import (
	"fmt"
	"math"
	"time"
)

//DbStrToMinut covert "2006-01-02T15:04:05Z" time str to minutes from now()
func DbStrToMinut(str string) int64 {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	last, err1 := time.Parse("2006-01-02 15:04:05", str)
	if err1 != nil {
		fmt.Println(err1)
	}
	now, err2 := time.Parse("2006-01-02 15:04:05", nowStr)
	if err2 != nil {
		fmt.Println(err2)
	}
	return int64(math.Floor(now.Sub(last).Minutes()))
}

//StrToMinut covert "2006-01-02T15:04:05Z" time str to minutes from now()
func StrToMinut(str string) int64 {
	nowStr := time.Now().UTC().Format("2006-01-02 15:04:05")
	last, err1 := time.Parse("2006-01-02T15:04:05Z", str)
	if err1 != nil {
		fmt.Println(err1)
	}
	now, err2 := time.Parse("2006-01-02 15:04:05", nowStr)
	if err2 != nil {
		fmt.Println(err2)
	}
	return int64(math.Floor(now.Sub(last).Minutes()))
}

//FormatToHM format date as HH:MM from now
func FormatToHM(str string) string {
	return FormatToHMPlus(str, 0)
}

func FormatToHMPlus(str string, plus int64) string {
	minutes := float64(StrToMinut(str) - plus)

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
