package balance

import (
	"strings"
)

const (
	OPEN_PARENTHESES  = "[({"
	CLOSE_PARENTHESES = "])}"
)

func Check(str string) (valid bool, err error) {
	return CheckCustom(str, OPEN_PARENTHESES, CLOSE_PARENTHESES)
}

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
