package mymath

import (
	"testing"
)

type testCases []struct {
	input    float64
	expected float64
}

func TestAbs(t *testing.T) {
	testCasesAbs := testCases{
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: -1, expected: 1},
		{input: 2, expected: 2},
		{input: -2, expected: 2},
	}
	for _, tc := range testCasesAbs {
		actual := Abs(tc.input)
		if actual != tc.expected {
			t.Errorf("Abs(%v): expected %v, actual %v", tc.input, tc.expected, actual)
		}
	}
}

func TestAcos(t *testing.T) {
	testCasesAcos := testCases{
		{input: 0, expected: 1.5707963267948966},
		{input: 1, expected: 0},
		{input: -1, expected: 3.141592653589793},
	}
	for _, tc := range testCasesAcos {
		actual := Acos(tc.input)
		if actual != tc.expected {
			t.Errorf("Acos(%v): expected %v, actual %v", tc.input, tc.expected, actual)
		}
	}
}

func TestAcosh(t *testing.T) {
	testCasesAcosh := testCases{
		{input: 1, expected: 0},
		{input: 2, expected: 1.3169578969248166},
		{input: 3, expected: 1.7627471740390859},
	}
	for _, tc := range testCasesAcosh {
		actual := Acosh(tc.input)
		if actual != tc.expected {
			t.Errorf("Acosh(%v): expected %v, actual %v", tc.input, tc.expected, actual)
		}
	}
}

func TestAsin(t *testing.T) {
	testCasesAsin := testCases{
		{input: 0, expected: 0},
		{input: 1, expected: 1.5707963267948966},
		{input: -1, expected: -1.5707963267948966},
	}
	for _, tc := range testCasesAsin {
		actual := Asin(tc.input)
		if actual != tc.expected {
			t.Errorf("Asin(%v): expected %v, actual %v", tc.input, tc.expected, actual)
		}
	}
}
