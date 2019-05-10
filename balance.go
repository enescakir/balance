package balance

import (
	"strings"
)

const (
	OPEN_PARENTHESES  = "[({"
	CLOSE_PARENTHESES = "])}"
)

func Check(text string) (valid bool, err error) {

	stack := newStack()

	for i, ch := range text {
		if pos := strings.IndexRune(OPEN_PARENTHESES, ch); pos != -1 {
			stack.push(rune(CLOSE_PARENTHESES[pos]))
		} else if pos := strings.IndexRune(CLOSE_PARENTHESES, ch); pos != -1 {
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
