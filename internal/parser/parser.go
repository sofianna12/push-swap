// Package parser provides argument parsing and validation for push-swap and checker.
// It converts raw CLI arguments into a validated slice of integers, rejecting
// non-integer tokens, duplicates, and values outside the int range.
package parser

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// ErrInvalidInput is returned when the provided arguments are not valid push-swap
// numbers — bad tokens, duplicates, sign-only values, or overflow.
var ErrInvalidInput = errors.New("invalid input")

// ParseArgs converts raw CLI args into a validated integer slice.
// Supports both quoted input ("1 2 3") and split input ("1", "2", "3"), including mixed forms.
//
// Parameters:
//   - args: os.Args[1:] from the command line.
//
// Returns the parsed integers and nil on success, or nil and ErrInvalidInput
// on bad tokens, duplicates, or overflow. Empty input returns ([]int{}, nil).
func ParseArgs(args []string) ([]int, error) {
	tokens := splitArgs(args)
	if len(tokens) == 0 {
		return []int{}, nil
	}

	out := make([]int, 0, len(tokens))
	seen := make(map[int]struct{}, len(tokens))

	for _, tok := range tokens {
		if tok == "" {
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

// splitArgs flattens CLI args into individual numeric tokens.
//
// Parameters:
//   - args: raw CLI arguments, each possibly containing space-separated values.
//
// Returns a flat slice of individual token strings.
func splitArgs(args []string) []string {
	var tokens []string
	for _, a := range args {
		parts := strings.Fields(a)
		tokens = append(tokens, parts...)
	}
	return tokens
}

// parseInt parses a single token as a base-10 integer.
// Rejects sign-only values ("+", "-"), non-numeric tokens, and values outside int range.
//
// Parameters:
//   - tok: a single whitespace-free token string.
//
// Returns the parsed int64 value, or an error if the token is invalid.
func parseInt(tok string) (int64, error) {
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
