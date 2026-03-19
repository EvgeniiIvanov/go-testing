### SonarCloud coverage status:
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=EvgeniiIvanov_go-testing&metric=coverage)](https://sonarcloud.io/summary/new_code?id=EvgeniiIvanov_go-testing)

# go-testing

A collection of Go testing examples and patterns. From basic unit tests to parallel execution, mocks, and temp files.

## What's inside

- **Unit tests** — basic tests, error/panic handling
- **Assertions** — using `testify/assert` for cleaner checks
- **Mocks** — simple dependency isolation
- **Parallel tests** — speed up test execution
- **Temp files** — working with filesystem in tests
- **More to come** — this repo grows as I explore new patterns

## Quick start

```bash
git clone https://github.com/EvgeniiIvanov/go-testing.git
cd go-testing/even
go test -v .
```