package parser

import "testing"

// TestParseArgs_OK checks the most common valid input shapes we expect from CLI usage.
func TestParseArgs_OK(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []int
	}{
		{"no args", []string{}, []int{}},
		{"only whitespace", []string{"   "}, []int{}},
		{"separate args", []string{"1", "2", "3"}, []int{1, 2, 3}},
		{"single quoted arg", []string{"1 2 3"}, []int{1, 2, 3}},
		{"mixed quoted and split", []string{"1 2", "3", "4 5"}, []int{1, 2, 3, 4, 5}},
		{"negative values", []string{"-1 -2", "0", "5"}, []int{-1, -2, 0, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if err != nil {
				t.Fatalf("ParseArgs returned unexpected error: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("parsed length mismatch: got %v want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("value mismatch at index %d: got %v want %v", i, got, tt.want)
				}
			}
		})
	}
}

// TestParseArgs_Invalid checks malformed input and duplicate handling.
func TestParseArgs_Invalid(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"non numeric token", []string{"1", "a"}},
		{"float token", []string{"1.2"}},
		{"sign only token", []string{"-"}},
		{"duplicate numbers", []string{"1 2 3 2"}},
		{"int32 overflow", []string{"2147483648"}},   // max int32 + 1
		{"int32 underflow", []string{"-2147483649"}}, // min int32 - 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseArgs(tt.args)
			if err == nil {
				t.Fatalf("expected ParseArgs to fail")
			}
			if err != ErrInvalidInput {
				t.Fatalf("expected ErrInvalidInput, got %v", err)
			}
		})
	}
}
