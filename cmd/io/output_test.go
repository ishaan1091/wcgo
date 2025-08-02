package io

import (
	"bytes"
	"log"
	"testing"
)

func TestOutputFormattedResult(t *testing.T) {
	tests := []struct {
		name            string
		bytesCount      *int
		charactersCount *int
		linesCount      *int
		wordsCount      *int
		filepath        string
		expectedOutput  string
	}{
		{
			name:            "should format output with all counts",
			bytesCount:      intPtr(11),
			charactersCount: intPtr(11),
			linesCount:      intPtr(1),
			wordsCount:      intPtr(2),
			filepath:        "test.txt",
			expectedOutput:  "\t1\t2\t11\t11 test.txt\n",
		},
		{
			name:            "should format output with only lines count",
			bytesCount:      nil,
			charactersCount: nil,
			linesCount:      intPtr(5),
			wordsCount:      nil,
			filepath:        "test.txt",
			expectedOutput:  "\t5 test.txt\n",
		},
		{
			name:            "should format output with only words count",
			bytesCount:      nil,
			charactersCount: nil,
			linesCount:      nil,
			wordsCount:      intPtr(10),
			filepath:        "test.txt",
			expectedOutput:  "\t10 test.txt\n",
		},
		{
			name:            "should format output with only bytes count",
			bytesCount:      intPtr(25),
			charactersCount: nil,
			linesCount:      nil,
			wordsCount:      nil,
			filepath:        "test.txt",
			expectedOutput:  "\t25 test.txt\n",
		},
		{
			name:            "should format output with only characters count",
			bytesCount:      nil,
			charactersCount: intPtr(30),
			linesCount:      nil,
			wordsCount:      nil,
			filepath:        "test.txt",
			expectedOutput:  "\t30 test.txt\n",
		},
		{
			name:            "should format output without filepath",
			bytesCount:      intPtr(11),
			charactersCount: intPtr(11),
			linesCount:      intPtr(1),
			wordsCount:      intPtr(2),
			filepath:        "",
			expectedOutput:  "\t1\t2\t11\t11\n",
		},
		{
			name:            "should format output with mixed counts",
			bytesCount:      intPtr(20),
			charactersCount: nil,
			linesCount:      intPtr(3),
			wordsCount:      intPtr(5),
			filepath:        "mixed.txt",
			expectedOutput:  "\t3\t5\t20 mixed.txt\n",
		},
		{
			name:            "should format output with zero counts",
			bytesCount:      intPtr(0),
			charactersCount: intPtr(0),
			linesCount:      intPtr(0),
			wordsCount:      intPtr(0),
			filepath:        "empty.txt",
			expectedOutput:  "\t0\t0\t0\t0 empty.txt\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture output
			var buf bytes.Buffer
			logger := log.New(&buf, "", 0)

			OutputFormattedResult(logger, tt.bytesCount, tt.charactersCount, tt.linesCount, tt.wordsCount, tt.filepath)

			output := buf.String()
			if output != tt.expectedOutput {
				t.Errorf("OutputFormattedResult() output = %q, want %q", output, tt.expectedOutput)
			}
		})
	}
}

// Helper function to create int pointers
func intPtr(i int) *int {
	return &i
}
