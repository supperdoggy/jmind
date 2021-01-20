package main

import (
	"fmt"
	"testing"
)

func TestGetBlockByNumber(t *testing.T) {
	type testCase struct {
		expectedResult int
		block          int64
	}

	cases := []testCase{
		{expectedResult: 241, block: 11508993},
		{expectedResult: 155, block: 11509797},
		{expectedResult: 1, block: 109789},
		{expectedResult: 0, block: 0},
	}

	for _, v := range cases {
		t.Run(fmt.Sprintf("%v", v.block), func(t *testing.T) {
			result, err := getBlockByNumber(v.block)
			if err != nil || len(result.R.Transactions) != v.expectedResult {
				t.Errorf("expected: %v %v, got: %v %v", v.expectedResult, nil, result, err)
			}
		})
	}
}
