package stack

// ArrayStack is an implementation of a stack.
type ArrayStack struct {
	elements []interface{}
}

// NewStack creates a new array stack.
func NewStack() *ArrayStack {
	return &ArrayStack{}
}

// Size returns the number of elements in the stack.
func (s *ArrayStack) Size() int {
	return len(s.elements)
}

// IsEmpty returns true or false whether the stack has zero elements or not.
func (s *ArrayStack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Push adds an element to the stack.
func (s *ArrayStack) Push(e interface{}) {
	s.elements = append(s.elements, e)
}

// Pop fetches the top element of the stack and removes it.
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	result := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return result
}

// Peek returns the top of element from the stack, but does not remove it.
func (s *ArrayStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.elements[len(s.elements)-1]
}
