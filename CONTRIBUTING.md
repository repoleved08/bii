# Contributing to bii

Thank you for considering contributing to bii! This document outlines the process and guidelines.

## How to Contribute

### Reporting Bugs

If you find a bug, please create an issue with:
- Clear description of the bug
- Steps to reproduce
- Expected behavior
- Actual behavior
- System information (OS, shell, Go version)
- Archive format and structure if relevant

### Suggesting Features

Feature suggestions are welcome! Please create an issue with:
- Clear description of the feature
- Use case and motivation
- Example usage

### Pull Requests

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes**
4. **Test thoroughly**
5. **Commit with clear messages**: `git commit -m "Add amazing feature"`
6. **Push to your fork**: `git push origin feature/amazing-feature`
7. **Open a Pull Request**

## Development Setup

### Prerequisites

- Go 1.19 or later
- Git
- Make (optional, but helpful)

### Building

```bash
# Clone the repository
git clone https://github.com/repoleved08/bii.git
cd bii

# Install dependencies
go mod download

# Build
go build -o bii .

# Or use Make
make build
```

### Testing

```bash
# Run tests
go test -v ./...

# Or use Make
make test
```

### Code Style

- Follow standard Go formatting: `go fmt ./...`
- Run linter: `go vet ./...`
- Keep functions focused and small
- Add comments for exported functions
- Use meaningful variable names

## Project Structure

```
bii/
├── cmd/              # CLI commands
│   └── root.go       # Root command and subcommands
├── pkg/              # Core packages
│   ├── archive/      # Archive handling (zip, tar)
│   ├── installer/    # Installation logic
│   └── shell/        # Shell detection and PATH management
├── main.go           # Entry point
├── go.mod            # Go module file
└── README.md         # Documentation
```

## Adding Support for New Features

### New Archive Format

1. Add detection in `pkg/archive/archive.go`
2. Implement extract function
3. Add tests
4. Update documentation

### New Shell Support

1. Add detection in `pkg/shell/shell.go`
2. Implement PATH update logic
3. Add tests
4. Update documentation

## Testing Your Changes

Before submitting:

1. **Build successfully**: `go build`
2. **Run tests**: `go test ./...`
3. **Test manually** with real archives
4. **Check different shells** (bash, zsh, fish)
5. **Verify documentation** is updated

## Code Review Process

1. Maintainers will review your PR
2. Address any feedback
3. Once approved, maintainers will merge

## Questions?

Feel free to open an issue for any questions!

## Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow
- Have fun!

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to bii! 🚀
