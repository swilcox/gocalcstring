package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func CalcString(i string) (result int, err error) {
	result = 0
	compStr := i
	reg := regexp.MustCompile(`,|\n`)
	delreg := regexp.MustCompile(`^//(?P<delimiters>.+)\n`)
	matchArray := delreg.FindStringSubmatch(i)
	if matchArray != nil {
		subReg := regexp.MustCompile(`\[(.+?)]`)
		delimiterArray := subReg.FindAllStringSubmatch(matchArray[1], -1)
		delimiters := []string{}
		if delimiterArray != nil {
			for _, match := range delimiterArray {
				delimiters = append(delimiters, regexp.QuoteMeta(match[1]))
			}
		} else {
			for _, chr := range matchArray[1] {
				delimiters = append(delimiters, regexp.QuoteMeta(string(chr)))
			}
		}
		reg = regexp.MustCompile(strings.Join(delimiters, "|"))
		compStr = strings.SplitN(i, "\n", 2)[1]
	}
	badValues := []string{}
	subStrings := reg.Split(compStr, -1)

	for _, s := range subStrings {
		if s != "" {
			value, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			if value < 0 {
				badValues = append(badValues, strconv.Itoa(value))
			}
			if value <= 1000 {
				result += value
			}
		}
	}
	if len(badValues) > 0 {
		return 0, errors.New("negatives not allowed [" + strings.Join(badValues, ", ") + "]")
	}
	return result, nil
}
