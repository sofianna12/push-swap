// Package sort implements the push-swap sorting algorithms.
//
// Sort dispatches to the optimal algorithm based on stack size:
//   - n=0,1: nothing
//   - n=2,3: sortTiny (hardcoded decision tree, at most 2 ops)
//   - n=4,5: sortFourFive (push minimum to b, sort remaining 3, pull back)
//   - n=6: sortSix (BFS over all permutations, guaranteed shortest sequence)
//   - n>6: sortLarge (Turkish rank-based chunking algorithm)
package sort

import (
	"fmt"
	"io"

	"push-swap/internal/stack"
)

// Sort sorts stack a using stack b as auxiliary storage.
// Operation names are written to w one per line after sorting completes.
//
// Parameters:
//   - a: the primary stack to sort in ascending order.
//   - b: the auxiliary stack, must be empty on entry.
//   - w: writer for operation names (one per line).
//
// Returns the number of operations performed.
func Sort(a, b *stack.Stack, w io.Writer) int {
	ops := SortCollect(a, b)
	for _, op := range ops {
		fmt.Fprintln(w, op) //nolint:errcheck
	}
	return len(ops)
}

// SortCollect sorts stack a using stack b and returns the operation names.
// Used by tests to verify both correctness and operation count.
//
// Parameters:
//   - a: the primary stack to sort in ascending order.
//   - b: the auxiliary stack, must be empty on entry.
//
// Returns the slice of operation names executed, or nil if already sorted.
func SortCollect(a, b *stack.Stack) []string {
	switch a.Len() {
	case 0, 1:
		return nil
	case 2, 3:
		return sortTiny(a)
	case 4, 5, 6:
		return sortSmall(a, b)
	default:
		return sortLarge(a, b)
	}
}
