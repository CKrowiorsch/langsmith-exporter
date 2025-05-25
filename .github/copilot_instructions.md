# Cilot Instruction File

## Language
- All code, comments, and documentation must be written in **English**.

## Go Best Practices
- Use idiomatic Go style (gofmt, golint, go vet).
- Organize code into packages with clear responsibilities.
- Use clear, descriptive names for variables, functions, and types.
- Handle errors explicitly; avoid silent failures.
- Use interfaces for abstractions and testing.
- Document all exported functions, types, and packages.
- Avoid global variables unless absolutely necessary.
- Use context for cancellation and timeouts in long-running operations.
- Prefer composition over inheritance.
- Keep functions small and focused.
- Use Go modules for dependency management.
- Write clear and concise README files.

## Unit Testing
- Write unit tests for all public functions and methods where possible.
- Use Go's built-in `testing` package.
- Name test files with `_test.go` suffix.
- Use table-driven tests for multiple scenarios.
- Mock dependencies using interfaces for isolated tests.
- Ensure tests are deterministic and independent.
- Run tests with `go test ./...` and ensure all pass before committing.

## Code Review
- All changes should be reviewed before merging.
- Encourage constructive feedback and knowledge sharing during reviews.

## Continuous Integration (CI)
- Set up CI to automatically run tests and linters on each commit or pull request.

## Dependency Management
- Regularly update dependencies and audit for vulnerabilities.

## Documentation
- Maintain up-to-date documentation, including usage examples and API references.

## Security
- Follow secure coding practices and validate all inputs.

## Performance
- Profile and optimize critical code paths as needed.

---

This file provides guidelines for contributing to this project. Please follow these instructions to ensure code quality and maintainability.
