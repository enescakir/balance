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
func (stack *stack) len() int {
	return stack.size
}

// push adds given rune to end of the stack.
func (stack *stack) push(value rune) {
	stack.top = &element{value, stack.top}
	stack.size++
}

// pop remove the last element of stack and returns it.
func (stack *stack) pop() rune {
	if stack.isEmpty() {
		return rune(0)
	}

	value := stack.top.value
	stack.top = stack.top.next
	stack.size--

	return value
}

// isEmpty checks the length of stack is zero or not.
func (stack *stack) isEmpty() bool {
	return stack.len() == 0
}
