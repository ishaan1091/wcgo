package io

import (
	"bufio"
	"fmt"
	"os"
	"wcgo/cmd/constants"
)

func ReadArgs() (string, string) {
	// Read args from command line
	args := os.Args[1:]

	// Read opType and filepath if provided
	var opType, filepath string

	if len(args) > 0 {
		if args[0] != "" && string(args[0][0]) == "-" {
			opType = args[0]

			if len(args) > 1 {
				filepath = args[1]
			}
		} else if args[0] != "" {
			filepath = args[0]
		}
	}

	if opType == "" || opType == "-" {
		opType = "-lwc"
	}

	return opType, filepath
}

func ValidateOpTypeArgs(opType string) error {
	// opType should only contain the possible opTypes
	if opType == "" {
		return nil
	}

	if string(opType[0]) == "-" {
		opType = opType[1:]
	}

	for _, op := range opType {
		if !constants.IsValidOpType(op) {
			return fmt.Errorf("invalid flag %v", string(op))
		}
	}

	return nil
}

func GetContent(filepath string) (string, error) {
	// Read from filepath if given
	if filepath != "" {
		content, err := os.ReadFile(filepath)
		if err != nil {
			return "", fmt.Errorf("failed to read file : %v, %v", filepath, err)
		}

		return string(content), err
	}

	// Else read from standard input
	var content string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	return content, nil
}
