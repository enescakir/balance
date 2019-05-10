package balance

import "fmt"

type MismatchError struct {
	index int
}

func (e *MismatchError) Error() string {
	return fmt.Sprintf("Mismatch at index: %d", e.index)
}

type UnknownCharacterError struct {
	index int
	char  rune
}

func (e *UnknownCharacterError) Error() string {
	return fmt.Sprintf("Unknown character %q at index: %d", e.char, e.index)
}

type UnclosedParenthesesError struct {
	count int
}

func (e *UnclosedParenthesesError) Error() string {
	return fmt.Sprintf("Unclosed %d parentheses", e.count)
}
