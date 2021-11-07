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

		if delimiterArray != nil {
			delimiters := []string{}
			for _, match := range delimiterArray {
				delimiters = append(delimiters, regexp.QuoteMeta(match[1]))
			}
			reg = regexp.MustCompile(strings.Join(delimiters, "|"))
		} else {
			reg = regexp.MustCompile(regexp.QuoteMeta(matchArray[1]))
		}

		compStr = strings.SplitN(i, "\n", 2)[1]
	}
	subStrings := reg.Split(compStr, -1)

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
