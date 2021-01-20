package main

import (
	"strconv"
	"testing"
)

func TestIntToHexa(t *testing.T) {
	type testCase struct {
		expectedResult string
		number         int
	}

	tests := []testCase{
		{number: 0, expectedResult: "0"},
		{number: 10, expectedResult: "a"},
		{number: 3132321, expectedResult: "2fcba1"},
		{number: -1, expectedResult: "-1"},
		{number: -12312332, expectedResult: "-bbdf0c"},
	}

	for _, v := range tests {
		t.Run(strconv.Itoa(v.number), func(t *testing.T) {
			result := intToHexa(v.number)
			if v.expectedResult != result {
				t.Errorf("wanted:%v, got:%v\n", v.expectedResult, result)
			}
		})
	}

}

func TestFromWei(t *testing.T) {
	type testCase struct {
		expectedResult string
		value          string
	}

	tests := []testCase{
		{value: "0x38d7ea4c68000", expectedResult: "0.001"},
		{value: "0x137e9165c0e3000", expectedResult: "0.087795"},
		{value: "0X0", expectedResult: "0"},
		{value: "", expectedResult: ""},
		{value: "12", expectedResult: ""},
		{value: "0x38d7ea4c6800038d7ea4c68000", expectedResult: "4.503599627370497e+12"},
	}

	for _, v := range tests {
		t.Run(v.value, func(t *testing.T) {
			result := FromWei(v.value)
			if v.expectedResult != result {
				t.Errorf("wanted:%v, got:%v\n", v.expectedResult, result)
			}
		})
	}
}
