package main

import "testing"

func TestCalcStringEmpty(t *testing.T) {
	result, err := CalcString("")
	if result != 0 || err != nil {
		t.Errorf("expected 0 result and nil error from empty string but got: %q, %q", result, err)
	}
}

func TestCalcStringNumber(t *testing.T) {
	result, err := CalcString("42")
	if result != 42 || err != nil {
		t.Errorf("expected 42 result and nil error from single number string but got: %q, %q", result, err)
	}
}

func TestCalcStringCommas(t *testing.T) {
	result, err := CalcString("42,1")
	if result != 43 || err != nil {
		t.Errorf("expected 43 result and nil error from comma delimited string but got: %q, %q", result, err)
	}
}

func TestCalcStringNewLine(t *testing.T) {
	result, err := CalcString("42\n1")
	if result != 43 || err != nil {
		t.Errorf("expected 43 result and nil error from newline delimited string but got: %q, %q", result, err)
	}
}

func TestCalcStringMixedDelimiters(t *testing.T) {
	result, err := CalcString("20\n21,23")
	if result != 64 || err != nil {
		t.Errorf("expected 64 result and nil error from 3 number mixed delimiter string but got: %q, %q", result, err)
	}
}

func TestCalcStringNegative(t *testing.T) {
	result, err := CalcString("-1")
	if result != 0 || err == nil {
		t.Errorf("expected nil result and error from negative number string but got: %q, %q", result, err)
	}
}

func TestCalcStringIgnoreGreater1000(t *testing.T) {
	result, err := CalcString("42\n1001")
	if result != 42 || err != nil {
		t.Errorf("expected 42 result and nil error from string with greater than 1000 in it but got: %q, %q", result, err)
	}
}

func TestCalcStringCustomDelimiter(t *testing.T) {
	result, err := CalcString("//;\n1;2")
	if result != 3 || err != nil {
		t.Errorf("expected 3 result and nil error from string with custom ; delimiter: %q, %q", result, err)
	}
}

func TestCalcStringCustomLengthDelimiter(t *testing.T) {
	result, err := CalcString("//[***]\n1***2***3")
	if result != 6 || err != nil {
		t.Errorf("expected 6 result and nil error from string with custom *** delimiter: %q, %q", result, err)
	}
}

func TestCalcStringCustomLengthMultipleDelimiters(t *testing.T) {
	result, err := CalcString("//[*][%]\n1*2%3")
	if result != 6 || err != nil {
		t.Errorf("expected 6 result and nil error from string with custom delimiters: %q, %q", result, err)
	}
}
