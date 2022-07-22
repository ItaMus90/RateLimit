package service

import (
	"flag"
	"strconv"
)

func ConvertStringToInt(str string) (int, bool) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}

	return num, true
}

func CheckInputArguments() (bool, int) {
	var threshold string
	var ttl string

	flag.StringVar(&threshold, "threshold", "2", "Max number of requests per URL within a time period")
	flag.StringVar(&ttl, "ttl", "10", "The time period in which URL visits will be counted.")

	flag.Parse()

	thresholdNum, isValidThreshold := ConvertStringToInt(threshold)
	ttlNum, isValidTtl := ConvertStringToInt(ttl)

	if !isValidThreshold || !isValidTtl {
		return false, 0
	}

	InputData.Threshold = thresholdNum
	InputData.Ttl = ttlNum

	return true, ttlNum
}
