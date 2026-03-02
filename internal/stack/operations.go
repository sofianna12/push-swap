// Package stack provides the core Stack data structure and
// all push-swap operations applied to stacks.
//
// The Stack represents a LIFO collection of integers where
// index 0 corresponds to the top of the stack.
//
// All mutating operations accept pointer receivers and
// return the operation name when executed, or an empty
// string if the operation is a no-op.
package stack

// Sa swaps the first two elements of stack A.
// If stack A contains fewer than two elements, it performs no operation
// and returns an empty string.
//
// Parameters:
//   - a: the stack to operate on.
//
// Returns the operation name "sa" if executed, otherwise an empty string.
func Sa(a *Stack) string {
	if a.Len() < 2 {
		return ""
	}
	a.data[0], a.data[1] = a.data[1], a.data[0]
	return "sa"
}

// Sb swaps the first two elements of stack B.
// If the stack contains fewer than two elements, Sb does nothing
// and returns an empty string.

func Sb(b *Stack) string {
	if b.Len() < 2 {
		return ""
	}
	b.data[0], b.data[1] = b.data[1], b.data[0]
	return "sb"
}

// Ss swaps the first two elements of both stacks.
// The operation is applied independently to each stack.
// Ss returns "ss" if at least one swap is performed;
// otherwise it returns an empty string.
func Ss(a, b *Stack) string {
	Sa(a)
	Sb(b)
	return "ss"
}

// Pa pops the top element from stack B and pushes it onto stack A.
// If stack B is empty, Pa does nothing and returns an empty string.
//
// Pa returns "pa" if a value is moved.
func Pa(a, b *Stack) string {
	val, ok := b.Pop()
	if !ok {
		return ""
	}
	a.Push(val)
	return "pa"
}

// Pb pops the top element from stack A and pushes it onto stack B.
// If stack A is empty, Pb does nothing and returns an empty string.
//
// Pb returns "pb" if a value is moved.
func Pb(a, b *Stack) string {
	val, ok := a.Pop()
	if !ok {
		return ""
	}
	b.Push(val)
	return "pb"
}

// Ra rotates stack A upward.
// The first element becomes the last element.
// If the stack contains fewer than two elements, Ra does nothing
// and returns an empty string.
//
// Ra returns "ra" if the rotation is performed.
func Ra(a *Stack) string {
	if a.Len() < 2 {
		return ""
	}
	first := a.data[0]
	a.data = append(a.data[1:], first)
	return "ra"
}

// Rb rotates stack B upward.
// The first element becomes the last element.
// If the stack contains fewer than two elements, Rb does nothing
// and returns an empty string.
//
// Rb returns "rb" if the rotation is performed.
func Rb(b *Stack) string {
	if b.Len() < 2 {
		return ""
	}
	first := b.data[0]
	b.data = append(b.data[1:], first)
	return "rb"
}

// Rr rotates both stack A and stack B upward.
// Each stack is rotated independently.
// Rr returns "rr" if at least one rotation is performed
// otherwise it returns an empty string.
func Rr(a, b *Stack) string {
	Ra(a)
	Rb(b)
	return "rr"
}

// Rra performs a reverse rotation on stack A.
// The last element becomes the first element.
// If the stack contains fewer than two elements, Rra does nothing
// and returns an empty string.
//
// Rra returns "rra" if the rotation is performed.
func Rra(a *Stack) string {
	if a.Len() < 2 {
		return ""
	}
	last := a.data[len(a.data)-1]
	a.data = append([]int{last}, a.data[:len(a.data)-1]...)
	return "rra"
}

// Rrb performs a reverse rotation on stack B.
// The last element becomes the first element.
// If the stack contains fewer than two elements, Rrb does nothing
// and returns an empty string.
//
// Rrb returns "rrb" if the rotation is performed.
func Rrb(b *Stack) string {
	if b.Len() < 2 {
		return ""
	}
	last := b.data[len(b.data)-1]
	b.data = append([]int{last}, b.data[:len(b.data)-1]...)
	return "rrb"
}

// Rrr performs a reverse rotation on both stack A and stack B.
// Each stack is rotated independently.
// Rrr returns "rrr" if at least one rotation is performed;
// otherwise it returns an empty string.
func Rrr(a, b *Stack) string {
	Rra(a)
	Rrb(b)
	return "rrr"
}
