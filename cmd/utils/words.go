package utils

import (
	"strings"
	"unicode"
	"wcgo/cmd/constants"
)

func GetWordsCountIfRequired(opType string, content string) (*int, error) {
	if !strings.Contains(opType, constants.WordsCountOp) {
		return nil, nil
	}

	count := 0
	readingWord := false

	for _, r := range content {
		if unicode.IsSpace(r) {
			if readingWord {
				count++
				readingWord = false
			}
		} else {
			readingWord = true
		}
	}
	if readingWord {
		count++
	}

	return &count, nil
}
