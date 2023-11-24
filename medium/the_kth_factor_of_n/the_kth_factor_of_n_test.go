package the_kth_factor_of_n

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	testCases := []struct {
		in, in2 int
		result  int
	}{
		{
			in:     12,
			in2:    3,
			result: 3,
		},
		{
			in:     7,
			in2:    2,
			result: 7,
		},
		{
			in:     4,
			in2:    4,
			result: -1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i+1), func(t *testing.T) {
			assert.Equal(t, tc.result, kthFactor(tc.in, tc.in2))
		})
	}
}
