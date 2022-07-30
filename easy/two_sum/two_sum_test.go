package two_sum

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		input, pairIndexes []int
		target             int
	}{
		{input: []int{3, 3}, pairIndexes: []int{0, 1}, target: 6},
		{input: []int{2, 7, 11, 15}, pairIndexes: []int{0, 1}, target: 9},
		{input: []int{3, 2, 4}, pairIndexes: []int{1, 2}, target: 6},
		{input: []int{-5, 0, 7}, pairIndexes: []int{0, 2}, target: 2},
	}

	for _, test := range testCases {
		test := test

		t.Run(fmt.Sprintf("input: %v | result: %v", test.input, test.pairIndexes), func(t *testing.T) {
			result := TwoSum(test.input, test.target)
			require.Equal(t, test.pairIndexes, result)

			resultSlow := TwoSumSlow(test.input, test.target)
			require.Equal(t, test.pairIndexes, resultSlow)
		})
	}
}
