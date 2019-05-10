package balance

import "fmt"

// MismatchError is returned when given string has a mismatched parentheses at index.
type MismatchError struct {
	Index int // index is position of mismatched parentheses
}

func (e *MismatchError) Error() string {
	return fmt.Sprintf("Mismatch at index: %d", e.Index)
}

// UnknownCharacterError is returned when given string has a unknown character rune at index.
type UnknownCharacterError struct {
	Index int  // index is position of unknown character
	Char  rune // char is unknown character
}

func (e *UnknownCharacterError) Error() string {
	return fmt.Sprintf("Unknown character %q at index: %d", e.Char, e.Index)
}

// UnclosedParenthesesError is returned when given string has a unclosed parentheses at index.
type UnclosedParenthesesError struct {
	Count int // count is number of unclosed parentheses
}

func (e *UnclosedParenthesesError) Error() string {
	return fmt.Sprintf("Unclosed %d parentheses", e.Count)
}

// CustomPairError is returned when given custom pair strings has different lengths.
type CustomPairError struct {
	Opens  string // opens is opening elements that has error
	Closes string // closes is closings elements that has error
}

func (e *CustomPairError) Error() string {
	return fmt.Sprintf("Custom pair strings should have same length. Opens: %q Closes: %q", e.Opens, e.Closes)
}
