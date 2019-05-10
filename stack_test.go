package balance

import "testing"

func TestStack(t *testing.T) {
	cases := []struct {
		in, want string
		len      int
	}{
		{"123456", "654321", 6},
		{"abcd", "dcba", 4},
		{"", "", 0},
	}

	for _, c := range cases {
		stack := newStack()

		if !stack.isEmpty() {
			t.Errorf("Stack have to be empty at the beginning of test")
		}

		for _, ch := range c.in {
			stack.push(ch)
		}

		if stack.len() != c.len {
			t.Errorf("Stack lenght is %q, expected %q", stack.len(), c.len)
		}

		for i, ch := range c.want {
			if ch2 := stack.pop(); ch2 != ch {
				t.Errorf("Element[%q] == %q, expected %q", i, ch2, ch)
			}
		}

		if !stack.isEmpty() {
			t.Errorf("Stack have to be empty at the end of test")
		}
	}
}
