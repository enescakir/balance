package balance

import (
	"strings"
	"testing"
)

func TestMismatchError_Error(t *testing.T) {
	e := MismatchError{3}
	m := e.Error()

	if strings.Index(m, "Mismatch") == -1 {
		t.Errorf("MismatchedError test fails")
	}
}

func TestUnknownCharacterError_Error(t *testing.T) {
	e := UnknownCharacterError{3, rune(30)}
	m := e.Error()

	if strings.Index(m, "Unknown character") == -1 {
		t.Errorf("UnknownCharacterError test fails")
	}
}

func TestUnclosedParenthesesError_Error(t *testing.T) {
	e := UnclosedParenthesesError{3}
	m := e.Error()

	if strings.Index(m, "Unclosed") == -1 {
		t.Errorf("UnclosedParenthesesError test fails")
	}
}

func TestCustomPairError_Error(t *testing.T) {
	e := CustomPairError{"((", ")"}
	m := e.Error()

	if strings.Index(m, "Custom pair strings") == -1 {
		t.Errorf("CustomPairError test fails")
	}
}
