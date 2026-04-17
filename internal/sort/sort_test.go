package sort

import (
	"fmt"
	"math/rand"
	"testing"

	"push-swap/internal/operations"
	"push-swap/internal/stack"
)

func requireSorted(t *testing.T, nums []int, maxOps int) {
	t.Helper()

	a := stack.New("a", append([]int(nil), nums...))
	b := stack.New("b", nil)
	ops := SortCollect(a, b)

	if maxOps >= 0 && len(ops) > maxOps {
		t.Errorf("SortCollect(%v): %d ops, want <= %d", nums, len(ops), maxOps)
	}

	a2 := stack.New("a", append([]int(nil), nums...))
	b2 := stack.New("b", nil)
	for _, op := range ops {
		if !operations.Execute(op, a2, b2) {
			t.Fatalf("SortCollect(%v): unknown op %q in output", nums, op)
		}
	}

	if !a2.IsSorted() {
		t.Errorf("SortCollect(%v): a not sorted after replay, got %v", nums, a2.Values())
	}
	if b2.Len() != 0 {
		t.Errorf("SortCollect(%v): b not empty after replay, got %v", nums, b2.Values())
	}
}

func TestSort_Zero(t *testing.T) {
	requireSorted(t, []int{}, 0)
}

func TestSort_One(t *testing.T) {
	requireSorted(t, []int{42}, 0)
}

func TestSort_Two(t *testing.T) {
	tests := []struct {
		name string
		in   []int
	}{
		{"already sorted", []int{1, 2}},
		{"reversed", []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireSorted(t, tt.in, 1)
		})
	}
}

func TestSort_Three(t *testing.T) {
	perms := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	for _, p := range perms {
		t.Run(fmt.Sprintf("%v", p), func(t *testing.T) {
			requireSorted(t, p, 2)
		})
	}
}

func TestSort_Four(t *testing.T) {
	tests := [][]int{
		{4, 3, 2, 1},
		{2, 1, 4, 3},
		{3, 1, 2, 4},
		{4, 2, 1, 3},
		{1, 4, 3, 2},
		{3, 4, 1, 2},
	}
	for _, nums := range tests {
		t.Run(fmt.Sprintf("%v", nums), func(t *testing.T) {
			requireSorted(t, nums, 11)
		})
	}
}

func TestSort_Five(t *testing.T) {
	tests := [][]int{
		{5, 4, 3, 2, 1},
		{1, 5, 2, 4, 3},
		{4, 5, 1, 6, 7},
		{4, 5, 6, 7, 1},
		{2, 1, 3, 5, 4},
		{3, 2, 5, 1, 4},
		{5, 1, 2, 3, 4},
		{3, 5, 4, 2, 1},
		{2, 4, 1, 5, 3},
		{1, 3, 5, 2, 4},
	}
	for _, nums := range tests {
		t.Run(fmt.Sprintf("%v", nums), func(t *testing.T) {
			requireSorted(t, nums, 11)
		})
	}
}

func TestSort_Six(t *testing.T) {
	tests := []struct {
		nums   []int
		maxOps int
	}{
		{[]int{2, 1, 3, 6, 5, 8}, 8},
		{[]int{1, 6, 2, 5, 3, 4}, 8},
		{[]int{6, 1, 5, 2, 4, 3}, 8},
		{[]int{3, 6, 1, 4, 2, 5}, 8},
		{[]int{5, 3, 1, 6, 4, 2}, 8},
		{[]int{2, 5, 3, 1, 6, 4}, 8},
		{[]int{4, 1, 3, 5, 2, 6}, 8},
		{[]int{6, 4, 2, 1, 3, 5}, 8},
		{[]int{6, 5, 4, 3, 2, 1}, 10},
		{[]int{4, 2, 6, 1, 5, 3}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			requireSorted(t, tt.nums, tt.maxOps)
		})
	}
}

func TestSort_AlreadySorted(t *testing.T) {
	tests := [][]int{
		{1, 2, 3},
		{1, 2, 3, 4, 5},
		{-3, -2, -1, 0, 1, 2, 3},
	}
	for _, nums := range tests {
		t.Run(fmt.Sprintf("%v", nums), func(t *testing.T) {
			requireSorted(t, nums, 0)
		})
	}
}

func TestSort_ReverseSorted(t *testing.T) {
	tests := [][]int{
		{3, 2, 1},
		{5, 4, 3, 2, 1},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}
	for _, nums := range tests {
		t.Run(fmt.Sprintf("%v", nums), func(t *testing.T) {
			requireSorted(t, nums, -1)
		})
	}
}

func TestSort_100_Random(t *testing.T) {
	seeds := []int64{42, 99, 12345}
	for _, seed := range seeds {
		t.Run(fmt.Sprintf("seed_%d", seed), func(t *testing.T) {
			r := rand.New(rand.NewSource(seed))
			perm := r.Perm(1000)
			nums := perm[:100]
			requireSorted(t, nums, 699)
		})
	}
}

func TestSort_Negatives(t *testing.T) {
	requireSorted(t, []int{-1, -5, -3, -2, -4}, 11)
}
