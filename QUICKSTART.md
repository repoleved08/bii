# Quick Start Guide for bii

## What is bii?

`bii` (Binary Installation Interface) is a command-line tool that simplifies installing downloaded binaries from ZIP and TAR archives. It automatically detects executables, extracts them to the right location, and configures your PATH.

## Installation

### Quick Install (Linux)

```bash
# Download the binary
wget https://github.com/repoleved08/bii/releases/latest/download/bii-linux-amd64

# Make it executable
chmod +x bii-linux-amd64

# Move to your local bin (no sudo needed)
mkdir -p ~/.local/bin
mv bii-linux-amd64 ~/.local/bin/bii

# Add to PATH (if not already there)
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Verify installation
bii version
```

### Build from Source

```bash
git clone https://github.com/repoleved08/bii.git
cd bii
make install
```

## Common Usage

### 1. Inspect before installing

```bash
bii inspect go1.21.5.linux-amd64.tar.gz
```

**Output:**
```
📦 Inspecting: go1.21.5.linux-amd64.tar.gz

✅ Found 3 executable(s):
  • go/bin/go
  • go/bin/gofmt
  • go/bin/godoc
```

### 2. Install binaries

```bash
bii install go1.21.5.linux-amd64.tar.gz
```

**What happens:**
- ✅ Extracts binaries to `~/.local/bin`
- ✅ Sets correct permissions
- ✅ Updates your shell config (if needed)
- ✅ Shows you what to do next

### 3. Custom installation

```bash
# Install to specific directory
bii install --dest /opt/mytools tool.zip

# Non-interactive (for scripts)
bii install --yes package.tar.gz

# Skip PATH configuration
bii install --skip-path archive.tgz
```

## Real-World Examples

### Installing Go

```bash
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
bii install go1.21.5.linux-amd64.tar.gz
source ~/.bashrc
go version
```

### Installing Hugo

```bash
wget https://github.com/gohugoio/hugo/releases/download/v0.120.0/hugo_0.120.0_linux-amd64.tar.gz
bii install hugo_0.120.0_linux-amd64.tar.gz
hugo version
```

### Installing kubectl

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
mkdir -p kubectl-bin/bin && mv kubectl kubectl-bin/bin/
tar -czf kubectl.tar.gz kubectl-bin/
bii install kubectl.tar.gz
kubectl version --client
```

## Supported Formats

- ✅ `.zip`
- ✅ `.tar`
- ✅ `.tar.gz`
- ✅ `.tgz`

## Supported Shells

- ✅ Bash
- ✅ Zsh
- ✅ Fish

## How It Works

1. **Scans** the archive for executables
   - Files in `bin/` directories
   - Files with executable permissions
   
2. **Extracts** only the binaries (not entire directory structures)

3. **Installs** to destination directory (default: `~/.local/bin`)

4. **Configures** your shell's PATH automatically

## Tips

💡 **Always inspect first**: Run `bii inspect` to see what will be installed

💡 **Use ~/.local/bin**: It's the standard user-local binary directory (no sudo needed)

💡 **Restart your shell**: After installation, restart your terminal or run `source ~/.bashrc`

💡 **Check installation**: Use `which <command>` to verify the binary is in your PATH

## Troubleshooting

### "command not found" after installation

```bash
# Make sure ~/.local/bin is in PATH
echo $PATH | grep .local/bin

# If not, add it:
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### Permission errors

```bash
# Use local directory instead of system directories
bii install --dest ~/.local/bin tool.zip

# Make sure the directory exists
mkdir -p ~/.local/bin
```

### Archive not recognized

Supported formats: `.zip`, `.tar`, `.tar.gz`, `.tgz`

If your file has a different extension, rename it:
```bash
mv tool.archive tool.tar.gz
bii install tool.tar.gz
```

## Getting Help

```bash
# General help
bii --help

# Command-specific help
bii install --help
bii inspect --help

# Version info
bii version
```

## Next Steps

1. ✅ Install bii
2. ✅ Download an archive (e.g., Go, Hugo, kubectl)
3. ✅ Run `bii inspect <archive>` to preview
4. ✅ Run `bii install <archive>` to install
5. ✅ Restart your shell
6. ✅ Use your newly installed tool!

## More Information

- **Full documentation**: [README.md](README.md)
- **Installation guide**: [INSTALL.md](INSTALL.md)
- **Examples**: [EXAMPLES.md](EXAMPLES.md)
- **Contributing**: [CONTRIBUTING.md](CONTRIBUTING.md)

---

**Need help?** Open an issue on [GitHub](https://github.com/repoleved08/bii/issues)
