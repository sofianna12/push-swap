package parser

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// ErrInvalidInput is returned when the provided arguments are not valid push-swap numbers.
//
// This covers bad tokens (like "one"), duplicates, sign-only values, and numbers
// outside int boundaries.
var ErrInvalidInput = errors.New("invalid input")

// ParseArgs converts raw CLI args into a validated integer slice.
//
// It supports both quoted input ("1 2 3") and split input (1 2 3), including mixed forms.
// Empty input returns an empty slice with no error.
func ParseArgs(args []string) ([]int, error) {
	tokens := splitArgs(args)
	if len(tokens) == 0 {
		return []int{}, nil
	}

	out := make([]int, 0, len(tokens))
	seen := make(map[int]struct{}, len(tokens))

	for _, tok := range tokens {
		if tok == "" {
			// strings.Fields already removes empty pieces, but we keep this guard
			// for defensive parsing.
			continue
		}

		n64, err := parseInt(tok)
		if err != nil {
			return nil, ErrInvalidInput
		}
		n := int(n64)

		if _, ok := seen[n]; ok {
			return nil, ErrInvalidInput
		}
		seen[n] = struct{}{}
		out = append(out, n)
	}

	return out, nil
}

// splitArgs flattens CLI args into numeric tokens.
//
// Each incoming argument may contain one value or many values separated by spaces.
func splitArgs(args []string) []string {
	var tokens []string
	for _, a := range args {
		parts := strings.Fields(a) // splits on any whitespace, ignores repeats
		tokens = append(tokens, parts...)
	}
	return tokens
}

// parseInt parses a single token as base-10 int.
//
// It rejects sign-only values, non-integers, and out-of-range numbers.
func parseInt(tok string) (int64, error) {
	// ParseInt with bitSize=64 allows full parsing first; int bounds are checked below.
	if tok == "+" || tok == "-" {
		return 0, errors.New("sign only")
	}
	n, err := strconv.ParseInt(tok, 10, 64)
	if err != nil {
		return 0, err
	}
	if n < int64(math.MinInt) || n > int64(math.MaxInt) {
		return 0, errors.New("out of int range")
	}
	return n, nil
}
