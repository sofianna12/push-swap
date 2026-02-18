# AGENTS.md — Claude Code Setup for Push-Swap

This file documents the AI-assisted development environment for this project.
It is read automatically at the start of every Claude Code session via `CLAUDE.md`.

---

## Agents

Agents are specialised AI sub-processes. Invoke them by saying _"use the X agent"_ in the chat.

| Agent | Model | Can Edit Files? | Purpose |
|---|---|---|---|
| `code-reviewer` | Opus | No (read-only) | Reviews PRs for correctness, style, spec compliance, and op count limits |
| `test-writer` | Sonnet | Yes | Writes table-driven Go tests following TDD. Covers all 11 operations + sort constraints |
| `refactorer` | Sonnet | Yes | Improves code structure without changing behaviour or degrading op counts |
| `documentation` | Sonnet | Yes | Generates and maintains README, CHANGELOG, CONTRIBUTING, Makefile, mermaid diagrams |
| `security-auditor` | Opus | No (read-only) | Audits input validation and error routing before audits or releases |

### When to use which agent
- **Before opening a PR**: `code-reviewer` + `security-auditor`
- **After writing a new function**: `test-writer`
- **When tests pass but code feels messy**: `refactorer`
- **Before the official zone01 audit**: `security-auditor` → `/run-audit`
- **For a new release**: `documentation` (update CHANGELOG) → `/release`

---

## Skills (Slash Commands)

Skills are custom slash commands. Type `/skill-name` in the Claude Code prompt.

| Skill | Usage | What it does |
|---|---|---|
| `/review-code` | `/review-code [file]` | Reviews a file or the latest changes for correctness and style |
| `/run-audit` | `/run-audit` | Runs every check from `OFFICIAL_AUDITS.md` against live binaries |
| `/check-coverage` | `/check-coverage` | Reports per-package test coverage vs targets |
| `/fix-issue` | `/fix-issue [description]` | Reproduce → diagnose → fix → add regression test |
| `/explain-code` | `/explain-code [file or concept]` | Explains code in plain language (great for the Turkish algorithm) |
| `/benchmark` | `/benchmark [n]` | Measures op counts for n=5 (<12), n=6 (<9), n=100 (<700) |
| `/release` | `/release [v1.0.0]` | Quality gate → build binaries → smoke test → git tag → dual push |

---

## Hooks

Git hooks are installed via `bash .claude/hooks/install-hooks.sh` (run once after cloning).

| Hook | Trigger | What it checks |
|---|---|---|
| `pre-commit` | Every `git commit` | gofmt, go vet, no external deps, no conflict markers |
| `commit-msg` | Every `git commit` | Conventional Commits format (type(scope): description) |

To bypass hooks when needed: `git commit --no-verify`

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

## CI / AI PR Review

Every PR to `main` runs two workflows:

- **`ci.yml`** — gofmt, go vet, build, `go test -race`, coverage report
- **`claude.yml`** — Claude Opus automatically reviews the PR and posts a comment with verdict

For `claude.yml` to work, the repo needs an `ANTHROPIC_API_KEY` GitHub secret.
Team members can also trigger Claude on a PR comment by writing `@claude [question]`.

---

## Memory

Claude Code has persistent memory for this project at:
```
~/.claude/projects/-home-teovaira-push-swap/memory/MEMORY.md
```

This file is on your local machine (not in git). It stores key facts Claude should remember across sessions — architecture, constraints, team structure.
