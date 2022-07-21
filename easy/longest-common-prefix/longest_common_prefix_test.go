package two_sum

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		input  []string
		result string
	}{
		{input: []string{"flower", "flow", "flight"}, result: "fl"},
		{input: []string{"dog", "racecar", "car"}, result: ""},
		{input: []string{"", "b"}, result: ""},
		{input: []string{"ba", "b"}, result: "b"},
		{input: []string{"", ""}, result: ""},
		{input: []string{"equal", "equal "}, result: "equal"},
		{input: []string{" equal", "equal "}, result: ""},
	}

	for _, test := range testCases {
		test := test

		t.Run(fmt.Sprintf("input: %v | result: %v", test.input, test.result), func(t *testing.T) {
			result := LongestCommonPrefix(test.input)
			require.Equal(t, test.result, result)
		})
	}
}
