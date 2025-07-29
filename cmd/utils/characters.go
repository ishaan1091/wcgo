package utils

import (
	"strings"
	"wcgo/cmd/constants"
)

func GetCharactersCountIfRequired(opType string, content string) (*int, error) {
	if !strings.Contains(opType, constants.CharactersCountOp) {
		return nil, nil
	}

	count := 0

	for range content {
		count++
	}

	return &count, nil
}
