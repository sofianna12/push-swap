package parser

import (
	"math"
	"strconv"
	"testing"
)

func TestParseArgs_OK(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []int
	}{
		{"nil args", nil, []int{}},
		{"no args", []string{}, []int{}},
		{"only whitespace", []string{"   "}, []int{}},
		{"separate args", []string{"1", "2", "3"}, []int{1, 2, 3}},
		{"single quoted arg", []string{"1 2 3"}, []int{1, 2, 3}},
		{"mixed quoted and split", []string{"1 2", "3", "4 5"}, []int{1, 2, 3, 4, 5}},
		{"negative values", []string{"-1 -2", "0", "5"}, []int{-1, -2, 0, 5}},
		{"int boundaries", []string{strconv.Itoa(math.MinInt), strconv.Itoa(math.MaxInt)}, []int{math.MinInt, math.MaxInt}},
		{"order preserved", []string{"3 1 2"}, []int{3, 1, 2}},
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

func TestParseArgs_Invalid(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"non numeric token", []string{"1", "a"}},
		{"float token", []string{"1.2"}},
		{"plus sign only token", []string{"+"}},
		{"sign only token", []string{"-"}},
		{"duplicate numbers", []string{"1 2 3 2"}},
		{"int64 overflow", []string{"9223372036854775808"}},
		{"int64 underflow", []string{"-9223372036854775809"}},
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
