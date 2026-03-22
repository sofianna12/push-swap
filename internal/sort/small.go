package sort

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// sortSmall sorts stack a when it contains between 2 and 6 elements.
// It uses stack b as auxiliary storage and returns the operation names executed.
//
// Strategy:
//  1. While a has more than 3 elements, push the smallest to b using minimal rotations.
//  2. Sort the remaining 2 or 3 elements in a with a hardcoded decision tree.
//  3. Push everything back from b to a.
func sortSmall(a, b *stack.Stack) []string {
	var ops []string

	// Move the smallest elements to b until only 3 remain in a.
	for a.Len() > 3 {
		vals := a.Values()
		minIdx := 0
		for i, v := range vals {
			if v < vals[minIdx] {
				minIdx = i
			}
		}

		n := a.Len()
		if minIdx <= n/2 {
			for i := 0; i < minIdx; i++ {
				operations.Ra(a)
				ops = append(ops, "ra")
			}
		} else {
			for i := 0; i < n-minIdx; i++ {
				operations.Rra(a)
				ops = append(ops, "rra")
			}
		}

		operations.Pb(a, b)
		ops = append(ops, "pb")
	}

	// Sort the remaining 2 or 3 elements in a.
	ops = append(ops, sortTiny(a)...)

	// Push everything back from b to a.
	for b.Len() > 0 {
		operations.Pa(a, b)
		ops = append(ops, "pa")
	}

	return ops
}
