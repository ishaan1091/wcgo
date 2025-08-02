package utils

import (
	"testing"
)

func TestGetCharactersCountIfRequired(t *testing.T) {
	tests := []struct {
		name     string
		opType   string
		content  string
		expected *int
	}{
		{
			name:     "should return nil when characters operation not requested",
			opType:   "-lw",
			content:  "hello world",
			expected: nil,
		},
		{
			name:     "should count characters in simple text",
			opType:   "-m",
			content:  "hello",
			expected: intPtr(5),
		},
		{
			name:     "should count characters with spaces",
			opType:   "-m",
			content:  "hello world",
			expected: intPtr(11),
		},
		{
			name:     "should count characters in empty string",
			opType:   "-m",
			content:  "",
			expected: intPtr(0),
		},
		{
			name:     "should count characters with newlines",
			opType:   "-m",
			content:  "hello\nworld",
			expected: intPtr(11),
		},
		{
			name:     "should count characters with tabs",
			opType:   "-m",
			content:  "hello\tworld",
			expected: intPtr(11),
		},
		{
			name:     "should count characters with special characters",
			opType:   "-m",
			content:  "hello@world!test#case",
			expected: intPtr(21),
		},
		{
			name:     "should count characters with numbers",
			opType:   "-m",
			content:  "hello123world456",
			expected: intPtr(16),
		},
		{
			name:     "should count characters with unicode",
			opType:   "-m",
			content:  "hello世界world",
			expected: intPtr(12),
		},
		{
			name:     "should count characters with mixed operations",
			opType:   "-lmc",
			content:  "hello world",
			expected: intPtr(11),
		},
		{
			name:     "should count characters with only whitespace",
			opType:   "-m",
			content:  "   \t\n  ",
			expected: intPtr(7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetCharactersCountIfRequired(tt.opType, tt.content)

			if err != nil {
				t.Errorf("GetCharactersCountIfRequired() error = %v", err)
				return
			}

			if tt.expected == nil && result != nil {
				t.Errorf("GetCharactersCountIfRequired() = %v, want nil", *result)
			} else if tt.expected != nil && result == nil {
				t.Errorf("GetCharactersCountIfRequired() = nil, want %v", *tt.expected)
			} else if tt.expected != nil && result != nil && *result != *tt.expected {
				t.Errorf("GetCharactersCountIfRequired() = %v, want %v", *result, *tt.expected)
			}
		})
	}
}
