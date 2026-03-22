package stack

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantLen int
	}{
		{"empty", []int{}, 0},
		{"single", []int{5}, 1},
		{"multiple", []int{1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New("a", tt.nums)
			if s.Len() != tt.wantLen {
				t.Fatalf("New(%v).Len() = %d, want %d", tt.nums, s.Len(), tt.wantLen)
			}
		})
	}
}

func TestNew_DefensiveCopy(t *testing.T) {
	nums := []int{1, 2, 3}
	s := New("a", nums)
	nums[0] = 99
	if got := s.Values()[0]; got != 1 {
		t.Fatalf("New should copy input slice, but top changed to %d", got)
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		push    int
		want    []int
	}{
		{"push onto empty", []int{}, 7, []int{7}},
		{"push onto non-empty", []int{2, 3}, 1, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New("a", tt.initial)
			s.Push(tt.push)
			if got := s.Values(); !equalSlice(got, tt.want) {
				t.Fatalf("Push(%d) on %v = %v, want %v", tt.push, tt.initial, got, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		wantVal int
		wantOK  bool
		wantLen int
	}{
		{"non-empty", []int{1, 2, 3}, 1, true, 2},
		{"single element", []int{5}, 5, true, 0},
		{"empty", []int{}, 0, false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New("a", tt.initial)
			v, ok := s.Pop()
			if v != tt.wantVal || ok != tt.wantOK {
				t.Fatalf("Pop() = (%d, %v), want (%d, %v)", v, ok, tt.wantVal, tt.wantOK)
			}
			if s.Len() != tt.wantLen {
				t.Fatalf("Len after Pop = %d, want %d", s.Len(), tt.wantLen)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		wantVal int
		wantOK  bool
	}{
		{"non-empty", []int{7, 8, 9}, 7, true},
		{"single element", []int{3}, 3, true},
		{"empty", []int{}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New("a", tt.initial)
			v, ok := s.Peek()
			if v != tt.wantVal || ok != tt.wantOK {
				t.Fatalf("Peek() = (%d, %v), want (%d, %v)", v, ok, tt.wantVal, tt.wantOK)
			}
			if s.Len() != len(tt.initial) {
				t.Fatalf("Peek must not modify stack, Len = %d, want %d", s.Len(), len(tt.initial))
			}
		})
	}
}

func TestLen(t *testing.T) {
	s := New("a", []int{1, 2, 3})
	if s.Len() != 3 {
		t.Fatalf("Len = %d, want 3", s.Len())
	}
	s.Pop()
	if s.Len() != 2 {
		t.Fatalf("Len after Pop = %d, want 2", s.Len())
	}
	s.Push(99)
	if s.Len() != 3 {
		t.Fatalf("Len after Push = %d, want 3", s.Len())
	}
}

func TestValues_ImmutabilityOfReturn(t *testing.T) {
	s := New("a", []int{1, 2, 3})
	v := s.Values()
	v[0] = 99
	if got := s.Values()[0]; got != 1 {
		t.Fatalf("Values() should return a copy, but stack top changed to %d", got)
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want bool
	}{
		{"empty", []int{}, true},
		{"single", []int{5}, true},
		{"sorted ascending", []int{1, 2, 3, 4}, true},
		{"unsorted", []int{2, 1, 3}, false},
		{"reverse sorted", []int{3, 2, 1}, false},
		{"two sorted", []int{1, 2}, true},
		{"two unsorted", []int{2, 1}, false},
		{"equal adjacent", []int{1, 1, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New("a", tt.in)
			if got := s.IsSorted(); got != tt.want {
				t.Fatalf("IsSorted(%v) = %v, want %v", tt.in, got, tt.want)
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
