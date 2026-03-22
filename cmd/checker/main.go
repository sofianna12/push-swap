package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"push-swap/internal/operations"
	"push-swap/internal/parser"
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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		op := strings.TrimSpace(scanner.Text())
		if op == "" {
			continue
		}
		if !operations.Execute(op, a, b) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if a.IsSorted() && b.Len() == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
