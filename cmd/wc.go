package main

import (
	"log"
	"os"
	"wcgo/cmd/io"
	"wcgo/cmd/utils"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	opType, filepath := io.ReadArgs()

	if err := io.ValidateOpTypeArgs(opType); err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	text, err := io.GetContent(filepath)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	bytesCount, err := utils.GetBytesCountIfRequired(opType, text)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	charactersCount, err := utils.GetCharactersCountIfRequired(opType, text)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	linesCount, err := utils.GetLinesCountIfRequired(opType, text)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	wordsCount, err := utils.GetWordsCountIfRequired(opType, text)
	if err != nil {
		io.OutputFatalErrorAndExit(logger, err)
	}

	io.OutputFormattedResult(logger, bytesCount, charactersCount, linesCount, wordsCount, filepath)
}
