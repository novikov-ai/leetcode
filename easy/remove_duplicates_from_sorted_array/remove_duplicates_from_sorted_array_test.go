package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input  []int
		unique int
	}{
		{input: []int{1, 1, 1, 1, 1, 1}, unique: 1},
		{input: []int{1, 1, 2}, unique: 2},
		{input: []int{0, 0, 1, 1, 2}, unique: 3},
		{input: []int{-5, -5, -3, -2, -1, -1, 0, 0, 1, 1, 2, 3, 3, 4}, unique: 9},
	}

	for _, test := range testCases {
		test := test
		t.Run("", func(t *testing.T) {
			actual := RemoveDuplicates(test.input)
			require.Equal(t, test.unique, actual)
		})
	}
}
