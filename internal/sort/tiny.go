package sort

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// sortTiny sorts a stack of 2 or 3 elements using a hardcoded decision tree.
// Returns the operation names executed. Does nothing for 0 or 1 elements.
func sortTiny(a *stack.Stack) []string {
	switch a.Len() {
	case 2:
		return sortTwo(a)
	case 3:
		return sortThree(a)
	}
	return nil
}

// sortTwo sorts a stack of exactly 2 elements.
func sortTwo(a *stack.Stack) []string {
	vals := a.Values()
	if vals[0] > vals[1] {
		operations.Sa(a)
		return []string{"sa"}
	}
	return nil
}

// sortThree sorts a stack of exactly 3 elements using at most 2 operations.
func sortThree(a *stack.Stack) []string {
	var ops []string
	v := a.Values()
	top, mid, bot := v[0], v[1], v[2]

	if top > mid && mid < bot && top < bot {
		// e.g. 2 1 3 → sa
		operations.Sa(a)
		ops = append(ops, "sa")
	} else if top > mid && mid > bot {
		// e.g. 3 2 1 → sa + rra
		operations.Sa(a)
		ops = append(ops, "sa")
		operations.Rra(a)
		ops = append(ops, "rra")
	} else if top > bot && mid < bot {
		// e.g. 3 1 2 → ra
		operations.Ra(a)
		ops = append(ops, "ra")
	} else if top < mid && mid > bot && top < bot {
		// e.g. 1 3 2 → sa + ra
		operations.Sa(a)
		ops = append(ops, "sa")
		operations.Ra(a)
		ops = append(ops, "ra")
	} else if top < mid && top > bot {
		// e.g. 2 3 1 → rra
		operations.Rra(a)
		ops = append(ops, "rra")
	}
	// top < mid && mid < bot → already sorted, nothing to do

	return ops
}
