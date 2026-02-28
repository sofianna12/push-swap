package large

import (
	"math/rand"
	"testing"

	"push-swap/internal/helper"
	"push-swap/internal/stack"
)

// applyOps replays operation strings against stacks A and B.
//
// This keeps tests focused on algorithm output while reusing real stack behavior.
func applyOps(a, b *stack.Stack, ops []string) {
	for _, op := range ops {
		switch op {
		case "sa":
			stack.Sa(a)
		case "sb":
			stack.Sb(b)
		case "ss":
			stack.Ss(a, b)
		case "pa":
			stack.Pa(a, b)
		case "pb":
			stack.Pb(a, b)
		case "ra":
			stack.Ra(a)
		case "rb":
			stack.Rb(b)
		case "rr":
			stack.Rr(a, b)
		case "rra":
			stack.Rra(a)
		case "rrb":
			stack.Rrb(b)
		case "rrr":
			stack.Rrr(a, b)
		case "":
			// ignore
		default:
			panic("unknown op: " + op)
		}
	}
}

// TestSortLarge_SortsAndEmptiesB validates the basic contract of Sort:
// A must be sorted and B must end empty.
func TestSortLarge_SortsAndEmptiesB(t *testing.T) {
	a := stack.New()
	b := stack.New()
	for _, v := range []int{3, 1, 5, 2, 4, 0} {
		a.Push(v)
	}

	ops := Sort(a, b)
	applyOps(a, b, ops)

	if b.Len() != 0 {
		t.Fatalf("expected B empty, got %v", b.Values())
	}
	if !helper.IsSorted(a.Values()) {
		t.Fatalf("expected A sorted, got %v", a.Values())
	}
}

// TestSortLarge_Random100ReasonableOps checks correctness on a deterministic
// random 100-number input and ensures the operation count stays in a sane range.
func TestSortLarge_Random100ReasonableOps(t *testing.T) {
	const n = 100
	r := rand.New(rand.NewSource(42))

	// make permutation-like unique values
	values := make([]int, 0, n)
	used := map[int]struct{}{}
	for len(values) < n {
		v := r.Intn(10000) - 5000
		if _, ok := used[v]; ok {
			continue
		}
		used[v] = struct{}{}
		values = append(values, v)
	}

	a := stack.New()
	b := stack.New()
	for _, v := range values {
		a.Push(v)
	}

	ops := Sort(a, b)
	applyOps(a, b, ops)

	if b.Len() != 0 {
		t.Fatalf("expected B empty")
	}
	if !helper.IsSorted(a.Values()) {
		t.Fatalf("expected sorted A")
	}
	// This is a heuristic threshold; adjust if your spec gives a target.
	if len(ops) > 1200 {
		t.Fatalf("too many ops: %d", len(ops))
	}
}
