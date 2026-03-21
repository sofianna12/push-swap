package large

import (
	"fmt"
	"math/rand"
	"testing"

	"push-swap/internal/helper"
	"push-swap/internal/stack"
)

// applyOps replays operation strings against stacks A and B.
//
// This keeps tests focused on algorithm output while reusing real stack behavior.
func applyOps(t *testing.T, a, b *stack.Stack, ops []string) {
	t.Helper()

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
			t.Fatalf("unknown op: %s", op)
		}
	}
}

func cloneInts(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}

func buildStacks(values []int) (*stack.Stack, *stack.Stack) {
	return stack.New("a", cloneInts(values)), stack.New("b", []int{})
}

// TestSortLarge_SortsAndEmptiesB validates the basic contract of Sort:
// A must be sorted and B must end empty.
func TestSortLarge_SortsAndEmptiesB(t *testing.T) {
	values := []int{3, 1, 5, 2, 4, 0}

	a, b := buildStacks(values)
	ops := Sort(a, b)

	replayA, replayB := buildStacks(values)
	applyOps(t, replayA, replayB, ops)

	if replayB.Len() != 0 {
		t.Fatalf("expected B empty, got %v", replayB.Values())
	}
	if !helper.IsSorted(replayA.Values()) {
		t.Fatalf("expected A sorted, got %v", replayA.Values())
	}
}

// TestSortLarge_Random100ReasonableOps checks correctness on a deterministic
// random 100-number input and ensures the operation count stays in a sane range.
func TestSortLarge_Random100ReasonableOps(t *testing.T) {
	const n = 100
	seeds := []int64{42, 99, 12345}

	for _, seed := range seeds {
		t.Run(fmt.Sprintf("seed_%d", seed), func(t *testing.T) {
			r := rand.New(rand.NewSource(seed))

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

			a, b := buildStacks(values)
			ops := Sort(a, b)

			replayA, replayB := buildStacks(values)
			applyOps(t, replayA, replayB, ops)

			if replayB.Len() != 0 {
				t.Fatalf("expected B empty")
			}
			if !helper.IsSorted(replayA.Values()) {
				t.Fatalf("expected sorted A")
			}
			if len(ops) > 699 {
				t.Fatalf("too many ops for n=100: %d", len(ops))
			}
		})
	}
}

func TestSortLarge_AlreadySorted_NoOps(t *testing.T) {
	values := []int{-2, -1, 0, 1, 2}
	a, b := buildStacks(values)

	ops := Sort(a, b)
	if len(ops) != 0 {
		t.Fatalf("expected no ops for already sorted input, got %d", len(ops))
	}
}

func TestSortLarge_TwoUnsorted(t *testing.T) {
	values := []int{2, 1}
	a, b := buildStacks(values)

	ops := Sort(a, b)

	replayA, replayB := buildStacks(values)
	applyOps(t, replayA, replayB, ops)

	if replayB.Len() != 0 {
		t.Fatalf("expected B empty")
	}
	if !helper.IsSorted(replayA.Values()) {
		t.Fatalf("expected sorted A, got %v", replayA.Values())
	}
}
