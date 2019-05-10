// Package balance provides functions for validating parentheses balance of strings
package balance

import "strings"

const (
	OpenParentheses  = "[({" // OPEN_PARENTHESES keeps default opening elements for checking.
	CloseParentheses = "])}" // CLOSE_PARENTHESES keeps default closing elements for checking.
)

// Check validates given string with default parenthesises via CheckCustom.
// Default parenthesises: (), [], {}
func Check(str string) (valid bool, err error) {
	return CheckCustom(str, OpenParentheses, CloseParentheses)
}

// CheckCustom validates given string for balance of opening and closing characters.
// It uses stack data structure for validation.
// Basically, it iterates given string, push opening elements to stack and expect to pop closing elements from stack.
// It returns special errors when string is not valid.
// Possible errors: MismatchError, UnclosedParenthesesError, UnknownCharacterError, CustomPairError.
func CheckCustom(str string, opens string, closes string) (valid bool, err error) {

	if len(opens) != len(closes) {
		return false, &CustomPairError{opens, closes}
	}

	stack := newStack()

	for i, ch := range str {
		if pos := strings.IndexRune(opens, ch); pos != -1 {
			stack.push(rune(closes[pos]))
		} else if pos := strings.IndexRune(closes, ch); pos != -1 {
			if ch2 := stack.pop(); ch != ch2 {
				return false, &MismatchError{i}
			}
		} else {
			return false, &UnknownCharacterError{i, ch}
		}
	}

	if !stack.isEmpty() {
		return false, &UnclosedParenthesesError{stack.size}
	}

	return true, nil
}
