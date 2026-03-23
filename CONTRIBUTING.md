# Contributing to push-swap

## Team

| Member | Role | Can merge to `main`? |
|---|---|---|
| Theo (`@teovaira`) | Leader | Yes |
| Anna (`@sofianna12`) | Developer | No — open a PR |
| Alex (`@arigopou`) | Developer | No — open a PR |

---

## First-time setup

```bash
git clone <repo-url>
cd push-swap
go build ./...                         # verify it compiles
go test -race ./...                    # verify tests pass
```

---

## Branch naming

```
feature/<slug>   # new functionality
fix/<slug>       # bug fix
```

Examples: `feature/large-sort`, `fix/rotation-minidx2`

Always branch from the latest `develop`:

```bash
git fetch origin
git checkout -b feature/your-feature origin/develop
```

---

## Commit format (Conventional Commits)

```
<type>(<scope>): <description>
```

**Types:** `feat`, `fix`, `test`, `refactor`, `perf`, `docs`, `chore`, `ci`

**Scopes:** `stack`, `ops`, `sort`, `parser`, `push-swap`, `checker`, `ci`

```bash
# Good
feat(stack): add IsSorted method with ascending order check
fix(sort): correct rotation for minIdx==2 in 4-element stack
test(ops): add table-driven tests for all 11 operations

# Bad
fixed stuff
Update sort.go
```

Rules: imperative mood, lowercase, no trailing period.

---

## Pre-PR checklist

Run all of these before opening a PR:

```bash
gofmt -w .               # format
go vet ./...             # static analysis
go test -race -cover -v ./...      # tests with race detector, coverage summary, verbose
go build ./...           # verify compile
```

If your change touches sort logic, also verify operation counts:

```bash
go build -o push-swap ./cmd/push-swap && go build -o checker ./cmd/checker
./push-swap "1 5 2 4 3" | wc -l                        # must be < 12 (n=5)
./push-swap "2 1 3 6 5 8" | wc -l                      # must be < 9  (n=6)
ARG=$(shuf -i 1-1000 -n 100 | tr '\n' ' ')
./push-swap "$ARG" | wc -l                             # must be < 700 (n=100)
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"  # must be OK
```

---

## Opening a PR

1. Push your branch: `git push origin feature/your-feature`
2. Open a PR on GitHub: `feature/your-feature → develop`
3. PR title must follow Conventional Commits format
4. PR body must include:
   - What changed and why
   - How to test it
   - Operation counts if sort-related
5. Request a review — do **not** merge your own PR

---

## Error handling rule

The only valid pattern for user-facing errors in this project:

```go
fmt.Fprintln(os.Stderr, "Error")
os.Exit(1)
```

Never use `log.Fatal`, `panic`, or `fmt.Println` for errors. Never add extra detail to the message — the spec requires exactly `"Error"`.

---

## Allowed packages

Only the Go standard library. `go.mod` must **never** contain a `require` block. If you accidentally add one, remove it before committing — the pre-commit hook will catch it.

---

## Coverage targets

| Package | Target |
|---|---|
| `internal/stack` | ≥ 95% |
| `internal/operations` | ≥ 95% |
| `internal/parser` | ≥ 90% |
| `internal/sort` | ≥ 80% |
| Overall | ≥ 85% |
