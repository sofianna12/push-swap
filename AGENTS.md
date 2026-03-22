# AGENTS.md — Push-Swap

Context and instructions for AI coding agents working on this project.

---

## Project Overview

Push-swap is a Go sorting algorithm project. It implements two binaries:

- **push-swap** — reads integers from arguments, outputs the minimum instruction sequence to sort them onto stack `a`
- **checker** — reads integers from arguments, reads instructions from stdin, outputs `OK` if the result is sorted or `KO` otherwise

Both programs operate on two stacks (`a` and `b`) using exactly 11 allowed operations.

---

## Repository Structure

```
push-swap/
├── cmd/
│   ├── push-swap/main.go    # push-swap binary entry point
│   └── checker/main.go      # checker binary entry point
├── internal/
│   ├── stack/               # Stack data structure
│   ├── operations/          # All 11 push-swap operations
│   ├── sort/                # Sort algorithm (tiny, small, large)
│   └── parser/              # Argument parsing and validation
├── .github/
│   └── workflows/           # CI, release, lint, coverage, e2e, PR checks
├── go.mod
├── CONTRIBUTING.md
├── CHANGELOG.md
└── Makefile
```

---

## Build

```bash
# Build both binaries
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker

# Verify all packages compile
go build ./...
```

---

## Test

```bash
# Run all tests with race detector
go test -race ./...

# Run with coverage
go test -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

# Run a specific package
go test -race -v ./internal/stack/
go test -race -v ./internal/operations/
go test -race -v ./internal/sort/
go test -race -v ./internal/parser/
```

---

## Lint and Format

```bash
gofmt -w .       # format all Go files
gofmt -l .       # list unformatted files (CI fails if any are listed)
go vet ./...     # static analysis
```

---

## Operations

The 11 allowed stack operations:

| Op | Description |
|---|---|
| `sa` | Swap top 2 elements of stack a |
| `sb` | Swap top 2 elements of stack b |
| `ss` | `sa` and `sb` simultaneously |
| `pa` | Push top of stack b onto stack a |
| `pb` | Push top of stack a onto stack b |
| `ra` | Rotate stack a up (first becomes last) |
| `rb` | Rotate stack b up |
| `rr` | `ra` and `rb` simultaneously |
| `rra` | Reverse rotate stack a (last becomes first) |
| `rrb` | Reverse rotate stack b |
| `rrr` | `rra` and `rrb` simultaneously |

All operations are no-ops on stacks with 0 or 1 elements — they never panic.

---

## Architecture and Dependency Rules

```
cmd/push-swap  →  parser, stack, sort
cmd/checker    →  parser, stack, operations
sort           →  stack, operations
operations     →  stack
parser         →  stdlib only
stack          →  stdlib only
```

- `internal/stack` must not import any other internal package
- `internal/parser` must not import any other internal package
- `internal/operations` must not import `sort` or `parser`
- No circular imports

---

## Critical Constraints

| Input size | Max operations |
|---|---|
| n = 5 | < 12 |
| n = 6 | < 9 |
| n = 100 | < 700 |

These are hard limits enforced by the school spec and tested in CI.

---

## Error Handling

The only valid pattern for user-facing errors:

```go
fmt.Fprintln(os.Stderr, "Error")
os.Exit(1)
```

- Output must go to `os.Stderr`, never `os.Stdout`
- The message must be exactly `"Error"` — no extra detail
- Never use `log.Fatal`, `panic`, or `fmt.Println` for errors

---

## Allowed Packages

Only the Go standard library. `go.mod` must never contain a `require` block. No external dependencies.

---

## Code Style

- All files must be formatted with `gofmt` — CI enforces this
- All exported types and functions must have godoc comments
- Operation functions are named with a capital first letter: `Sa`, `Pb`, `Rra`, etc.
- Combined operations delegate to their constituents — never duplicate logic:
  ```go
  func Ss(a, b *stack.Stack)  { Sa(a); Sb(b) }
  func Rr(a, b *stack.Stack)  { Ra(a); Rb(b) }
  func Rrr(a, b *stack.Stack) { Rra(a); Rrb(b) }
  ```

---

## Commit Format

All commits must follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>
```

**Types:** `feat`, `fix`, `test`, `refactor`, `perf`, `docs`, `chore`, `ci`

**Scopes:** `stack`, `ops`, `sort`, `parser`, `push-swap`, `checker`, `ci`

Examples:
```
feat(stack): add IsSorted method
fix(sort): correct rotation for minIdx==2 in 4-element stack
test(ops): add table-driven tests for all 11 operations
```

---

## Coverage Targets

| Package | Target |
|---|---|
| `internal/stack` | ≥ 95% |
| `internal/operations` | ≥ 95% |
| `internal/parser` | ≥ 90% |
| `internal/sort` | ≥ 80% |
| Overall | ≥ 85% |

---

## CI Workflows

| Workflow | Trigger | What it does |
|---|---|---|
| `ci.yml` | Every PR and push to `main` | gofmt, go vet, build, test with race detector, coverage |
| `pr-title-check.yml` | Every PR | Validates Conventional Commits format on PR title |
| `lint.yml` | Every PR and push to `main` | golangci-lint |
| `coverage-gate.yml` | Every PR and push to `main` | Fails if overall coverage drops below 85% |
| `e2e.yml` | Every PR and push to `main` | Runs both binaries, checks op counts, verifies checker output |
| `label.yml` | Every PR | Auto-labels by changed files |
| `changelog.yml` | On version tag | Updates CHANGELOG.md |
| `release.yml` | On version tag | Builds and attaches binaries to GitHub Release |
