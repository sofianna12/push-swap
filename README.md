# push-swap

A sorting algorithm project in Go. Given a list of integers on two stacks and a set of 11 operations, find the minimum instruction sequence to sort stack `a` in ascending order.

## Programs

| Binary | Description |
|---|---|
| `push-swap` | Reads integers from arguments, outputs the minimum instruction sequence to sort them |
| `checker` | Reads integers from arguments, reads instructions from stdin, outputs `OK` or `KO` |

## Operations

| Instruction | Effect |
|---|---|
| `sa` / `sb` / `ss` | Swap the top 2 elements of stack a / b / both |
| `pa` / `pb` | Push top of b onto a / top of a onto b |
| `ra` / `rb` / `rr` | Rotate up (first becomes last) in a / b / both |
| `rra` / `rrb` / `rrr` | Reverse rotate (last becomes first) in a / b / both |

## Usage

```bash
# Build
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker

# Sort 5 numbers
./push-swap "4 67 3 87 23"

# Verify correctness
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
# → OK

# Error cases
./push-swap "0 one 2 3"    # → Error (non-integer)
./push-swap "1 2 2 3"      # → Error (duplicate)
./push-swap                 # → (no output, nothing to sort)
```

## Performance Targets

| Input size | Max instructions |
|---|---|
| n = 5 | < 12 |
| n = 6 | < 9 |
| n = 100 | < 700 |

## Architecture

```
push-swap/
├── cmd/
│   ├── push-swap/main.go    # parse args → sort → print ops to stdout
│   └── checker/main.go      # parse args → read stdin ops → OK/KO
├── internal/
│   ├── stack/               # Stack type: Push, Pop, Peek, Len, IsSorted
│   ├── operations/          # All 11 operations + Execute dispatcher
│   ├── sort/                # Sort algorithms (n=2–3 hardcoded, n=4–6 small, n>6 large)
│   └── parser/              # ParseArgs: string → []int, validates no dups, no overflow
└── go.mod
```

```
cmd/push-swap → parser, stack, sort
cmd/checker   → parser, stack, operations
sort          → stack, operations
operations    → stack
parser        → stdlib only
stack         → stdlib only
```

## Development

```bash
# Run all tests
go test ./...

# Check formatting (lists unformatted files) and static analysis
gofmt -l .
go vet ./...

# Test coverage
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

## Team

| Member | Role |
|---|---|
| Theo (`@teovaira`) | Leader — `cmd/push-swap`, `cmd/checker`, CI/CD, docs, merges all PRs to `main` |
| Anna (`@sofianna12`) | Developer — stack, operations, small sort |
| Alex (`@arigopou`) | Developer — parser, large sort |

Branch strategy: `feature/<slug>` → PR → Theo merges to `main`.
See [PERMISSIONS.md](PERMISSIONS.md) for the full workflow.
