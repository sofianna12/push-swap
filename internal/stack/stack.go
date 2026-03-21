// Package stack provides the Stack data structure used by push-swap and checker.
//
// The Stack represents a LIFO (Last-In, First-Out) collection of integers
// where index 0 corresponds to the top of the stack.
//
// All methods that modify the stack operate on pointer receivers.
package stack

// Stack represents a LIFO collection of integers.
// The element at index 0 is considered the top of the stack.
// name is "a" or "b" and is used for debugging only — it never appears in output.
type Stack struct {
	data []int
	name string
}

// New creates a new Stack initialised with the given integers.
// The input slice is defensively copied so the caller cannot mutate the stack.
//
// Parameters:
//   - name: "a" or "b", used for debugging only, never printed to output.
//   - nums: initial values where index 0 is the top of the stack.
//
// Returns a pointer to the newly created Stack.
func New(name string, nums []int) *Stack {
	cp := make([]int, len(nums))
	copy(cp, nums)
	return &Stack{name: name, data: cp}
}

// Push inserts a value at the top of the stack.
//
// Parameters:
//   - v: the integer to push onto the top.
func (s *Stack) Push(v int) {
	s.data = append([]int{v}, s.data...)
}

// Pop removes and returns the top element of the stack.
// Returns (0, false) if the stack is empty — never panics.
//
// Returns:
//   - the removed integer value
//   - true if an element was removed, false if the stack was empty
func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	val := s.data[0]
	s.data = s.data[1:]
	return val, true
}

// Peek returns the top element without removing it.
// Returns (0, false) if the stack is empty — never panics.
//
// Returns:
//   - the top integer value
//   - true if an element exists, false if the stack was empty
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

// Values returns a defensive copy of the stack contents as a slice.
// Index 0 of the returned slice corresponds to the top of the stack.
// Mutating the returned slice does not affect the stack.
func (s *Stack) Values() []int {
	cp := make([]int, len(s.data))
	copy(cp, s.data)
	return cp
}
