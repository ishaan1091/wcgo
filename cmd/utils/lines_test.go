package utils

import (
	"testing"
)

func TestGetLinesCountIfRequired(t *testing.T) {
	tests := []struct {
		name     string
		opType   string
		content  string
		expected *int
	}{
		{
			name:     "should return nil when lines operation not requested",
			opType:   "-wc",
			content:  "hello\nworld",
			expected: nil,
		},
		{
			name:     "should count lines in simple text",
			opType:   "-l",
			content:  "hello\nworld",
			expected: intPtr(2),
		},
		{
			name:     "should count single line",
			opType:   "-l",
			content:  "hello world",
			expected: intPtr(1),
		},
		{
			name:     "should count lines with empty string",
			opType:   "-l",
			content:  "",
			expected: intPtr(0),
		},
		{
			name:     "should count lines ending with newline",
			opType:   "-l",
			content:  "hello\nworld\n",
			expected: intPtr(2),
		},
		{
			name:     "should count lines with multiple consecutive newlines",
			opType:   "-l",
			content:  "hello\n\nworld",
			expected: intPtr(3),
		},
		{
			name:     "should count lines with only newlines",
			opType:   "-l",
			content:  "\n\n\n",
			expected: intPtr(3),
		},
		{
			name:     "should count lines with mixed operations",
			opType:   "-lwc",
			content:  "hello\nworld\ntest",
			expected: intPtr(3),
		},
		{
			name:     "should count lines with carriage return and newline",
			opType:   "-l",
			content:  "hello\r\nworld",
			expected: intPtr(2),
		},
		{
			name:     "should count lines with tabs and spaces",
			opType:   "-l",
			content:  "hello\tworld\n  test  line",
			expected: intPtr(2),
		},
		{
			name:     "should count lines with special characters",
			opType:   "-l",
			content:  "hello@world\n!test#case",
			expected: intPtr(2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetLinesCountIfRequired(tt.opType, tt.content)

			if err != nil {
				t.Errorf("GetLinesCountIfRequired() error = %v", err)
				return
			}

			if tt.expected == nil && result != nil {
				t.Errorf("GetLinesCountIfRequired() = %v, want nil", *result)
			} else if tt.expected != nil && result == nil {
				t.Errorf("GetLinesCountIfRequired() = nil, want %v", *tt.expected)
			} else if tt.expected != nil && result != nil && *result != *tt.expected {
				t.Errorf("GetLinesCountIfRequired() = %v, want %v", *result, *tt.expected)
			}
		})
	}
}
