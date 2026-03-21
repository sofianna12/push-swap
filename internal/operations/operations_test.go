package operations

import (
	"testing"

	"push-swap/internal/stack"
)

func TestSa(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"swap two", []int{1, 2}, []int{2, 1}},
		{"swap three", []int{1, 2, 3}, []int{2, 1, 3}},
		{"single no-op", []int{5}, []int{5}},
		{"empty no-op", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := stack.New("a", tt.in)
			Sa(a)
			if got := a.Values(); !equalSlice(got, tt.want) {
				t.Fatalf("Sa(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestSb(t *testing.T) {
	b := stack.New("b", []int{3, 4, 5})
	Sb(b)
	if got := b.Values(); !equalSlice(got, []int{4, 3, 5}) {
		t.Fatalf("Sb([3 4 5]) = %v, want [4 3 5]", got)
	}
}

func TestSs(t *testing.T) {
	a := stack.New("a", []int{1, 2})
	b := stack.New("b", []int{3, 4})
	Ss(a, b)
	if got := a.Values(); !equalSlice(got, []int{2, 1}) {
		t.Fatalf("Ss a = %v, want [2 1]", got)
	}
	if got := b.Values(); !equalSlice(got, []int{4, 3}) {
		t.Fatalf("Ss b = %v, want [4 3]", got)
	}
}

func TestPa(t *testing.T) {
	a := stack.New("a", []int{1})
	b := stack.New("b", []int{9, 8})
	Pa(a, b)
	if got := a.Values(); !equalSlice(got, []int{9, 1}) {
		t.Fatalf("Pa a = %v, want [9 1]", got)
	}
	if got := b.Values(); !equalSlice(got, []int{8}) {
		t.Fatalf("Pa b = %v, want [8]", got)
	}

	// no-op when b is empty
	Pa(a, b)
	Pa(a, b) // b now empty
	Pa(a, b) // should be no-op
	if a.Len() != 3 {
		t.Fatalf("Pa on empty b should be no-op, got len %d", a.Len())
	}
}

func TestPb(t *testing.T) {
	a := stack.New("a", []int{5, 6})
	b := stack.New("b", []int{})
	Pb(a, b)
	if got := a.Values(); !equalSlice(got, []int{6}) {
		t.Fatalf("Pb a = %v, want [6]", got)
	}
	if got := b.Values(); !equalSlice(got, []int{5}) {
		t.Fatalf("Pb b = %v, want [5]", got)
	}

	// no-op when a is empty
	Pb(a, b)
	Pb(a, b) // a now empty
	Pb(a, b) // should be no-op
	if b.Len() != 2 {
		t.Fatalf("Pb on empty a should be no-op, got len %d", b.Len())
	}
}

func TestRa(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"three elements", []int{1, 2, 3}, []int{2, 3, 1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"single no-op", []int{5}, []int{5}},
		{"empty no-op", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := stack.New("a", tt.in)
			Ra(a)
			if got := a.Values(); !equalSlice(got, tt.want) {
				t.Fatalf("Ra(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestRb(t *testing.T) {
	b := stack.New("b", []int{1, 2, 3})
	Rb(b)
	if got := b.Values(); !equalSlice(got, []int{2, 3, 1}) {
		t.Fatalf("Rb([1 2 3]) = %v, want [2 3 1]", got)
	}
}

func TestRr(t *testing.T) {
	a := stack.New("a", []int{1, 2, 3})
	b := stack.New("b", []int{4, 5, 6})
	Rr(a, b)
	if got := a.Values(); !equalSlice(got, []int{2, 3, 1}) {
		t.Fatalf("Rr a = %v, want [2 3 1]", got)
	}
	if got := b.Values(); !equalSlice(got, []int{5, 6, 4}) {
		t.Fatalf("Rr b = %v, want [5 6 4]", got)
	}
}

func TestRra(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"three elements", []int{1, 2, 3}, []int{3, 1, 2}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"single no-op", []int{5}, []int{5}},
		{"empty no-op", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := stack.New("a", tt.in)
			Rra(a)
			if got := a.Values(); !equalSlice(got, tt.want) {
				t.Fatalf("Rra(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestRrb(t *testing.T) {
	b := stack.New("b", []int{1, 2, 3})
	Rrb(b)
	if got := b.Values(); !equalSlice(got, []int{3, 1, 2}) {
		t.Fatalf("Rrb([1 2 3]) = %v, want [3 1 2]", got)
	}
}

func TestRrr(t *testing.T) {
	a := stack.New("a", []int{1, 2, 3})
	b := stack.New("b", []int{4, 5, 6})
	Rrr(a, b)
	if got := a.Values(); !equalSlice(got, []int{3, 1, 2}) {
		t.Fatalf("Rrr a = %v, want [3 1 2]", got)
	}
	if got := b.Values(); !equalSlice(got, []int{6, 4, 5}) {
		t.Fatalf("Rrr b = %v, want [6 4 5]", got)
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		op     string
		wantOK bool
	}{
		{"sa", true},
		{"sb", true},
		{"ss", true},
		{"pa", true},
		{"pb", true},
		{"ra", true},
		{"rb", true},
		{"rr", true},
		{"rra", true},
		{"rrb", true},
		{"rrr", true},
		{"xx", false},
		{"SA", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.op, func(t *testing.T) {
			a := stack.New("a", []int{1, 2, 3})
			b := stack.New("b", []int{4, 5, 6})
			got := Execute(tt.op, a, b)
			if got != tt.wantOK {
				t.Fatalf("Execute(%q) = %v, want %v", tt.op, got, tt.wantOK)
			}
		})
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
