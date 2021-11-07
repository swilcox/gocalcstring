package main

import (
	"errors"
	"testing"
)

func TestCalcString(t *testing.T) {
	testCases := []struct {
		InputString   string
		ExpectedValue int
		ExpectedError error
		Description   string
	}{
		{"", 0, nil, "Empty String"},
		{"42", 42, nil, "Single Value"},
		{"1,2", 3, nil, "Comma-separated"},
		{"42\n1", 43, nil, "Newline-separated"},
		{"20\n21,23", 64, nil, "multiple separators"},
		{"-1,-2", 0, errors.New("negatives not allowed [-1, -2]"), "negatives should throw error"},
		{"1001,1", 1, nil, "ignore values greater than 1000"},
		{"//;\n1;2", 3, nil, "Custom separator"},
		{"//;*\n1;2*3", 6, nil, "Multiple custom separators"},
		{"//[***]\n1***2***3", 6, nil, "Custom Length separator"},
		{"//[**][--]\n1**2--3", 6, nil, "Multiple Custom length separators"},
	}
	for _, tCase := range testCases {
		t.Run(tCase.Description, func(t *testing.T) {
			result, err := CalcString(tCase.InputString)
			if result != tCase.ExpectedValue || (tCase.ExpectedError == nil && err != nil) || (tCase.ExpectedError != nil && (err == nil || err.Error() != tCase.ExpectedError.Error())) {
				t.Errorf("SUB-TEST: %q\nexpected result: %v and error: %q but received result: %v and error: %q", tCase.Description, tCase.ExpectedValue, tCase.ExpectedError, result, err)
			}
		})
	}
}
