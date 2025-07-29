package utils

import (
	"strings"
	"wcgo/cmd/constants"
)

func GetLinesCountIfRequired(opType string, content string) (*int, error) {
	if !strings.Contains(opType, constants.LinesCountOp) {
		return nil, nil
	}

	lines := strings.Split(content, "\n")

	count := len(lines)

	if lines[count-1] == "" {
		count--
	}

	return &count, nil
}
