package sqrtx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	testCases := []struct {
		in     int
		result int
	}{
		{
			in:     0,
			result: 0,
		},
		{
			in:     1,
			result: 1,
		},
		{
			in:     2,
			result: 1,
		},
		{
			in:     4,
			result: 2,
		},
		{
			in:     8,
			result: 2,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i+1), func(t *testing.T) {
			assert.Equal(t, tc.result, mySqrt(tc.in))
		})
	}
}
