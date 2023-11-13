package add_binary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	testCases := []struct {
		in     string
		in2    string
		result string
	}{
		{
			in:     "11",
			in2:    "1",
			result: "100",
		},
		{
			in:     "1010",
			in2:    "1011",
			result: "10101",
		},
		{
			in:     "1",
			in2:    "0",
			result: "1",
		},
		{
			in:     "1111",
			in2:    "1111",
			result: "11110",
		},
		{
			in:     "100",
			in2:    "110010",
			result: "110110",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i+1), func(t *testing.T) {
			assert.Equal(t, tc.result, addBinary(tc.in, tc.in2))
		})
	}
}
