package main

import (
	"errors"
	"regexp"
	"strconv"
)

func CalcString(i string) (result int, err error) {
	result = 0
	reg := regexp.MustCompile(`[,\n]`)
	subStrings := reg.Split(i, -1)
	// subStrings := strings.Split(i, ",")

	for _, s := range subStrings {
		if s != "" {
			value, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			if value < 0 {
				return 0, errors.New("bad")
			}
			if value <= 1000 {
				result += value
			}
		}
	}
	return result, nil
}
