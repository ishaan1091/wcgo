package utils

import "strings"

func GetBytesCountIfRequired(opType string, content string) (int, bool, error) {
	if !strings.Contains(opType, "c") {
		return 0, false, nil
	}

	return 0, false, nil
}
