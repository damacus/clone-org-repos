# Repository Guidelines

## Project Structure & Module Organization
- Entry point: `main.go`, which delegates CLI execution to `cmd/root.go`.
- CLI wiring: `cmd/` contains Cobra command setup and flag parsing.
- Core logic: `checkout/checkout.go` handles GitHub GraphQL queries and clone/pull behavior.
- Tests: keep `_test.go` files next to the code they validate (for example, `checkout/checkout_test.go`).
- Repo metadata: `.github/workflows/` for automation, `CHANGELOG.md` managed by release tooling, `renovate.json` for dependency updates.

## Build, Test, and Development Commands
- `go mod tidy`: sync module dependencies before committing.
- `go build ./...`: compile all packages and catch build-time issues.
- `go build -o clone-org-repos .`: produce the local CLI binary.
- `go test ./...`: run all tests.
- `go test -cover ./...`: run tests with coverage output.
- `go run . -o <org> -p <path>`: run locally without building a binary.

Example:
```bash
go run . -o sous-chefs -p ~/dev/sous-chefs
```

## Coding Style & Naming Conventions
- Follow standard Go formatting: run `gofmt -w .` before opening a PR.
- Prefer short, focused functions and package-local helpers for shared behavior.
- Use Go naming conventions: exported identifiers in `CamelCase`, unexported in `camelCase`.
- Keep package names lowercase and concise (`checkout`, `cmd`).
- Maintain existing CLI flag style: long flag with short alias (`--org`, `-o`).

## Testing Guidelines
- Use Go’s built-in `testing` package.
- Name tests as `Test<FunctionOrBehavior>` (for example, `TestCheckout_InvalidToken`).
- Favor table-driven tests for multiple input/output cases.
- Add tests for new logic and error paths in `checkout/` and any new packages.

## Commit & Pull Request Guidelines
- Use Conventional Commit style where possible (`fix:`, `feat:`, `chore:`); this aligns with release-please changelog generation.
- Keep commits scoped and atomic; avoid mixing refactors with behavior changes.
- PRs should include:
  - a brief problem/solution summary,
  - linked issue(s) when applicable,
  - test evidence (`go test ./...` output or coverage notes),
  - CLI output examples for behavior changes.

## Security & Configuration Tips
- `GITHUB_TOKEN` is required at runtime; never hardcode or commit tokens.
- Prefer environment variables for credentials and keep local overrides out of version control.
