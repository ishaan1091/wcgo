package utils

import (
	"strings"
	"wcgo/cmd/constants"
)

func GetBytesCountIfRequired(opType string, content string) (*int, error) {
	if !strings.Contains(opType, constants.BytesCountOp) {
		return nil, nil
	}

	count := len(content)

	return &count, nil
}
