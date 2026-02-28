package large

import (
	"sort"

	"push-swap/internal/helper"
	"push-swap/internal/stack"
)

// Sort generates push-swap operations that sort stack A in ascending order.
//
// The strategy is rank-based chunking: push values from A to B by rank windows,
// then rebuild A by repeatedly bringing the current max in B to the top.
func Sort(a, b *stack.Stack) []string {
	vals := a.Values()
	if len(vals) <= 1 || helper.IsSorted(vals) {
		return nil
	}

	// Rank compression: value -> rank [0..n-1]
	rank := buildRank(vals)

	// Choose chunk size (common heuristic)
	n := len(vals)
	chunk := 20
	if n > 200 {
		chunk = 35
	}
	if n > 500 {
		chunk = 45
	}

	ops := make([]string, 0, n*8)

	// Push to B in chunks by rank window
	nextRank := 0
	for nextRank < n {
		limit := nextRank + chunk
		if limit > n {
			limit = n
		}

		for countInWindow(a, rank, nextRank, limit) > 0 {
			// Bring an element whose rank is in [nextRank, limit) to top of A
			up, down := distanceToWindow(a, rank, nextRank, limit)

			if up <= down {
				for i := 0; i < up; i++ {
					if op := stack.Ra(a); op != "" {
						ops = append(ops, op)
					}
				}
			} else {
				for i := 0; i < down; i++ {
					if op := stack.Rra(a); op != "" {
						ops = append(ops, op)
					}
				}
			}

			// Push it to B
			if op := stack.Pb(a, b); op != "" {
				ops = append(ops, op)
			}

			// Small heuristic: if the pushed rank is in lower half of the window,
			// rotate B to keep bigger ranks nearer the top.
			top := b.Peek()
			if top != nil {
				r := rank[*top]
				mid := (nextRank + limit) / 2
				if r < mid {
					if op := stack.Rb(b); op != "" {
						ops = append(ops, op)
					}
				}
			}
		}

		nextRank = limit
	}

	// Push back from B to A by always extracting the current max in B
	for b.Len() > 0 {
		up, down := distanceToMax(b, rank)
		if up <= down {
			for i := 0; i < up; i++ {
				if op := stack.Rb(b); op != "" {
					ops = append(ops, op)
				}
			}
		} else {
			for i := 0; i < down; i++ {
				if op := stack.Rrb(b); op != "" {
					ops = append(ops, op)
				}
			}
		}
		if op := stack.Pa(a, b); op != "" {
			ops = append(ops, op)
		}
	}

	return ops
}

// buildRank maps each value to its sorted position (rank) in the input set.
func buildRank(vals []int) map[int]int {
	cp := append([]int(nil), vals...)
	sort.Ints(cp)
	r := make(map[int]int, len(cp))
	for i, v := range cp {
		r[v] = i
	}
	return r
}

// countInWindow returns how many values in A have ranks inside [lo, hi).
func countInWindow(a *stack.Stack, rank map[int]int, lo, hi int) int {
	c := 0
	for _, v := range a.Values() {
		r := rank[v]
		if r >= lo && r < hi {
			c++
		}
	}
	return c
}

// distanceToWindow calculates the shortest forward/backward rotations needed
// to bring any value with rank in [lo, hi) to the top of stack A.
func distanceToWindow(a *stack.Stack, rank map[int]int, lo, hi int) (up int, down int) {
	vals := a.Values()
	n := len(vals)
	up = n
	for i, v := range vals {
		r := rank[v]
		if r >= lo && r < hi {
			up = i
			break
		}
	}
	down = n
	for i := n - 1; i >= 0; i-- {
		r := rank[vals[i]]
		if r >= lo && r < hi {
			down = n - i
			break
		}
	}
	if up == n {
		up = 0
	}
	if down == n {
		down = 0
	}
	return up, down
}

// distanceToMax calculates the rotations needed to move the highest-rank value
// in B to the top using either rb or rrb.
func distanceToMax(b *stack.Stack, rank map[int]int) (up int, down int) {
	vals := b.Values()
	n := len(vals)

	maxIdx := 0
	maxRank := -1
	for i, v := range vals {
		r := rank[v]
		if r > maxRank {
			maxRank = r
			maxIdx = i
		}
	}

	up = maxIdx
	down = n - maxIdx
	if down == n {
		down = 0
	}
	return up, down
}
