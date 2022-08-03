package plus_one

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		input  []int
		result []int
	}{
		{input: []int{1, 2, 3}, result: []int{1, 2, 4}},
		{input: []int{4, 3, 2}, result: []int{4, 3, 3}},
		{input: []int{9}, result: []int{1, 0}},
	}

	for _, test := range testCases {
		test := test

		t.Run(fmt.Sprintf("input: %v | result: %v", test.input, test.result), func(t *testing.T) {
			result := PlusOne(test.input)
			require.Equal(t, test.result, result)
		})
	}
}
