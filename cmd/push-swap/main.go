// Package main implements the push-swap binary.
// It reads integers from command-line arguments, validates them, and prints
// the minimum sequence of push-swap operations to sort them onto stack a.
package main

import (
	"fmt"
	"os"

	"push-swap/internal/parser"
	"push-swap/internal/sort"
	"push-swap/internal/stack"
)

func main() {
	nums, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	if len(nums) == 0 {
		return
	}
	a := stack.New("a", nums)
	b := stack.New("b", nil)
	sort.Sort(a, b, os.Stdout)
}
