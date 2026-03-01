package stack

import "testing"

func TestStackBasics(t *testing.T) {
	vals := []int{1, 2, 3}

	// NewStack should copy the slice
	s := NewStack(vals)
	vals[0] = 99
	if s.Len() != 3 {
		t.Fatalf("expected len 3, got %d", s.Len())
	}

	got := s.Values()
	if got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("expected values [1 2 3], got %v", got)
	}

	// Peek
	if v, ok := s.Peek(); !ok || v != 1 {
		t.Fatalf("unexpected peek: %v %v", v, ok)
	}

	// Pop
	if v, ok := s.Pop(); !ok || v != 1 {
		t.Fatalf("unexpected pop: %v %v", v, ok)
	}
	if s.Len() != 2 {
		t.Fatalf("expected len 2 after pop, got %d", s.Len())
	}

	// Push
	s.Push(7)
	if v, ok := s.Peek(); !ok || v != 7 {
		t.Fatalf("unexpected peek after push: %v %v", v, ok)
	}

	// Pop all and check IsEmpty
	_, _ = s.Pop()
	_, _ = s.Pop()
	_, _ = s.Pop()
	if !s.IsEmpty() {
		t.Fatalf("expected empty stack")
	}

	// Peek/Pop on empty
	if _, ok := s.Peek(); ok {
		t.Fatalf("expected peek on empty to be false")
	}
	if _, ok := s.Pop(); ok {
		t.Fatalf("expected pop on empty to be false")
	}
}

func TestValuesImmutability(t *testing.T) {
	a := NewStack([]int{1, 2, 3})
	vals := a.Values()
	vals[0] = 99
	// original stack should be unchanged
	if v := a.Values()[0]; v != 1 {
		t.Fatalf("expected original stack unchanged, got %v", a.Values())
	}
}
