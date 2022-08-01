package remove_element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input  []int
		delete int
		left   int
	}{
		{input: []int{1, 2, 3, 3, 4, 3}, delete: 3, left: 3},
		{input: []int{0, 2, 4, 6, 7, 7, 8}, delete: 7, left: 5},
	}

	for _, test := range testCases {
		test := test
		t.Run("", func(t *testing.T) {
			actual := RemoveElement(test.input, test.delete)
			require.Equal(t, test.left, actual)
		})
	}
}
