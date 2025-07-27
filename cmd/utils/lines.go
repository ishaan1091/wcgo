package utils

import (
	"strings"
	"wcgo/cmd/constants"
)

func GetLinesCountIfRequired(opType string, content string) (*int, error) {
	if !strings.Contains(opType, constants.LinesCountOp) {
		return nil, nil
	}

	count := 0

	return &count, nil
}
