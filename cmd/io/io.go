package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"wcgo/cmd/constants"
)

func ReadArgs() (string, string) {
	// Read args from command line
	args := os.Args[1:]

	// Read opType and filepath if provided
	var opType, filepath string

	if len(args) > 0 {
		opType = args[0]

		if len(args) > 1 {
			filepath = args[1]
		}
	}

	return opType, filepath
}

func ValidateOpTypeArgs(opType string) error {
	// opType should only contain the possible opTypes
	opTypes := opType[1:]
	for _, op := range opTypes {
		if !constants.IsValidOpType(op) {
			return fmt.Errorf("invalid flag %v", string(op))
		}
	}

	return nil
}

func OutputFatalErrorAndExit(logger *log.Logger, err error) {
	logger.Fatal(err)
	os.Exit(1)
}

func GetContent(filepath string) (string, error) {
	if filepath != "" {
		content, err := os.ReadFile(filepath)
		if err != nil {
			return "", fmt.Errorf("failed to read file : %v, %v", filepath, err)
		}

		return string(content), err
	}

	var content string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		content += scanner.Text()
	}

	return content, nil
}
