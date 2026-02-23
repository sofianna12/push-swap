package small

import "push-swap/internal/stack"

// Sort sorts small stacks (2–5 elements) and returns the operations
func Sort(a, b *stack.Stack) []string {
	var ops []string
	n := a.Len()

	// 2 elements
	if n == 2 {
		if a.Values()[0] > a.Values()[1] {
			ops = append(ops, stack.Sa(a))
		}
		return ops
	}

	// 3 elements
	if n == 3 {
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
		return ops
	}

	// 4–5 elements
	// strategy: push the smallest elements to b, sort remaining 3 in a, then push back
	for a.Len() > 3 {
		vals := a.Values()
		// find index of smallest value
		minIdx := 0
		minVal := vals[0]
		for i, v := range vals {
			if v < minVal {
				minVal = v
				minIdx = i
			}
		}

		// rotate a until minVal is on top
		if minIdx == 1 {
			ops = append(ops, stack.Ra(a))
		} else if minIdx == 2 && a.Len() == 4 {
			ops = append(ops, stack.Rra(a))
		} else if minIdx == 2 && a.Len() == 5 {
			ops = append(ops, stack.Ra(a))
			ops = append(ops, stack.Ra(a))
		} else if minIdx == 3 {
			ops = append(ops, stack.Rra(a))
		} else if minIdx == 4 {
			ops = append(ops, stack.Rra(a))
			ops = append(ops, stack.Rra(a))
		}

		// push smallest to b
		ops = append(ops, stack.Pb(a, b))
	}

	// sort remaining 3 elements in a
	ops = append(ops, Sort(a, b)...)

	// push back from b to a
	for !b.IsEmpty() {
		ops = append(ops, stack.Pa(a, b))
	}

	return ops
}
