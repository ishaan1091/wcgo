package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"wcgo/cmd/io"
)

func TestProcessWcIntegration(t *testing.T) {
	tests := []struct {
		name             string
		opType           string
		filepath         string
		fileContent      string
		expectedResult   *WcResult
		expectedOutput   string
		expectError      bool
		errorContains    string
		testOutput       bool // Whether to test output formatting
	}{
		{
			name:           "should process default operation with file",
			opType:         "-lwc",
			filepath:       "test_integration.txt",
			fileContent:    "hello world\ntest line",
			expectedResult: &WcResult{
				LinesCount: intPtr(2),
				WordsCount: intPtr(4),
				BytesCount: intPtr(21),
				Filepath:   "test_integration.txt",
			},
			expectedOutput: "\t2\t4\t21 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should process only lines operation",
			opType:         "-l",
			filepath:       "test_integration.txt",
			fileContent:    "hello world\ntest line\nthird line",
			expectedResult: &WcResult{
				LinesCount: intPtr(3),
				Filepath:   "test_integration.txt",
			},
			expectedOutput: "\t3 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should process only words operation",
			opType:         "-w",
			filepath:       "test_integration.txt",
			fileContent:    "hello world test",
			expectedResult: &WcResult{
				WordsCount: intPtr(3),
				Filepath:   "test_integration.txt",
			},
			expectedOutput: "\t3 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should process only bytes operation",
			opType:         "-c",
			filepath:       "test_integration.txt",
			fileContent:    "hello",
			expectedResult: &WcResult{
				BytesCount: intPtr(5),
				Filepath:   "test_integration.txt",
			},
			expectedOutput: "\t5 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should process only characters operation",
			opType:         "-m",
			filepath:       "test_integration.txt",
			fileContent:    "hello世界",
			expectedResult: &WcResult{
				CharactersCount: intPtr(7),
				Filepath:        "test_integration.txt",
			},
			expectedOutput: "\t7 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should process mixed operations",
			opType:         "-lw",
			filepath:       "test_integration.txt",
			fileContent:    "hello world\ntest line",
			expectedResult: &WcResult{
				LinesCount: intPtr(2),
				WordsCount: intPtr(4),
				Filepath:   "test_integration.txt",
			},
			expectedOutput: "\t2\t4 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should handle empty file",
			opType:         "-lwc",
			filepath:       "test_integration.txt",
			fileContent:    "",
			expectedResult: &WcResult{
				LinesCount: intPtr(0),
				WordsCount: intPtr(0),
				BytesCount: intPtr(0),
				Filepath:   "test_integration.txt",
			},
			expectedOutput: "\t0\t0\t0 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should handle whitespace only file",
			opType:         "-lwc",
			filepath:       "test_integration.txt",
			fileContent:    "   \t\n  ",
			expectedResult: &WcResult{
				LinesCount: intPtr(2),
				WordsCount: intPtr(0),
				BytesCount: intPtr(7),
				Filepath:   "test_integration.txt",
			},
			expectedOutput: "\t2\t0\t7 test_integration.txt\n",
			expectError:    false,
			testOutput:     true,
		},
		{
			name:           "should reject invalid operation type",
			opType:         "-x",
			filepath:       "test_integration.txt",
			fileContent:    "hello world",
			expectedResult: nil,
			expectedOutput: "",
			expectError:    true,
			errorContains:  "invalid flag x",
			testOutput:     false,
		},
		{
			name:           "should handle non-existent file",
			opType:         "-l",
			filepath:       "non_existent_file.txt",
			fileContent:    "",
			expectedResult: nil,
			expectedOutput: "",
			expectError:    true,
			errorContains:  "failed to read file",
			testOutput:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file if not testing non-existent file
			if tt.filepath != "non_existent_file.txt" {
				err := os.WriteFile(tt.filepath, []byte(tt.fileContent), 0644)
				if err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
				defer os.Remove(tt.filepath) // Clean up after test
			}

			// Test the actual integration function
			result, err := ProcessWc(tt.opType, tt.filepath)

			// Check error expectations
			if tt.expectError {
				if err == nil {
					t.Errorf("ProcessWc() expected error but got nil")
					return
				}
				if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("ProcessWc() error = %v, want error containing %v", err.Error(), tt.errorContains)
				}
			} else {
				if err != nil {
					t.Errorf("ProcessWc() unexpected error = %v", err)
					return
				}
			}

			// Check result expectations
			if !tt.expectError && tt.expectedResult != nil {
				assertWcResult(t, result, tt.expectedResult)
			}

			// Test output formatting if requested
			if tt.testOutput && !tt.expectError {
				// Capture output
				var buf bytes.Buffer
				logger := log.New(&buf, "", 0)

				// Format output
				io.OutputFormattedResult(logger, result.BytesCount, result.CharactersCount, result.LinesCount, result.WordsCount, result.Filepath)

				// Check output
				output := buf.String()
				if output != tt.expectedOutput {
					t.Errorf("Output = %q, want %q", output, tt.expectedOutput)
				}
			}
		})
	}
}

// Helper function to assert WcResult values
func assertWcResult(t *testing.T, actual, expected *WcResult) {
	if expected == nil && actual != nil {
		t.Errorf("WcResult = %+v, want nil", actual)
		return
	}
	if expected != nil && actual == nil {
		t.Errorf("WcResult = nil, want %+v", expected)
		return
	}
	if expected == nil && actual == nil {
		return
	}

	// Compare individual fields
	if expected.BytesCount != nil && actual.BytesCount != nil && *actual.BytesCount != *expected.BytesCount {
		t.Errorf("BytesCount = %v, want %v", *actual.BytesCount, *expected.BytesCount)
	}
	if expected.CharactersCount != nil && actual.CharactersCount != nil && *actual.CharactersCount != *expected.CharactersCount {
		t.Errorf("CharactersCount = %v, want %v", *actual.CharactersCount, *expected.CharactersCount)
	}
	if expected.LinesCount != nil && actual.LinesCount != nil && *actual.LinesCount != *expected.LinesCount {
		t.Errorf("LinesCount = %v, want %v", *actual.LinesCount, *expected.LinesCount)
	}
	if expected.WordsCount != nil && actual.WordsCount != nil && *actual.WordsCount != *expected.WordsCount {
		t.Errorf("WordsCount = %v, want %v", *actual.WordsCount, *expected.WordsCount)
	}
	if actual.Filepath != expected.Filepath {
		t.Errorf("Filepath = %v, want %v", actual.Filepath, expected.Filepath)
	}
}



// Helper function to create int pointers
func intPtr(i int) *int {
	return &i
}
