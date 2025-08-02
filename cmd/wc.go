package main

import (
	"log"
	"os"
	"wcgo/cmd/io"
	"wcgo/cmd/utils"
)

// WcResult holds the results of word count operations
type WcResult struct {
	BytesCount      *int
	CharactersCount *int
	LinesCount      *int
	WordsCount      *int
	Filepath        string
}

// ProcessWc processes the word count operations and returns the results
func ProcessWc(opType string, filepath string) (*WcResult, error) {
	if err := io.ValidateOpTypeArgs(opType); err != nil {
		return nil, err
	}

	text, err := io.GetContent(filepath)
	if err != nil {
		return nil, err
	}

	bytesCount, err := utils.GetBytesCountIfRequired(opType, text)
	if err != nil {
		return nil, err
	}

	charactersCount, err := utils.GetCharactersCountIfRequired(opType, text)
	if err != nil {
		return nil, err
	}

	linesCount, err := utils.GetLinesCountIfRequired(opType, text)
	if err != nil {
		return nil, err
	}

	wordsCount, err := utils.GetWordsCountIfRequired(opType, text)
	if err != nil {
		return nil, err
	}

	return &WcResult{
		BytesCount:      bytesCount,
		CharactersCount: charactersCount,
		LinesCount:      linesCount,
		WordsCount:      wordsCount,
		Filepath:        filepath,
	}, nil
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	opType, filepath := io.ReadArgs()

	result, err := ProcessWc(opType, filepath)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	io.OutputFormattedResult(logger, result.BytesCount, result.CharactersCount, result.LinesCount, result.WordsCount, result.Filepath)
}
