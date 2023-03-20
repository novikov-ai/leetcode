package mean_of_array_after_removing_some_elements

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimMean(t *testing.T) {
	testCases := []struct {
		in     []int
		result float64
	}{
		{
			in:     []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3},
			result: 2.00,
		},
		{
			in:     []int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0},
			result: 4.00,
		},
		{
			in:     []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4},
			result: 4.77778,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i+1), func(t *testing.T) {
			accuracy := 100000.00
			assert.Equal(t, testCase.result, math.Round(TrimMean(testCase.in)*accuracy)/accuracy)
		})
	}
}
