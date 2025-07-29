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

	// TODO: Fix for multilingual input like 你好世界 evaluates to Hello world in chinese which is two words.
	// wc identifies and counts this correctly as two words but we end up counting this as one word only

	for _, line := range strings.Split(content, "\n") {
		words := strings.Split(line, " ")
		if len(words) > 1 || words[0] != "" {
			count += len(words)
		}
	}

	return &count, nil
}
