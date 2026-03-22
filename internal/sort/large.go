package sort

import (
	"sort"

	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// sortLarge sorts stack a when it contains more than 6 elements using
// rank-based chunking (Turkish algorithm):
//  1. Compress values to ranks 0..n-1 so chunk boundaries are uniform.
//  2. Push values from a to b in rank windows (chunks), rotating a to minimize moves.
//  3. Pull from b back to a by always bringing the current maximum rank to the top.
//
// Parameters:
//   - a: the primary stack (more than 6 elements).
//   - b: the auxiliary stack, must be empty on entry.
//
// Returns the operation names executed, or nil if already sorted.
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
//
// Parameters:
//   - vals: the original stack values.
//
// Returns a map from value to rank (0-indexed).
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
//
// Parameters:
//   - a: the stack to scan.
//   - rank: value-to-rank mapping.
//   - lo, hi: inclusive lower and exclusive upper rank bounds.
//
// Returns the count of matching elements.
func countInWindow(a *stack.Stack, rank map[int]int, lo, hi int) int {
	c := 0
	for _, v := range a.Values() {
		if r := rank[v]; r >= lo && r < hi {
			c++
		}
	}
	return c
}

// distanceToWindow returns the minimum forward and backward rotations
// needed to bring any element with rank in [lo, hi) to the top of a.
//
// Parameters:
//   - a: the stack to scan.
//   - rank: value-to-rank mapping.
//   - lo, hi: inclusive lower and exclusive upper rank bounds.
//
// Returns (up, down): rotations via ra and rra respectively.
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
// in b to the top.
//
// Parameters:
//   - b: the stack to scan.
//   - rank: value-to-rank mapping.
//
// Returns (up, down): rotations via rb and rrb respectively.
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
