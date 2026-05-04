## Summary

<!-- What changed and why. Link related issues. -->

## Type of change

- [ ] Bug fix
- [ ] New feature / enum
- [ ] Refactor (no behavior change)
- [ ] Docs / spec update
- [ ] CI / tooling

## Checklist

- [ ] `./run.sh tc` (or `./run.ps1 TC`) passes locally with ≥ 60% coverage
- [ ] `golangci-lint run --timeout=5m` is clean
- [ ] `go vet ./...` is clean
- [ ] `CHANGELOG.md` updated under `## [Unreleased]` if user-visible
- [ ] No `@latest` / `@main` pins added in workflows or scripts
