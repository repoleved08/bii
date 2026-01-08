# Examples

## Basic Usage

### 1. Inspect an archive

Before installing, you can inspect what binaries are in an archive:

```bash
$ bii inspect go1.21.5.linux-amd64.tar.gz
📦 Inspecting: go1.21.5.linux-amd64.tar.gz

✅ Found 3 executable(s):
  • go/bin/go
  • go/bin/gofmt
  • go/bin/godoc
```

### 2. Install from archive

Install all detected binaries to `~/.local/bin` (default):

```bash
$ bii install go1.21.5.linux-amd64.tar.gz
📦 Installing from: go1.21.5.linux-amd64.tar.gz
📁 Destination: /home/user/.local/bin

✅ Found 3 executable(s):
  • go/bin/go
  • go/bin/gofmt
  • go/bin/godoc

Continue with installation? [Y/n]: y

✅ Successfully installed 3 binary(ies) to /home/user/.local/bin
  • go
  • gofmt
  • godoc

🐚 Detected shell: bash
📝 Adding /home/user/.local/bin to PATH...
✅ PATH updated in shell configuration
💡 Restart your shell or run: source ~/.bashrc
```

### 3. Custom installation directory

Install to `/opt`:

```bash
$ sudo bii install --dest /opt/golang go1.21.5.linux-amd64.tar.gz
```

### 4. Skip PATH configuration

If you want to handle PATH yourself:

```bash
$ bii install --skip-path mytool.zip
```

### 5. Non-interactive installation

For scripts or automation:

```bash
$ bii install --yes terraform.zip
```

## Real-world Examples

### Installing Go

```bash
# Download Go
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz

# Install with bii
bii install go1.21.5.linux-amd64.tar.gz

# Verify
go version
```

### Installing kubectl

```bash
# Download kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

# Create archive (if not already archived)
mkdir -p kubectl-bin/bin
mv kubectl kubectl-bin/bin/
tar -czf kubectl.tar.gz kubectl-bin/

# Install with bii
bii install kubectl.tar.gz

# Verify
kubectl version --client
```

### Installing Hugo

```bash
# Download Hugo
wget https://github.com/gohugoio/hugo/releases/download/v0.120.0/hugo_0.120.0_linux-amd64.tar.gz

# Install with bii
bii install hugo_0.120.0_linux-amd64.tar.gz

# Verify
hugo version
```

### Installing to system-wide location

```bash
# Install to /usr/local/bin (requires sudo)
sudo bii install --dest /usr/local/bin --skip-path mytool.zip
```

## Advanced Usage

### Inspecting multiple archives

```bash
for archive in *.tar.gz; do
    echo "Checking $archive..."
    bii inspect "$archive"
    echo ""
done
```

### Batch installation

```bash
#!/bin/bash
archives=(
    "tool1.tar.gz"
    "tool2.zip"
    "tool3.tgz"
)

for archive in "${archives[@]}"; do
    bii install --yes "$archive"
done
```

### Installation script for new systems

```bash
#!/bin/bash
# setup-dev-tools.sh

set -e

# Install bii first
echo "Installing bii..."
wget https://github.com/repoleved08/bii/releases/latest/download/bii-linux-amd64
chmod +x bii-linux-amd64
mkdir -p ~/.local/bin
mv bii-linux-amd64 ~/.local/bin/bii

# Add to PATH
export PATH="$HOME/.local/bin:$PATH"

# Install tools
echo "Installing Go..."
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
bii install --yes go1.21.5.linux-amd64.tar.gz

echo "Installing Node.js..."
wget https://nodejs.org/dist/v20.10.0/node-v20.10.0-linux-x64.tar.gz
bii install --yes node-v20.10.0-linux-x64.tar.gz

echo "All tools installed! Restart your shell."
```

## Shell-specific Examples

### Bash

```bash
# After installation, reload config
source ~/.bashrc

# Check if bii is available
which bii

# Check installed tools
ls -l ~/.local/bin/
```

### Zsh

```zsh
# After installation, reload config
source ~/.zshrc

# Check if bii is available
which bii
```

### Fish

```fish
# After installation, reload config
source ~/.config/fish/config.fish

# Check if bii is available
which bii

# View PATH
echo $PATH
```

## Tips and Tricks

### 1. Check what will be installed first

Always run `inspect` before `install` to see what binaries will be extracted:

```bash
bii inspect unknown-tool.tar.gz
```

### 2. Create your own tool archives

If you have a binary you want to distribute:

```bash
mkdir -p mytool/bin
cp mytool mytool/bin/
tar -czf mytool.tar.gz mytool/
```

Now others can install with `bii install mytool.tar.gz`

### 3. Combine with version managers

Use bii for tools that don't have version managers:

```bash
# Tools with version managers - use those
# Node.js -> nvm
# Python -> pyenv
# Ruby -> rbenv

# Standalone tools - use bii
bii install terraform.zip
bii install hugo.tar.gz
bii install kubectl.tar.gz
```

### 4. Local bin directory structure

Keep your `~/.local/bin` organized:

```bash
~/.local/bin/
├── bii
├── go
├── gofmt
├── hugo
├── kubectl
└── terraform
```
