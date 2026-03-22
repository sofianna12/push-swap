package sort

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// sortTiny sorts a stack of 2 or 3 elements using a hardcoded decision tree.
//
// Parameters:
//   - a: the stack to sort (must have 2 or 3 elements).
//
// Returns the operation names executed, or nil if already sorted.
func sortTiny(a *stack.Stack) []string {
	switch a.Len() {
	case 2:
		return sortTwo(a)
	case 3:
		return sortThree(a)
	}
	return nil
}

// sortTwo sorts a stack of exactly 2 elements. At most 1 swap.
//
// Parameters:
//   - a: the stack to sort (must have exactly 2 elements).
//
// Returns the operation names executed, or nil if already sorted.
func sortTwo(a *stack.Stack) []string {
	vals := a.Values()
	if vals[0] > vals[1] {
		operations.Sa(a)
		return []string{"sa"}
	}
	return nil
}

// sortThree sorts a stack of exactly 3 elements using at most 2 operations.
//
// Parameters:
//   - a: the stack to sort (must have exactly 3 elements).
//
// Returns the operation names executed, or nil if already sorted.
func sortThree(a *stack.Stack) []string {
	var ops []string
	v := a.Values()
	top, mid, bot := v[0], v[1], v[2]

	if top > mid && mid < bot && top < bot {
		operations.Sa(a)
		ops = append(ops, "sa")
	} else if top > mid && mid > bot {
		operations.Sa(a)
		ops = append(ops, "sa")
		operations.Rra(a)
		ops = append(ops, "rra")
	} else if top > bot && mid < bot {
		operations.Ra(a)
		ops = append(ops, "ra")
	} else if top < mid && mid > bot && top < bot {
		operations.Sa(a)
		ops = append(ops, "sa")
		operations.Ra(a)
		ops = append(ops, "ra")
	} else if top < mid && top > bot {
		operations.Rra(a)
		ops = append(ops, "rra")
	}

	return ops
}
