package multiplystrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiply(t *testing.T) {
	tests := []struct {
		name   string
		num1   string
		num2   string
		result string
	}{
		{
			name:   "one number",
			num1:   "2",
			num2:   "3",
			result: "6",
		},
		{
			name:   "zero",
			num1:   "0",
			num2:   "3",
			result: "0",
		},
		{
			name:   "zero",
			num1:   "123",
			num2:   "0",
			result: "0",
		},
		{
			name:   "long num1",
			num1:   "123",
			num2:   "2",
			result: "246",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := multiply(tc.num1, tc.num2)

			assert.Equal(t, tc.result, got)
		})
	}
}
