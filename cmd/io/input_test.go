package io

import (
	"testing"
)

func TestValidateOpTypeArgs(t *testing.T) {
	tests := []struct {
		name        string
		opType      string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "should accept empty opType",
			opType:      "",
			expectError: false,
		},
		{
			name:        "should accept empty opType",
			opType:      "-",
			expectError: false,
		},
		{
			name:        "should accept valid single operation",
			opType:      "-l",
			expectError: false,
		},
		{
			name:        "should accept valid multiple operations",
			opType:      "-lwc",
			expectError: false,
		},
		{
			name:        "should accept valid operations without dash",
			opType:      "lwc",
			expectError: false,
		},
		{
			name:        "should accept all valid operations",
			opType:      "-lmwc",
			expectError: false,
		},
		{
			name:        "should reject invalid operation",
			opType:      "-x",
			expectError: true,
			errorMsg:    "invalid flag x",
		},
		{
			name:        "should reject mixed valid and invalid operations",
			opType:      "-lwx",
			expectError: true,
			errorMsg:    "invalid flag x",
		},
		{
			name:        "should reject uppercase operation",
			opType:      "-L",
			expectError: true,
			errorMsg:    "invalid flag L",
		},
		{
			name:        "should reject number",
			opType:      "-1",
			expectError: true,
			errorMsg:    "invalid flag 1",
		},
		{
			name:        "should reject special character",
			opType:      "-@",
			expectError: true,
			errorMsg:    "invalid flag @",
		},
		{
			name:        "should reject space",
			opType:      "- ",
			expectError: true,
			errorMsg:    "invalid flag  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOpTypeArgs(tt.opType)

			if tt.expectError {
				if err == nil {
					t.Errorf("ValidateOpTypeArgs() expected error but got nil")
					return
				}
				if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("ValidateOpTypeArgs() error = %v, want %v", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateOpTypeArgs() unexpected error = %v", err)
				}
			}
		})
	}
}
