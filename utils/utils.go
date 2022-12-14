package utils

import (
	"fmt"
	"strings"
)

var maxSeconds int64 = 922337203685
var minSeconds int64 = -922337203685
var ticksPerSecond int64 = 10000000

func IsLeapYear(year int) bool {
	if year < 1 || year > 9999 {
		return false
	}
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func TimeToTicks(hour, minute, second int) (int64, error) {
	totalSeconds := int64(hour)*3600 + int64(minute)*60 + int64(second)
	if totalSeconds > maxSeconds || totalSeconds < minSeconds {
		return 0, fmt.Errorf("totalseconds:%d is out of range", totalSeconds)
	}
	return totalSeconds * ticksPerSecond, nil
}
func ConvertLayout(cslayout string) string {
	//yyyy-MM-dd HH:mm:ss
	//2006-01-02 15:04:05
	//year
	if strings.Contains(cslayout, "yyyy") {
		cslayout = strings.ReplaceAll(cslayout, "yyyy", "2006")
	}
	if strings.Contains(cslayout, "yy") {
		cslayout = strings.ReplaceAll(cslayout, "yy", "06")
	}
	if strings.Contains(cslayout, "y") {
		cslayout = strings.ReplaceAll(cslayout, "y", "6")
	}
	//Moth
	if strings.Contains(cslayout, "MM") {
		cslayout = strings.ReplaceAll(cslayout, "MM", "01")
	}
	if strings.Contains(cslayout, "M") {
		cslayout = strings.ReplaceAll(cslayout, "M", "1")
	}
	//DAY
	if strings.Contains(cslayout, "dd") {
		cslayout = strings.ReplaceAll(cslayout, "dd", "02")
	}
	if strings.Contains(cslayout, "d") {
		cslayout = strings.ReplaceAll(cslayout, "d", "2")
	}
	//Hour
	if strings.Contains(cslayout, "HH") {
		cslayout = strings.ReplaceAll(cslayout, "HH", "15")
	}
	if strings.Contains(cslayout, "H") {
		cslayout = strings.ReplaceAll(cslayout, "H", "5")
	}
	//Minute
	if strings.Contains(cslayout, "mm") {
		cslayout = strings.ReplaceAll(cslayout, "mm", "04")
	}
	if strings.Contains(cslayout, "m") {
		cslayout = strings.ReplaceAll(cslayout, "m", "4")
	}
	//Second
	if strings.Contains(cslayout, "ss") {
		cslayout = strings.ReplaceAll(cslayout, "ss", "05")
	}
	if strings.Contains(cslayout, "s") {
		cslayout = strings.ReplaceAll(cslayout, "d", "5")
	}
	//millisecond
	if strings.Contains(cslayout, "f") {
		cslayout = strings.ReplaceAll(cslayout, "f", "0")
	}
	return cslayout
}
