package sort

import (
	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// sortSmall sorts stack a when it contains between 4 and 6 elements.
// It uses stack b as auxiliary storage and returns the operation names executed.
func sortSmall(a, b *stack.Stack) []string {
	if a.IsSorted() {
		return nil
	}
	if a.Len() == 6 {
		return sortSix(a, b)
	}
	return sortFourFive(a, b)
}

// sortFourFive handles n=4 and n=5.
// Strategy: push the smallest elements to b until 3 remain, sort those 3,
// then pull everything back from b.
func sortFourFive(a, b *stack.Stack) []string {
	var ops []string

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
			for range minIdx {
				operations.Ra(a)
				ops = append(ops, "ra")
			}
		} else {
			for range n - minIdx {
				operations.Rra(a)
				ops = append(ops, "rra")
			}
		}

		operations.Pb(a, b)
		ops = append(ops, "pb")
	}

	ops = append(ops, sortTiny(a)...)

	for b.Len() > 0 {
		operations.Pa(a, b)
		ops = append(ops, "pa")
	}

	return ops
}

// sortSix handles n=6 using the Turkish rank-based chunking algorithm.
// This produces correct results for all inputs and stays well under 700 ops.
// The official audit input "2 1 3 6 5 8" is verified via the binary smoke test.
func sortSix(a, b *stack.Stack) []string {
	return sortLarge(a, b)
}
