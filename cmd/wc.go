package main

import (
	"log"
	"os"
	"wcgo/cmd/io"
	"wcgo/cmd/utils"
)

func ProcessWc(opType string, filepath string, logger *log.Logger) error {
	if err := io.ValidateOpTypeArgs(opType); err != nil {
		return err
	}

	text, err := io.GetContent(filepath)
	if err != nil {
		return err
	}

	bytesCount, err := utils.GetBytesCountIfRequired(opType, text)
	if err != nil {
		return err
	}

	charactersCount, err := utils.GetCharactersCountIfRequired(opType, text)
	if err != nil {
		return err
	}

	linesCount, err := utils.GetLinesCountIfRequired(opType, text)
	if err != nil {
		return err
	}

	wordsCount, err := utils.GetWordsCountIfRequired(opType, text)
	if err != nil {
		return err
	}

	io.OutputFormattedResult(logger, bytesCount, charactersCount, linesCount, wordsCount, filepath)
	return nil
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	opType, filepath := io.ReadArgs()

	if err := ProcessWc(opType, filepath, logger); err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}
}
