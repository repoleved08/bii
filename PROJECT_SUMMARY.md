# bii Project Summary

## Project Information

**Name**: bii (Binary Installation Interface)  
**Author**: Norman Bii (geekbii08@gmail.com)  
**GitHub**: https://github.com/repoleved08/bii  
**License**: MIT  
**Language**: Go 1.21+

## Project Description

`bii` is a command-line tool that simplifies the installation of binary tools distributed as archives (ZIP, TAR, etc.). It automatically detects executables, extracts them, and configures your system PATH - eliminating the manual steps typically required when installing downloaded tools.

## Problem Solved

When downloading tools like Go, Node.js, Hugo, kubectl, etc., users typically have to:
1. Extract the archive
2. Navigate through directory structures to find binaries
3. Copy binaries to a directory in PATH (e.g., /usr/local/bin, ~/.local/bin)
4. Manually update shell configuration files
5. Set correct permissions

`bii` automates all of these steps.

## Key Features

1. **Multi-Format Support**: Handles .zip, .tar, .tar.gz, .tgz
2. **Smart Detection**: Finds executables automatically (bin/ directories, executable permissions)
3. **Safe Extraction**: Only extracts binaries, not entire directory structures
4. **Shell Integration**: Detects bash/zsh/fish and updates configuration
5. **Flexible Installation**: User-local or custom directories
6. **Validation**: Inspect archives before installing
7. **Cross-Platform**: Builds for Linux, macOS, Windows

## Architecture

```
bii/
├── cmd/                    # CLI commands (Cobra framework)
│   └── root.go            # Main commands: install, inspect, version
├── pkg/
│   ├── archive/           # Archive handling
│   │   ├── archive.go     # Detection and extraction logic
│   │   └── archive_test.go
│   ├── installer/         # Installation coordinator
│   │   └── installer.go
│   └── shell/             # Shell detection & PATH management
│       ├── shell.go
│       └── shell_test.go
└── main.go                # Entry point
```

## Core Components

### 1. Archive Package
- Detects executables in ZIP/TAR archives
- Extracts specific files to destination
- Supports multiple archive formats

### 2. Installer Package
- Coordinates the installation process
- Validates inputs and destinations

### 3. Shell Package
- Detects user's shell (bash/zsh/fish)
- Checks if directory is in PATH
- Updates shell configuration files

### 4. CLI (Cobra)
- `bii install <archive>` - Install binaries
- `bii inspect <archive>` - Preview contents
- `bii version` - Version information

## Commands

### Install
```bash
bii install [flags] <archive>

Flags:
  -d, --dest string    Destination directory (default: ~/.local/bin)
  -s, --skip-path      Skip PATH configuration
  -y, --yes            Skip confirmation prompts
```

### Inspect
```bash
bii inspect <archive>

Shows all detected executables without installing.
```

### Version
```bash
bii version

Shows version and author information.
```

## Testing

- **Unit tests**: Complete coverage for archive and shell packages
- **Integration tests**: Test data with sample archives
- **CI/CD**: GitHub Actions for automated testing

Test coverage:
- Archive detection (ZIP, TAR, TAR.GZ)
- Shell detection and PATH management
- Extraction and installation

## Build & Release

### Local Build
```bash
make build          # Build for current platform
make install        # Build and install to ~/.local/bin
make build-all      # Build for multiple platforms
make test           # Run tests
```

### CI/CD
- **GitHub Actions CI**: Runs on every push/PR
  - Build verification
  - Run tests
  - Run linter

- **GitHub Actions Release**: Triggered by version tags
  - Builds for Linux (amd64, arm64)
  - Builds for macOS (amd64, arm64)
  - Builds for Windows (amd64)
  - Creates GitHub release with binaries

## Dependencies

- **spf13/cobra**: CLI framework
- Go standard library (archive/zip, archive/tar, compress/gzip)

No runtime dependencies - single static binary.

## Documentation

1. **README.md**: Overview and features
2. **QUICKSTART.md**: Quick installation and usage
3. **INSTALL.md**: Detailed installation guide
4. **EXAMPLES.md**: Real-world usage examples
5. **CONTRIBUTING.md**: Contribution guidelines

## Usage Examples

### Basic Usage
```bash
# Download and inspect
bii inspect go1.21.5.linux-amd64.tar.gz

# Install
bii install go1.21.5.linux-amd64.tar.gz
```

### Advanced Usage
```bash
# Custom destination
bii install --dest /opt/tools hugo.tar.gz

# Automated (no prompts)
bii install --yes terraform.zip

# Skip PATH configuration
bii install --skip-path kubectl.tar.gz
```

## Distribution

### GitHub Releases
Binaries available for:
- Linux: amd64, arm64
- macOS: amd64 (Intel), arm64 (Apple Silicon)
- Windows: amd64

### Installation Methods
1. Download binary from releases
2. Build from source with Go
3. Use `go install` command

## Future Enhancements (Ideas)

1. Support for more archive formats (.7z, .rar)
2. Checksum verification
3. Update mechanism (bii update <tool>)
4. List installed tools
5. Uninstall command
6. Configuration file for defaults
7. Homebrew/AUR packages
8. Download archives directly from URLs

## Project Status

✅ **Completed**:
- Core functionality
- Archive detection (ZIP, TAR, TAR.GZ)
- Shell integration (bash, zsh, fish)
- PATH management
- Tests and documentation
- CI/CD setup

🚀 **Ready to Use**: The tool is fully functional and ready for production use.

## How to Use This Project

### For Users
1. Download the binary from releases
2. Install to ~/.local/bin
3. Use `bii install <archive>` to install tools

### For Contributors
1. Clone the repository
2. Review CONTRIBUTING.md
3. Make changes and add tests
4. Submit pull request

### For Developers Learning Go
This project demonstrates:
- CLI development with Cobra
- Working with archives (ZIP, TAR)
- File system operations
- Testing in Go
- GitHub Actions CI/CD
- Project structure and organization

## Contact

**Author**: Norman Bii  
**Email**: geekbii08@gmail.com  
**GitHub**: https://github.com/repoleved08

## License

MIT License - Free to use, modify, and distribute.

---

**Project Created**: January 2026  
**Status**: Active Development  
**Version**: 0.1.0
