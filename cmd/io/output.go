package io

import (
	"fmt"
	"log"
	"os"
)

func OutputFormattedResult(logger *log.Logger, bytesCount *int, charactersCount *int, linesCount *int, wordsCount *int, filepath string) {
	var output string

	if bytesCount != nil {
		output += fmt.Sprintf("\t%v", *bytesCount)
	}

	if charactersCount != nil {
		output += fmt.Sprintf("\t%v", *charactersCount)
	}

	if linesCount != nil {
		output += fmt.Sprintf("\t%v", *linesCount)
	}

	if wordsCount != nil {
		output += fmt.Sprintf("\t%v", *wordsCount)
	}

	if filepath != "" {
		output += fmt.Sprintf(" %v", filepath)
	}

	logger.Println(output)
}

func OutputFatalErrorAndExit(logger *log.Logger, err error) {
	logger.Fatal(err)
	os.Exit(1)
}
