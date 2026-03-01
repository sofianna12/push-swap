package stack

import "testing"

func TestSwapOps(t *testing.T) {
	a := NewStack([]int{1})
	if got := Sa(a); got != "" {
		t.Fatalf("expected empty op for Sa on len<2, got %q", got)
	}

	a = NewStack([]int{1, 2})
	if got := Sa(a); got != "sa" {
		t.Fatalf("expected sa, got %q", got)
	}
	if v := a.Values(); v[0] != 2 || v[1] != 1 {
		t.Fatalf("unexpected values after sa: %v", v)
	}

	b := NewStack([]int{3, 4})
	if got := Sb(b); got != "sb" {
		t.Fatalf("expected sb, got %q", got)
	}

	// ss should call both and return "ss"
	a = NewStack([]int{9, 8})
	b = NewStack([]int{7, 6})
	if got := Ss(a, b); got != "ss" {
		t.Fatalf("expected ss, got %q", got)
	}
}

func TestPushPopOps(t *testing.T) {
	a := NewStack([]int{})
	b := NewStack([]int{5})
	if got := Pa(a, b); got != "pa" {
		t.Fatalf("expected pa, got %q", got)
	}
	if a.Len() != 1 || b.Len() != 0 {
		t.Fatalf("unexpected lens after pa: a=%d b=%d", a.Len(), b.Len())
	}

	// pb when a empty should return empty
	a = NewStack([]int{})
	b = NewStack([]int{})
	if got := Pb(a, b); got != "" {
		t.Fatalf("expected empty pb, got %q", got)
	}
}

func TestRotateOps(t *testing.T) {
	a := NewStack([]int{1, 2, 3})
	if got := Ra(a); got != "ra" {
		t.Fatalf("expected ra, got %q", got)
	}
	if v := a.Values(); v[0] != 2 || v[1] != 3 || v[2] != 1 {
		t.Fatalf("unexpected values after ra: %v", v)
	}

	b := NewStack([]int{4, 5, 6})
	if got := Rb(b); got != "rb" {
		t.Fatalf("expected rb, got %q", got)
	}

	// rra
	a = NewStack([]int{1, 2, 3})
	if got := Rra(a); got != "rra" {
		t.Fatalf("expected rra, got %q", got)
	}
	if v := a.Values(); v[0] != 3 || v[1] != 1 || v[2] != 2 {
		t.Fatalf("unexpected values after rra: %v", v)
	}

	// combined ops
	a = NewStack([]int{1, 2})
	b = NewStack([]int{3, 4})
	if got := Rr(a, b); got != "rr" {
		t.Fatalf("expected rr, got %q", got)
	}
	if got := Rrr(a, b); got != "rrr" {
		t.Fatalf("expected rrr, got %q", got)
	}
}

func TestPushPopAdditional(t *testing.T) {
	// pb when a has value should push to b
	a := NewStack([]int{10})
	b := NewStack([]int{})
	if got := Pb(a, b); got != "pb" {
		t.Fatalf("expected pb, got %q", got)
	}
	if a.Len() != 0 || b.Len() != 1 {
		t.Fatalf("unexpected lens after pb: a=%d b=%d", a.Len(), b.Len())
	}

	// pa when b empty should be no-op
	a = NewStack([]int{})
	b = NewStack([]int{})
	if got := Pa(a, b); got != "" {
		t.Fatalf("expected empty pa when b empty, got %q", got)
	}
}

func TestRotateNoOps(t *testing.T) {
	// Ra on len<2
	a := NewStack([]int{1})
	if got := Ra(a); got != "" {
		t.Fatalf("expected empty ra on len<2, got %q", got)
	}
	// Rb on len<2
	b := NewStack([]int{2})
	if got := Rb(b); got != "" {
		t.Fatalf("expected empty rb on len<2, got %q", got)
	}
	// Rrb on len<2
	if got := Rrb(b); got != "" {
		t.Fatalf("expected empty rrb on len<2, got %q", got)
	}
	// Rr and Rrr should still run even if they are no-ops for small stacks
	a = NewStack([]int{1})
	b = NewStack([]int{2})
	if got := Rr(a, b); got != "rr" {
		t.Fatalf("expected rr, got %q", got)
	}
	if got := Rrr(a, b); got != "rrr" {
		t.Fatalf("expected rrr, got %q", got)
	}
}
