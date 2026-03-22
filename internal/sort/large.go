package sort

import (
	"sort"

	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// sortLarge sorts stack a when it contains more than 6 elements.
// It uses stack b as auxiliary storage and returns the operation names executed.
//
// Strategy — rank-based chunking (Turkish algorithm):
//  1. Compress values to ranks 0..n-1 so chunk boundaries are uniform.
//  2. Push values from a to b in rank windows (chunks), rotating a to minimise moves.
//  3. Pull from b back to a by always bringing the current maximum rank to the top.
func sortLarge(a, b *stack.Stack) []string {
	vals := a.Values()
	if len(vals) <= 1 || a.IsSorted() {
		return nil
	}

	rankMap := buildRank(vals)

	n := len(vals)
	chunk := 20
	if n > 200 {
		chunk = 35
	}
	if n > 500 {
		chunk = 45
	}

	ops := make([]string, 0, n*8)

	// Push to b in chunks by rank window.
	nextRank := 0
	for nextRank < n {
		limit := nextRank + chunk
		if limit > n {
			limit = n
		}

		for countInWindow(a, rankMap, nextRank, limit) > 0 {
			up, down := distanceToWindow(a, rankMap, nextRank, limit)

			if up <= down {
				for i := 0; i < up; i++ {
					operations.Ra(a)
					ops = append(ops, "ra")
				}
			} else {
				for i := 0; i < down; i++ {
					operations.Rra(a)
					ops = append(ops, "rra")
				}
			}

			operations.Pb(a, b)
			ops = append(ops, "pb")

			// Heuristic: rotate b so bigger ranks stay near the top.
			top, ok := b.Peek()
			if ok {
				mid := (nextRank + limit) / 2
				if rankMap[top] < mid {
					operations.Rb(b)
					ops = append(ops, "rb")
				}
			}
		}

		nextRank = limit
	}

	// Pull from b back to a by always extracting the current maximum rank.
	for b.Len() > 0 {
		up, down := distanceToMax(b, rankMap)
		if up <= down {
			for i := 0; i < up; i++ {
				operations.Rb(b)
				ops = append(ops, "rb")
			}
		} else {
			for i := 0; i < down; i++ {
				operations.Rrb(b)
				ops = append(ops, "rrb")
			}
		}
		operations.Pa(a, b)
		ops = append(ops, "pa")
	}

	return ops
}

// buildRank maps each value to its sorted position (rank) in the input set.
func buildRank(vals []int) map[int]int {
	cp := append([]int(nil), vals...)
	sort.Ints(cp)
	rankMap := make(map[int]int, len(cp))
	for i, v := range cp {
		rankMap[v] = i
	}
	return rankMap
}

// countInWindow returns how many values in a have ranks inside [lo, hi).
func countInWindow(a *stack.Stack, rank map[int]int, lo, hi int) int {
	c := 0
	for _, v := range a.Values() {
		if r := rank[v]; r >= lo && r < hi {
			c++
		}
	}
	return c
}

// distanceToWindow returns the minimum forward (up) and backward (down) rotations
// needed to bring any element with rank in [lo, hi) to the top of a.
func distanceToWindow(a *stack.Stack, rank map[int]int, lo, hi int) (up int, down int) {
	vals := a.Values()
	n := len(vals)
	up = n
	for i, v := range vals {
		if r := rank[v]; r >= lo && r < hi {
			up = i
			break
		}
	}
	down = n
	for i := n - 1; i >= 0; i-- {
		if r := rank[vals[i]]; r >= lo && r < hi {
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

// distanceToMax returns the rotations needed to bring the highest-rank element
// in b to the top using either rb or rrb.
func distanceToMax(b *stack.Stack, rank map[int]int) (up int, down int) {
	vals := b.Values()
	n := len(vals)
	maxIdx := 0
	maxRank := -1
	for i, v := range vals {
		if r := rank[v]; r > maxRank {
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
