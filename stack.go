package balance

type stack struct {
	top  *element
	size int
}

type element struct {
	value rune
	next  *element
}

func newStack() *stack {
	return &stack{}
}

func (stack *stack) len() int {
	return stack.size
}

func (stack *stack) push(value rune) {
	stack.top = &element{value, stack.top}
	stack.size++
}

func (stack *stack) pop() rune {
	if stack.isEmpty() {
		return rune(0)
	}

	value := stack.top.value
	stack.top = stack.top.next
	stack.size--

	return value
}

func (stack *stack) isEmpty() bool {
	return stack.len() == 0
}
