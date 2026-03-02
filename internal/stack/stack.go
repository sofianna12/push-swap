// Package stack provides the Stack data structure used by the push-swap
// project along with fundamental stack manipulation primitives.
//
// The Stack represents a LIFO (Last-In, First-Out) collection of integers
// where index 0 corresponds to the top of the stack.
//
// All methods that modify the stack operate on pointer receivers.
package stack

// Stack represents a LIFO collection of integers.
// The element at index 0 is considered the top of the stack.
type Stack struct {
	data []int
}

// NewStack creates a new Stack initialized with the given values.
// The input slice is defensively copied to preserve immutability.
//
// Parameters:
//   - values: the initial stack values, where index 0 is the top.
//
// Returns a pointer to the newly created Stack.
func NewStack(values []int) *Stack {
	cp := make([]int, len(values))
	copy(cp, values)
	return &Stack{data: cp}
}

// Push inserts a value at the top of the stack.
func (s *Stack) Push(v int) {
	s.data = append([]int{v}, s.data...)
}

// Pop removes and returns the top element of the stack.
// If the stack is empty, it returns 0 and false.
//
// Returns:
//   - the popped integer value
//   - a boolean indicating whether a value was successfully removed
func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	val := s.data[0]
	s.data = s.data[1:]
	return val, true
}

// Peek returns the top element of the stack without removing it.
// If the stack is empty, it returns 0 and false.
//
// Returns:
//   - the top integer value
//   - a boolean indicating whether a value exists
func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[0], true
}

// Len returns the number of elements currently in the stack.
func (s *Stack) Len() int {
	return len(s.data)
}

// IsEmpty reports whether the stack contains no elements.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Values returns a copy of the stack's underlying data slice.
// The returned slice is a defensive copy to prevent external mutation.
func (s *Stack) Values() []int {
	cp := make([]int, len(s.data))
	copy(cp, s.data)
	return cp
}
