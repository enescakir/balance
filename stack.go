package balance

// stack represents FIFO data structure.
// It has reference to its top element and its number of elements.
type stack struct {
	top  *element
	size int
}

// element represents an item of the stack.
// It has rune value and reference to next element of the stack.
type element struct {
	value rune
	next  *element
}

// newStack returns newly created stack reference.
func newStack() *stack {
	return &stack{}
}

// len returns the number of runes in the stack.
func (s *stack) len() int {
	return s.size
}

// push adds given rune to end of the stack.
func (s *stack) push(value rune) {
	s.top = &element{value, s.top}
	s.size++
}

// pop remove the last element of stack and returns it.
func (s *stack) pop() rune {
	if s.isEmpty() {
		return rune(0)
	}

	value := s.top.value
	s.top = s.top.next
	s.size--

	return value
}

// isEmpty checks the length of stack is zero or not.
func (s *stack) isEmpty() bool {
	return s.len() == 0
}
