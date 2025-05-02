# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands
- Build: `go build`
- Run: `./workdays -start YYYY-MM-DD -end YYYY-MM-DD`
- Test: `go test ./...`
- Run specific test: `go test -v -run TestName`
- Format code: `gofmt -w .`
- Lint with custom linter: `go run honnef.co/go/tools/cmd/staticcheck ./...`

## Nix Environment
- Activate: `nix develop` or `direnv allow` if using direnv

## Style Guidelines
- Follow standard Go style conventions
- Group imports: standard lib, then third-party libs with blank line separating
- Error handling: check errors and return with context where appropriate
- Variable naming: camelCase, descriptive but concise names
- Indentation: use tabs for indentation
- Function organization: smaller, focused functions preferred
- Comments: add comments for public APIs
- Return early: handle error cases at beginning of functions
- Prefer map[string]bool over []string for lookups

## Dependencies
- charmbracelet/lipgloss: terminal styling
- rickar/cal/v2: calendar and workday calculations