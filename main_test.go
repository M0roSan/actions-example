package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		name     string
		input1   int
		input2   int
		expected int
	}{
		{
			name:     "positive",
			input1:   1,
			input2:   2,
			expected: 3,
		},
		{
			name:     "negative",
			input1:   -100,
			input2:   -10,
			expected: -110,
		},
		{
			name:     "mix",
			input1:   -10,
			input2:   10,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := sum(test.input1, test.input2)
			a.Equal(test.expected, actual)
		})
	}
}
