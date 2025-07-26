package main

import (
	"fmt"
	"log"
	"os"
	"wcgo/cmd/io"
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

	fmt.Println(opType, filepath, text)
}
