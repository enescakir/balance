package balance

import "testing"

func TestValidCheck(t *testing.T) {

	cases := []string{"", "[]", "[()]", "[[]]", "([()])", "{[()]}"}

	for _, str := range cases {
		valid, err := Check(str)

		if !valid {
			t.Errorf("Text: %q, Status: Invalid,  Expected: Valid,  Error: %v", str, err)
		}
	}
}

func TestMismatchError(t *testing.T) {
	cases := []string{"([)]", "([)", ")()(", "([]{)}"}

	for _, str := range cases {
		valid, err := Check(str)

		if valid {
			t.Errorf("Text: %q, Status: Valid,  Expected: Invalid,  Error: %v", str, err)
		}

		if err != nil {
			switch err.(type) {
			case *MismatchError:
			default:
				t.Errorf("Text: %q,  Error: %v, Expected: MismatchError", str, err)
			}
		}
	}
}

func TestUnclosedParenthesesError(t *testing.T) {
	cases := []string{"(((", "[[]", "(())(", "{{{()}}"}

	for _, str := range cases {
		valid, err := Check(str)

		if valid {
			t.Errorf("Text: %q, Status: Valid,  Expected: Invalid,  Error: %v", str, err)
		}

		if err != nil {
			switch err.(type) {
			case *UnclosedParenthesesError:
			default:
				t.Errorf("Text: %q,  Error: %v, Expected: UnclosedParenthesesError", str, err)
			}
		}
	}
}

func TestUnknownCharacterError(t *testing.T) {
	cases := []string{"((a))", "[]abc", "abf"}

	for _, str := range cases {
		valid, err := Check(str)

		if valid {
			t.Errorf("Text: %q, Status: Valid,  Expected: Invalid,  Error: %v", str, err)
		}

		if err != nil {
			switch err.(type) {
			case *UnknownCharacterError:
			default:
				t.Errorf("Text: %q,  Error: %v, Expected: UnknownCharacterError", str, err)
			}
		}
	}
}

func TestCheckCustom(t *testing.T) {

	cases := []struct{
		str, opens, closes string
		valid bool
	}{
		{"<>", "<", ">", true},
		{"{{}}\\/", "\\{", "/}", true},
		{")))()(((", ")", "(", true},
		{"<<>><<<>>", "<", ">)", false},
	}

	for _, c := range cases {
		valid, err := CheckCustom(c.str, c.opens, c.closes)

		if valid != c.valid {
			t.Errorf("Text: %q, Valid: %v,  Expected: %v,  Error: %v", c.str, valid, c.valid, err)
		}
	}
}
