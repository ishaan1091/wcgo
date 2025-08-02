package utils

import (
	"testing"
)

func TestGetBytesCountIfRequired(t *testing.T) {
	tests := []struct {
		name     string
		opType   string
		content  string
		expected *int
	}{
		{
			name:     "should return nil when bytes operation not requested",
			opType:   "-lw",
			content:  "hello world",
			expected: nil,
		},
		{
			name:     "should count bytes in simple text",
			opType:   "-c",
			content:  "hello",
			expected: intPtr(5),
		},
		{
			name:     "should count bytes with spaces",
			opType:   "-c",
			content:  "hello world",
			expected: intPtr(11),
		},
		{
			name:     "should count bytes in empty string",
			opType:   "-c",
			content:  "",
			expected: intPtr(0),
		},
		{
			name:     "should count bytes with newlines",
			opType:   "-c",
			content:  "hello\nworld",
			expected: intPtr(11),
		},
		{
			name:     "should count bytes with tabs",
			opType:   "-c",
			content:  "hello\tworld",
			expected: intPtr(11),
		},
		{
			name:     "should count bytes with special characters",
			opType:   "-c",
			content:  "hello@world!test#case",
			expected: intPtr(21),
		},
		{
			name:     "should count bytes with numbers",
			opType:   "-c",
			content:  "hello123world456",
			expected: intPtr(16),
		},
		{
			name:     "should count bytes with unicode (multi-byte characters)",
			opType:   "-c",
			content:  "hello世界world",
			expected: intPtr(16), // 5 + 6 (世界 = 6 bytes) + 5
		},
		{
			name:     "should count bytes with mixed operations",
			opType:   "-lwc",
			content:  "hello world",
			expected: intPtr(11),
		},
		{
			name:     "should count bytes with only whitespace",
			opType:   "-c",
			content:  "   \t\n  ",
			expected: intPtr(7),
		},
		{
			name:     "should count bytes with emoji",
			opType:   "-c",
			content:  "hello😀world",
			expected: intPtr(14), // 5 + 4 (😀 = 4 bytes) + 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetBytesCountIfRequired(tt.opType, tt.content)

			if err != nil {
				t.Errorf("GetBytesCountIfRequired() error = %v", err)
				return
			}

			if tt.expected == nil && result != nil {
				t.Errorf("GetBytesCountIfRequired() = %v, want nil", *result)
			} else if tt.expected != nil && result == nil {
				t.Errorf("GetBytesCountIfRequired() = nil, want %v", *tt.expected)
			} else if tt.expected != nil && result != nil && *result != *tt.expected {
				t.Errorf("GetBytesCountIfRequired() = %v, want %v", *result, *tt.expected)
			}
		})
	}
}
