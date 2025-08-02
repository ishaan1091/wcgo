package utils

import (
	"testing"
)

func TestGetWordsCountIfRequired(t *testing.T) {
	tests := []struct {
		name     string
		opType   string
		content  string
		expected *int
	}{
		{
			name:     "should return nil when words operation not requested",
			opType:   "-lc",
			content:  "hello world",
			expected: nil,
		},
		{
			name:     "should count words in simple text",
			opType:   "-w",
			content:  "hello world",
			expected: intPtr(2),
		},
		{
			name:     "should count words with multiple spaces",
			opType:   "-w",
			content:  "hello   world   test",
			expected: intPtr(3),
		},
		{
			name:     "should count words with tabs and newlines",
			opType:   "-w",
			content:  "hello\tworld\ntest",
			expected: intPtr(3),
		},
		{
			name:     "should count single word",
			opType:   "-w",
			content:  "hello",
			expected: intPtr(1),
		},
		{
			name:     "should count words with leading/trailing spaces",
			opType:   "-w",
			content:  "  hello world  ",
			expected: intPtr(2),
		},
		{
			name:     "should count words in empty string",
			opType:   "-w",
			content:  "",
			expected: intPtr(0),
		},
		{
			name:     "should count words with only spaces",
			opType:   "-w",
			content:  "   \t\n  ",
			expected: intPtr(0),
		},
		{
			name:     "should count words with mixed operations",
			opType:   "-lwc",
			content:  "hello world\ntest line",
			expected: intPtr(4),
		},
		{
			name:     "should count words with special characters",
			opType:   "-w",
			content:  "hello@world!test#case",
			expected: intPtr(1),
		},
		{
			name:     "should count words with numbers",
			opType:   "-w",
			content:  "hello 123 world 456",
			expected: intPtr(4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetWordsCountIfRequired(tt.opType, tt.content)

			if err != nil {
				t.Errorf("GetWordsCountIfRequired() error = %v", err)
				return
			}

			if tt.expected == nil && result != nil {
				t.Errorf("GetWordsCountIfRequired() = %v, want nil", *result)
			} else if tt.expected != nil && result == nil {
				t.Errorf("GetWordsCountIfRequired() = nil, want %v", *tt.expected)
			} else if tt.expected != nil && result != nil && *result != *tt.expected {
				t.Errorf("GetWordsCountIfRequired() = %v, want %v", *result, *tt.expected)
			}
		})
	}
}
