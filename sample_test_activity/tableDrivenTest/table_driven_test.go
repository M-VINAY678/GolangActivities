package tabledriventest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsEven(n int) bool {
	return n%2 == 0
}
func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Even number", 4, true},
		{"Odd number", 7, false},
		{"Zero", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsEven(tt.input), "input is :%v", tt.input)
		})
	}
}
