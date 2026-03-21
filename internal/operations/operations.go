// Package operations implements the 11 push-swap operations:
// sa, sb, ss, pa, pb, ra, rb, rr, rra, rrb, rrr.
//
// All operations accept pointers to stack.Stack and mutate them in place.
// Operations never return errors — input validation is the parser's responsibility.
// Operations on stacks with 0 or 1 elements are always silent no-ops.
package operations

import "push-swap/internal/stack"

// Sa swaps the top two elements of stack a.
// Does nothing if a has fewer than two elements.
func Sa(a *stack.Stack) {
	if a.Len() < 2 {
		return
	}
	first, _ := a.Pop()
	second, _ := a.Pop()
	a.Push(first)
	a.Push(second)
}

// Sb swaps the top two elements of stack b.
// Does nothing if b has fewer than two elements.
func Sb(b *stack.Stack) {
	if b.Len() < 2 {
		return
	}
	first, _ := b.Pop()
	second, _ := b.Pop()
	b.Push(first)
	b.Push(second)
}

// Ss executes Sa and Sb simultaneously.
func Ss(a, b *stack.Stack) {
	Sa(a)
	Sb(b)
}

// Pa pops the top element of stack b and pushes it onto stack a.
// Does nothing if b is empty.
func Pa(a, b *stack.Stack) {
	val, ok := b.Pop()
	if !ok {
		return
	}
	a.Push(val)
}

// Pb pops the top element of stack a and pushes it onto stack b.
// Does nothing if a is empty.
func Pb(a, b *stack.Stack) {
	val, ok := a.Pop()
	if !ok {
		return
	}
	b.Push(val)
}

// Ra rotates stack a upward: the top element becomes the bottom.
// Does nothing if a has fewer than two elements.
func Ra(a *stack.Stack) {
	if a.Len() < 2 {
		return
	}
	vals := a.Values()
	for a.Len() > 0 {
		a.Pop()
	}
	// push in reverse order with vals[0] at the bottom
	a.Push(vals[0])
	for i := len(vals) - 1; i >= 1; i-- {
		a.Push(vals[i])
	}
}

// Rb rotates stack b upward: the top element becomes the bottom.
// Does nothing if b has fewer than two elements.
func Rb(b *stack.Stack) {
	if b.Len() < 2 {
		return
	}
	vals := b.Values()
	for b.Len() > 0 {
		b.Pop()
	}
	b.Push(vals[0])
	for i := len(vals) - 1; i >= 1; i-- {
		b.Push(vals[i])
	}
}

// Rr executes Ra and Rb simultaneously.
func Rr(a, b *stack.Stack) {
	Ra(a)
	Rb(b)
}

// Rra reverse-rotates stack a: the bottom element becomes the top.
// Does nothing if a has fewer than two elements.
func Rra(a *stack.Stack) {
	if a.Len() < 2 {
		return
	}
	vals := a.Values()
	last := vals[len(vals)-1]
	for a.Len() > 0 {
		a.Pop()
	}
	for i := len(vals) - 2; i >= 0; i-- {
		a.Push(vals[i])
	}
	a.Push(last)
}

// Rrb reverse-rotates stack b: the bottom element becomes the top.
// Does nothing if b has fewer than two elements.
func Rrb(b *stack.Stack) {
	if b.Len() < 2 {
		return
	}
	vals := b.Values()
	last := vals[len(vals)-1]
	for b.Len() > 0 {
		b.Pop()
	}
	for i := len(vals) - 2; i >= 0; i-- {
		b.Push(vals[i])
	}
	b.Push(last)
}

// Rrr executes Rra and Rrb simultaneously.
func Rrr(a, b *stack.Stack) {
	Rra(a)
	Rrb(b)
}

// Execute applies the named operation to stacks a and b.
// Returns true if the operation name was recognised, false otherwise.
// Used by checker to validate and apply each instruction from stdin.
func Execute(op string, a, b *stack.Stack) bool {
	switch op {
	case "sa":
		Sa(a)
	case "sb":
		Sb(b)
	case "ss":
		Ss(a, b)
	case "pa":
		Pa(a, b)
	case "pb":
		Pb(a, b)
	case "ra":
		Ra(a)
	case "rb":
		Rb(b)
	case "rr":
		Rr(a, b)
	case "rra":
		Rra(a)
	case "rrb":
		Rrb(b)
	case "rrr":
		Rrr(a, b)
	default:
		return false
	}
	return true
}
