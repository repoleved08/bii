# bii - Binary Installation Interface

[![CI](https://github.com/repoleved08/bii/workflows/CI/badge.svg)](https://github.com/repoleved08/bii/actions)
[![Release](https://img.shields.io/github/v/release/repoleved08/bii)](https://github.com/repoleved08/bii/releases)
[![License](https://img.shields.io/github/license/repoleved08/bii)](LICENSE)

> A simple, powerful tool to install binaries from archives on Linux/macOS systems

`bii` helps you quickly install downloaded tools that come as ZIP or TAR archives. No more manual extraction, copying to bin directories, or PATH configuration headaches!

## ✨ Features

- 📦 **Multi-format Support**: `.zip`, `.tar`, `.tar.gz`, `.tgz`
- 🔍 **Smart Detection**: Automatically finds executables in archives
- 🎯 **Flexible Installation**: User-local (`~/.local/bin`) or custom directories
- 🐚 **Shell-aware**: Detects and configures bash, zsh, and fish
- ⚡ **Fast & Safe**: Written in Go, single binary, no dependencies
- 🔗 **PATH Management**: Automatically updates your shell configuration
- ✅ **Validation**: Inspect archives before installing

## 🚀 Quick Start

### Installation

**Linux (amd64):**
```bash
# Download the latest release (v0.1.0)
wget https://github.com/repoleved08/bii/releases/download/v0.1.0/bii-linux-amd64
chmod +x bii-linux-amd64
mkdir -p ~/.local/bin
mv bii-linux-amd64 ~/.local/bin/bii

# Add to PATH (if not already)
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Verify installation
bii version
```

**Linux (arm64):**
```bash
wget https://github.com/repoleved08/bii/releases/download/v0.1.0/bii-linux-arm64
chmod +x bii-linux-arm64
mkdir -p ~/.local/bin
mv bii-linux-arm64 ~/.local/bin/bii
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

**macOS (Intel):**
```bash
curl -LO https://github.com/repoleved08/bii/releases/download/v0.1.0/bii-darwin-amd64
chmod +x bii-darwin-amd64
mkdir -p ~/.local/bin
mv bii-darwin-amd64 ~/.local/bin/bii
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

**macOS (Apple Silicon):**
```bash
curl -LO https://github.com/repoleved08/bii/releases/download/v0.1.0/bii-darwin-arm64
chmod +x bii-darwin-arm64
mkdir -p ~/.local/bin
mv bii-darwin-arm64 ~/.local/bin/bii
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

**Windows (amd64):**
```powershell
# Download from: https://github.com/repoleved08/bii/releases/download/v0.1.0/bii-windows-amd64.exe
# Move to a directory in your PATH or add the directory to PATH
```

**Or build from source:**
```bash
go install github.com/repoleved08/bii@v0.1.0
```

See [INSTALL.md](INSTALL.md) for more installation options and troubleshooting.

### Usage

```bash
# Inspect an archive first
bii inspect go1.21.5.linux-amd64.tar.gz

# Install binaries from archive
bii install go1.21.5.linux-amd64.tar.gz

# Install to custom location
bii install --dest /opt/mytools kubectl.tar.gz

# Non-interactive installation
bii install --yes terraform.zip

# Skip PATH configuration
bii install --skip-path hugo.tar.gz
```

## 📖 Documentation

- [Installation Guide](INSTALL.md)
- [Usage Examples](EXAMPLES.md)
- [Contributing Guidelines](CONTRIBUTING.md)

## 🎯 Use Cases

Perfect for installing:
- **Go**: `bii install go1.21.5.linux-amd64.tar.gz`
- **Node.js**: `bii install node-v20.10.0-linux-x64.tar.gz`
- **Hugo**: `bii install hugo_extended_0.120.0_linux-amd64.tar.gz`
- **kubectl**: `bii install kubectl.tar.gz`
- **Terraform**: `bii install terraform_1.6.0_linux_amd64.zip`
- Any tool distributed as a ZIP/TAR archive!

## 🛠️ How It Works

1. **Detection**: Scans archive for executables in `bin/` directories or with executable permissions
2. **Extraction**: Extracts only the binaries (not entire directory structures)
3. **Installation**: Copies to destination (default: `~/.local/bin`)
4. **PATH Setup**: Updates your shell config to include the installation directory

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📝 License

MIT License - see [LICENSE](LICENSE) file for details

## 👤 Author

**Norman Bii**
- Email: geekbii08@gmail.com
- GitHub: [@repoleved08](https://github.com/repoleved08)

## 🌟 Show Your Support

Give a ⭐️ if this project helped you!

## 📊 Project Status

This project is actively maintained. Issues and feature requests are welcome!

---

**Note**: `bii` itself is distributed as a single binary - once you have it installed, you can use it to install other tools! Consider it your "binary package manager" for standalone tools.
