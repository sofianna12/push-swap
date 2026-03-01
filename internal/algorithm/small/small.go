// Package small provides sorting algorithms optimized for small stacks
// (2–6 elements) for the push-swap project.
//
// The Sort function generates a minimal sequence of stack operations to
// sort stack A in ascending order using stack B as auxiliary storage.
//
// Strategy:
//  1. While more than 3 elements are in A, push the smallest element
//     to B using minimal rotations.
//  2. Sort the remaining 3 (or 2) elements in A using hardcoded patterns.
//  3. Push back all elements from B to A.
package small

import "push-swap/internal/stack"

// Sort sorts stack A (2–6 elements) using stack B as auxiliary storage.
// Returns the executed push-swap operation names in order.
func Sort(a, b *stack.Stack) []string {
	var ops []string

	// Move the smallest elements to B until only 3 remain in A
	for a.Len() > 3 {
		vals := a.Values()
		// Find the index of the smallest element
		minIdx := 0
		minVal := vals[0]
		for i, v := range vals {
			if v < minVal {
				minVal = v
				minIdx = i
			}
		}

		// Rotate A to bring the smallest element to the top with minimal moves
		if minIdx <= a.Len()/2 {
			for i := 0; i < minIdx; i++ {
				ops = append(ops, stack.Ra(a))
			}
		} else {
			for i := 0; i < a.Len()-minIdx; i++ {
				ops = append(ops, stack.Rra(a))
			}
		}

		// Push the smallest element to B
		ops = append(ops, stack.Pb(a, b))
	}

	// Sort the remaining 3 or 2 elements in A
	if a.Len() == 3 {
		vals := a.Values()
		if vals[0] > vals[1] && vals[1] < vals[2] && vals[0] < vals[2] {
			ops = append(ops, stack.Sa(a))
		} else if vals[0] > vals[1] && vals[1] > vals[2] {
			ops = append(ops, stack.Sa(a))
			ops = append(ops, stack.Rra(a))
		} else if vals[0] > vals[2] && vals[1] < vals[2] {
			ops = append(ops, stack.Ra(a))
		} else if vals[0] < vals[2] && vals[1] > vals[2] {
			ops = append(ops, stack.Sa(a))
			ops = append(ops, stack.Ra(a))
		} else if vals[0] < vals[1] && vals[0] > vals[2] {
			ops = append(ops, stack.Rra(a))
		}
	} else if a.Len() == 2 {
		vals := a.Values()
		if vals[0] > vals[1] {
			ops = append(ops, stack.Sa(a))
		}
	}

	// Push back all elements from B to A
	for !b.IsEmpty() {
		ops = append(ops, stack.Pa(a, b))
	}

	return ops
}
