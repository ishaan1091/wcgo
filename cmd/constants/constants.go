package constants

import "slices"

const (
	LinesCountOp      = "l"
	WordsCountOp      = "w"
	BytesCountOp      = "c"
	CharactersCountOp = "m"
)

func IsValidOpType(op rune) bool {
	validOpTypes := []string{LinesCountOp, WordsCountOp, BytesCountOp, CharactersCountOp}

	return slices.Contains(validOpTypes, string(op))
}
