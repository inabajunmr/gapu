package ufap

// Stack is Stack
type Stack struct {
	stack []interface{}
}

// Push new value
func (stack *Stack) Push(val interface{}) {
	stack.stack = append(stack.stack, val)
}

// Pop new value
func (stack *Stack) Pop() interface{} {
	val := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[0 : len(stack.stack)-1]
	return val
}

// IsEmpty return true when stack is empty
func (stack *Stack) IsEmpty() bool {
	return len(stack.stack) == 0
}
