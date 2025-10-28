package dexcom

import (
	"regexp"
	"strconv"
)

func parseDate(dateString string) int64 {
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(dateString, "")

	num, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		panic(err)
	}

	return num
}
