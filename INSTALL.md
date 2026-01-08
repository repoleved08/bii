# Installation Guide

## Prerequisites

- Linux system (tested on Ubuntu, Arch, Fedora)
- Go 1.19+ (only for building from source)

## Installation Methods

### Method 1: Download Pre-built Binary (Recommended)

```bash
# Download the latest release for your platform
wget https://github.com/repoleved08/bii/releases/latest/download/bii-linux-amd64

# Make it executable
chmod +x bii-linux-amd64

# Move to a directory in your PATH
sudo mv bii-linux-amd64 /usr/local/bin/bii

# Or install to user directory (no sudo needed)
mkdir -p ~/.local/bin
mv bii-linux-amd64 ~/.local/bin/bii
```

### Method 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/repoleved08/bii.git
cd bii

# Build and install
make install

# Or just build
make build
```

### Method 3: Using Go Install

```bash
go install github.com/repoleved08/bii@latest
```

## Quick Start

After installation, test bii:

```bash
bii version
```

## Examples

### Install a tool from a tar.gz archive

```bash
bii install golang.tar.gz
```

### Install to a custom location

```bash
bii install --dest /opt/mytools mytool.zip
```

### Inspect an archive before installing

```bash
bii inspect package.tar.gz
```

### Skip PATH configuration

```bash
bii install --skip-path tool.zip
```

### Non-interactive installation

```bash
bii install --yes archive.tar.gz
```

## Verify Installation

After installing a tool with bii:

1. Restart your shell or run: `source ~/.bashrc` (or `.zshrc`, `.config/fish/config.fish`)
2. Verify the tool is in PATH: `which <tool-name>`
3. Run the tool to test it works

## Troubleshooting

### PATH not updated

If bii doesn't automatically update your PATH:

**Bash/Zsh:**
```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

**Fish:**
```fish
set -Ux fish_user_paths $HOME/.local/bin $fish_user_paths
```

### Permission denied when installing

If you get permission errors:
- Use `--dest ~/.local/bin` instead of system directories
- Make sure the destination directory exists: `mkdir -p ~/.local/bin`

### Binary not detected in archive

Bii looks for:
- Files in `bin/` directories
- Files with executable permissions
- Excludes: `.sh`, `.py`, `.rb`, `.pl` scripts

If your archive has a different structure, you may need to manually extract and install.

## Supported Archive Formats

- `.zip`
- `.tar`
- `.tar.gz` / `.tgz`

## Supported Shells

- Bash
- Zsh
- Fish

## Uninstalling bii

```bash
rm ~/.local/bin/bii
# Also remove the PATH entry from your shell config file
```
