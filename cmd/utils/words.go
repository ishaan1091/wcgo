package utils

import (
	"strings"
	"wcgo/cmd/constants"
)

func GetWordsCountIfRequired(opType string, content string) (*int, error) {
	if !strings.Contains(opType, constants.WordsCountOp) {
		return nil, nil
	}

	count := 0

	return &count, nil
}
