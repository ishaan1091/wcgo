package constants

import (
	"testing"
)

func TestIsValidOpType(t *testing.T) {
	tests := []struct {
		name     string
		op       rune
		expected bool
	}{
		{
			name:     "should return true for lines operation",
			op:       'l',
			expected: true,
		},
		{
			name:     "should return true for words operation",
			op:       'w',
			expected: true,
		},
		{
			name:     "should return true for bytes operation",
			op:       'c',
			expected: true,
		},
		{
			name:     "should return true for characters operation",
			op:       'm',
			expected: true,
		},
		{
			name:     "should return false for invalid operation",
			op:       'x',
			expected: false,
		},
		{
			name:     "should return false for uppercase operation",
			op:       'L',
			expected: false,
		},
		{
			name:     "should return false for number",
			op:       '1',
			expected: false,
		},
		{
			name:     "should return false for special character",
			op:       '@',
			expected: false,
		},
		{
			name:     "should return false for space",
			op:       ' ',
			expected: false,
		},
		{
			name:     "should return false for newline",
			op:       '\n',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidOpType(tt.op)
			if result != tt.expected {
				t.Errorf("IsValidOpType(%q) = %v, want %v", tt.op, result, tt.expected)
			}
		})
	}
}
