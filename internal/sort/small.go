package sort

import (
	"fmt"
	gosort "sort"

	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

// sortSmall sorts stack a when it contains between 4 and 6 elements.
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
// Push minimums to b until 3 remain, sort those 3, pull back.
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

// bfsState represents the state of both stacks during BFS exploration.
type bfsState struct {
	av []int // stack a values (top first)
	bv []int // stack b values (top first)
}

// bfsEntry is a BFS queue element holding the state and the ops taken to reach it.
type bfsEntry struct {
	s   bfsState
	ops []string
}

// sortSix uses BFS to find the shortest sequence of operations to sort 6 elements.
// The input is normalised to ranks 0-5 so the BFS explores a bounded state space.
// Since 6! = 720 permutations and the maximum optimal depth is 10, BFS terminates
// quickly (typically under 5000 visited states).
func sortSix(a, b *stack.Stack) []string {
	// Normalise the input to ranks 0..5.
	av := a.Values()
	sorted := make([]int, len(av))
	copy(sorted, av)
	gosort.Ints(sorted)
	rankOf := make(map[int]int, len(av))
	for i, v := range sorted {
		rankOf[v] = i
	}
	ranked := make([]int, len(av))
	for i, v := range av {
		ranked[i] = rankOf[v]
	}

	start := bfsState{av: ranked, bv: nil}
	if bfsIsSorted(start) {
		return nil
	}

	visited := make(map[string]bool, 8192)
	visited[bfsKey(start)] = true

	queue := []bfsEntry{{s: start, ops: nil}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// Maximum depth 12 is sufficient for any 6-element permutation (worst case is 10).
		if len(cur.ops) >= 12 {
			continue
		}

		for _, op := range bfsOpNames {
			ns := bfsApplyOp(cur.s, op)
			k := bfsKey(ns)
			if visited[k] {
				continue
			}
			newOps := make([]string, len(cur.ops)+1)
			copy(newOps, cur.ops)
			newOps[len(cur.ops)] = op

			if bfsIsSorted(ns) {
				replayOps(a, b, newOps)
				return newOps
			}
			visited[k] = true
			queue = append(queue, bfsEntry{s: ns, ops: newOps})
		}
	}

	// Unreachable for valid 6-element input.
	return nil
}

// bfsOpNames lists all 11 push-swap operations used in BFS exploration.
var bfsOpNames = [11]string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

// bfsKey encodes a BFS state as a unique string for the visited map.
func bfsKey(s bfsState) string {
	return fmt.Sprintf("%v|%v", s.av, s.bv)
}

// bfsIsSorted reports whether stack a is sorted ascending and stack b is empty.
func bfsIsSorted(s bfsState) bool {
	if len(s.bv) != 0 {
		return false
	}
	for i := 1; i < len(s.av); i++ {
		if s.av[i-1] > s.av[i] {
			return false
		}
	}
	return true
}

// bfsApplyOp returns the new state after applying a single operation.
// All slices are freshly allocated to avoid aliasing.
func bfsApplyOp(s bfsState, op string) bfsState {
	av := make([]int, len(s.av))
	copy(av, s.av)
	bv := make([]int, len(s.bv))
	copy(bv, s.bv)

	switch op {
	case "sa":
		if len(av) >= 2 {
			av[0], av[1] = av[1], av[0]
		}
	case "sb":
		if len(bv) >= 2 {
			bv[0], bv[1] = bv[1], bv[0]
		}
	case "ss":
		if len(av) >= 2 {
			av[0], av[1] = av[1], av[0]
		}
		if len(bv) >= 2 {
			bv[0], bv[1] = bv[1], bv[0]
		}
	case "pa":
		if len(bv) > 0 {
			nav := make([]int, len(av)+1)
			nav[0] = bv[0]
			copy(nav[1:], av)
			av = nav
			nbv := make([]int, len(bv)-1)
			copy(nbv, bv[1:])
			bv = nbv
		}
	case "pb":
		if len(av) > 0 {
			nbv := make([]int, len(bv)+1)
			nbv[0] = av[0]
			copy(nbv[1:], bv)
			bv = nbv
			nav := make([]int, len(av)-1)
			copy(nav, av[1:])
			av = nav
		}
	case "ra":
		if len(av) >= 2 {
			av = bfsRotate(av)
		}
	case "rb":
		if len(bv) >= 2 {
			bv = bfsRotate(bv)
		}
	case "rr":
		if len(av) >= 2 {
			av = bfsRotate(av)
		}
		if len(bv) >= 2 {
			bv = bfsRotate(bv)
		}
	case "rra":
		if len(av) >= 2 {
			av = bfsReverseRotate(av)
		}
	case "rrb":
		if len(bv) >= 2 {
			bv = bfsReverseRotate(bv)
		}
	case "rrr":
		if len(av) >= 2 {
			av = bfsReverseRotate(av)
		}
		if len(bv) >= 2 {
			bv = bfsReverseRotate(bv)
		}
	}
	return bfsState{av: av, bv: bv}
}

// bfsRotate returns a new slice with the first element moved to the end.
func bfsRotate(s []int) []int {
	n := make([]int, len(s))
	copy(n, s[1:])
	n[len(s)-1] = s[0]
	return n
}

// bfsReverseRotate returns a new slice with the last element moved to the front.
func bfsReverseRotate(s []int) []int {
	n := make([]int, len(s))
	n[0] = s[len(s)-1]
	copy(n[1:], s[:len(s)-1])
	return n
}

// replayOps applies a slice of op names to the real stacks.
func replayOps(a, b *stack.Stack, ops []string) {
	for _, op := range ops {
		switch op {
		case "sa":
			operations.Sa(a)
		case "sb":
			operations.Sb(b)
		case "ss":
			operations.Ss(a, b)
		case "pa":
			operations.Pa(a, b)
		case "pb":
			operations.Pb(a, b)
		case "ra":
			operations.Ra(a)
		case "rb":
			operations.Rb(b)
		case "rr":
			operations.Rr(a, b)
		case "rra":
			operations.Rra(a)
		case "rrb":
			operations.Rrb(b)
		case "rrr":
			operations.Rrr(a, b)
		}
	}
}
