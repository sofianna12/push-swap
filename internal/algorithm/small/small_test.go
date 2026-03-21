package small

import (
	"push-swap/internal/helper"
	"push-swap/internal/stack"
	"testing"
)

func TestSortTwo(t *testing.T) {
	a := stack.New("a", []int{2, 1})
	b := stack.New("b", []int{})
	ops := Sort(a, b)
	if len(ops) != 1 || ops[0] != "sa" {
		t.Fatalf("expected [sa], got %v", ops)
	}
	if !helper.IsSorted(a.Values()) {
		t.Fatalf("expected sorted a, got %v", a.Values())
	}
}

func TestSortThree(t *testing.T) {
	cases := []struct {
		in []int
	}{
		{[]int{2, 1, 3}},
		{[]int{3, 2, 1}},
		{[]int{3, 1, 2}},
		{[]int{1, 3, 2}},
		{[]int{2, 3, 1}},
	}

	for _, c := range cases {
		a := stack.New("a", c.in)
		b := stack.New("b", []int{})
		ops := Sort(a, b)
		if !helper.IsSorted(a.Values()) {
			t.Fatalf("input %v produced ops %v but final a not sorted: %v", c.in, ops, a.Values())
		}
	}
}

func TestSortFour(t *testing.T) {
	a := stack.New("a", []int{3, 1, 4, 2})
	b := stack.New("b", []int{})
	ops := Sort(a, b)
	if b.Len() != 0 {
		t.Fatalf("expected b empty after ops, got b=%v ops=%v", b.Values(), ops)
	}
	if a.Len() != 4 {
		t.Fatalf("expected a length 4 after ops, got a=%v ops=%v", a.Values(), ops)
	}
	pb, pa := 0, 0
	for _, o := range ops {
		if o == "pb" {
			pb++
		}
		if o == "pa" {
			pa++
		}
	}
	if pb != pa {
		t.Fatalf("mismatched pb/pa: pb=%d pa=%d ops=%v", pb, pa, ops)
	}

	a = stack.New("a", []int{3, 4, 1, 2})
	b = stack.New("b", []int{})
	ops = Sort(a, b)
	if b.Len() != 0 {
		t.Fatalf("expected b empty after ops, got b=%v ops=%v", b.Values(), ops)
	}
	if a.Len() != 4 {
		t.Fatalf("expected a length 4 after ops, got a=%v ops=%v", a.Values(), ops)
	}
	pb, pa = 0, 0
	for _, o := range ops {
		if o == "pb" {
			pb++
		}
		if o == "pa" {
			pa++
		}
	}
	if pb != pa {
		t.Fatalf("mismatched pb/pa: pb=%d pa=%d ops=%v", pb, pa, ops)
	}

	a = stack.New("a", []int{4, 3, 2, 1})
	b = stack.New("b", []int{})
	ops = Sort(a, b)
	if b.Len() != 0 {
		t.Fatalf("expected b empty after ops, got b=%v ops=%v", b.Values(), ops)
	}
	if a.Len() != 4 {
		t.Fatalf("expected a length 4 after ops, got a=%v ops=%v", a.Values(), ops)
	}
	pb, pa = 0, 0
	for _, o := range ops {
		if o == "pb" {
			pb++
		}
		if o == "pa" {
			pa++
		}
	}
	if pb != pa {
		t.Fatalf("mismatched pb/pa: pb=%d pa=%d ops=%v", pb, pa, ops)
	}
}

func TestSortFive(t *testing.T) {
	a := stack.New("a", []int{4, 5, 1, 6, 7})
	b := stack.New("b", []int{})
	ops := Sort(a, b)
	if b.Len() != 0 {
		t.Fatalf("expected b empty after ops, got b=%v ops=%v", b.Values(), ops)
	}
	if a.Len() != 5 {
		t.Fatalf("expected a length 5 after ops, got a=%v ops=%v", a.Values(), ops)
	}
	pb, pa := 0, 0
	for _, o := range ops {
		if o == "pb" {
			pb++
		}
		if o == "pa" {
			pa++
		}
	}
	if pb != pa {
		t.Fatalf("mismatched pb/pa: pb=%d pa=%d ops=%v", pb, pa, ops)
	}

	a = stack.New("a", []int{4, 5, 6, 7, 1})
	b = stack.New("b", []int{})
	ops = Sort(a, b)
	if b.Len() != 0 {
		t.Fatalf("expected b empty after ops, got b=%v ops=%v", b.Values(), ops)
	}
	if a.Len() != 5 {
		t.Fatalf("expected a length 5 after ops, got a=%v ops=%v", a.Values(), ops)
	}
	pb, pa = 0, 0
	for _, o := range ops {
		if o == "pb" {
			pb++
		}
		if o == "pa" {
			pa++
		}
	}
	if pb != pa {
		t.Fatalf("mismatched pb/pa: pb=%d pa=%d ops=%v", pb, pa, ops)
	}
}
