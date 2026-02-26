# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [Unreleased]

### Added
- Stack data structure (`NewStack`, `Push`, `Pop`, `Peek`, `Len`, `IsEmpty`, `Values`) with defensive copies on input and output
- All 11 push-swap operations: `sa`, `sb`, `ss`, `pa`, `pb`, `ra`, `rb`, `rr`, `rra`, `rrb`, `rrr` — each a no-op on 0 or 1 element stacks
- Small stack sort algorithm for n=2–5 elements
- `IsSorted` helper for validating sort correctness in tests
- `cmd/push-swap/main.go` and `cmd/checker/main.go` entry point stubs
- CI pipeline: gofmt, go vet, build, test with race detector, coverage report
- AI-powered PR review via `claude.yml` workflow
- Project documentation: `AGENTS.md`, `PERMISSIONS.md`, `CONTRIBUTING.md`, `Makefile`, `README.md` and `CHANGELOG.md`

---

<!-- Add releases below as they are tagged -->
<!-- Format:
## [1.0.0] - YYYY-MM-DD

### Added
### Fixed
### Changed
### Removed
-->
