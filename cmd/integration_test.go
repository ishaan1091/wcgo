package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestProcessWcIntegration(t *testing.T) {
	tests := []struct {
		name           string
		opType         string
		filepath       string
		fileContent    string
		expectedOutput string
		expectError    bool
		errorContains  string
	}{
		{
			name:           "should process default operation with file",
			opType:         "-lwc",
			filepath:       "test_integration.txt",
			fileContent:    "hello world\ntest line",
			expectedOutput: "\t2\t4\t21 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should process only lines operation",
			opType:         "-l",
			filepath:       "test_integration.txt",
			fileContent:    "hello world\ntest line\nthird line",
			expectedOutput: "\t3 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should process only words operation",
			opType:         "-w",
			filepath:       "test_integration.txt",
			fileContent:    "hello world test",
			expectedOutput: "\t3 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should process only bytes operation",
			opType:         "-c",
			filepath:       "test_integration.txt",
			fileContent:    "hello",
			expectedOutput: "\t5 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should process only characters operation",
			opType:         "-m",
			filepath:       "test_integration.txt",
			fileContent:    "hello世界",
			expectedOutput: "\t7 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should process mixed operations",
			opType:         "-lw",
			filepath:       "test_integration.txt",
			fileContent:    "hello world\ntest line",
			expectedOutput: "\t2\t4 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should handle empty file",
			opType:         "-lwc",
			filepath:       "test_integration.txt",
			fileContent:    "",
			expectedOutput: "\t0\t0\t0 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should handle whitespace only file",
			opType:         "-lwc",
			filepath:       "test_integration.txt",
			fileContent:    "   \t\n  ",
			expectedOutput: "\t2\t0\t7 test_integration.txt\n",
			expectError:    false,
		},
		{
			name:           "should reject invalid operation type",
			opType:         "-x",
			filepath:       "test_integration.txt",
			fileContent:    "hello world",
			expectedOutput: "",
			expectError:    true,
			errorContains:  "invalid flag x",
		},
		{
			name:           "should handle non-existent file",
			opType:         "-l",
			filepath:       "non_existent_file.txt",
			fileContent:    "",
			expectedOutput: "",
			expectError:    true,
			errorContains:  "failed to read file",
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

			// Capture output
			var buf bytes.Buffer
			logger := log.New(&buf, "", 0)

			// Test the actual integration function
			err := ProcessWc(tt.opType, tt.filepath, logger)

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

			// Check output expectations
			if !tt.expectError {
				output := buf.String()
				if output != tt.expectedOutput {
					t.Errorf("ProcessWc() output = %q, want %q", output, tt.expectedOutput)
				}
			}
		})
	}
}
