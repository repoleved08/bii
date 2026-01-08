package shell

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDetectShell(t *testing.T) {
	// Save original SHELL
	originalShell := os.Getenv("SHELL")
	defer os.Setenv("SHELL", originalShell)

	tests := []struct {
		name     string
		shell    string
		expected string
	}{
		{"Bash", "/bin/bash", "bash"},
		{"Zsh", "/usr/bin/zsh", "zsh"},
		{"Fish", "/usr/local/bin/fish", "fish"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("SHELL", tt.shell)
			result, err := DetectShell()
			if err != nil {
				t.Fatalf("DetectShell failed: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestIsInPath(t *testing.T) {
	// Save original PATH
	originalPath := os.Getenv("PATH")
	defer os.Setenv("PATH", originalPath)

	testDir := "/test/bin"
	os.Setenv("PATH", "/usr/bin:/bin:"+testDir+":/usr/local/bin")

	inPath, err := IsInPath(testDir)
	if err != nil {
		t.Fatalf("IsInPath failed: %v", err)
	}

	if !inPath {
		t.Error("Expected directory to be in PATH")
	}

	// Test directory not in PATH
	inPath, err = IsInPath("/not/in/path")
	if err != nil {
		t.Fatalf("IsInPath failed: %v", err)
	}

	if inPath {
		t.Error("Expected directory to not be in PATH")
	}
}

func TestGetShellConfigPath(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get home directory")
	}

	tests := []struct {
		shell    string
		expected string
	}{
		{"bash", filepath.Join(home, ".bashrc")},
		{"zsh", filepath.Join(home, ".zshrc")},
		{"fish", filepath.Join(home, ".config", "fish", "config.fish")},
	}

	for _, tt := range tests {
		t.Run(tt.shell, func(t *testing.T) {
			result, err := GetShellConfigPath(tt.shell)
			if err != nil {
				t.Fatalf("GetShellConfigPath failed: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}

	// Test unsupported shell
	_, err = GetShellConfigPath("unsupported")
	if err == nil {
		t.Error("Expected error for unsupported shell")
	}
}

func TestAddToPath(t *testing.T) {
	// Create temporary home directory
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	testDir := "/test/bin"

	tests := []struct {
		shell          string
		configFile     string
		expectedSubstr string
	}{
		{"bash", ".bashrc", "export PATH=\"" + testDir + ":$PATH\""},
		{"zsh", ".zshrc", "export PATH=\"" + testDir + ":$PATH\""},
		{"fish", ".config/fish/config.fish", "set -gx PATH " + testDir + " $PATH"},
	}

	for _, tt := range tests {
		t.Run(tt.shell, func(t *testing.T) {
			// Create a fresh temp dir for each test
			shellTmpDir := t.TempDir()
			os.Setenv("HOME", shellTmpDir)

			err := AddToPath(tt.shell, testDir)
			if err != nil {
				t.Fatalf("AddToPath failed: %v", err)
			}

			// Verify config file was created and contains the PATH
			configPath := filepath.Join(shellTmpDir, tt.configFile)
			content, err := os.ReadFile(configPath)
			if err != nil {
				t.Fatalf("Failed to read config file: %v", err)
			}

			if !strings.Contains(string(content), tt.expectedSubstr) {
				t.Errorf("Config file does not contain expected PATH export.\nExpected substring: %s\nGot: %s",
					tt.expectedSubstr, string(content))
			}

			// Test that running again doesn't duplicate
			err = AddToPath(tt.shell, testDir)
			if err != nil {
				t.Fatalf("Second AddToPath failed: %v", err)
			}

			content2, err := os.ReadFile(configPath)
			if err != nil {
				t.Fatal(err)
			}

			// Count occurrences of testDir
			count := strings.Count(string(content2), testDir)
			if count != 1 {
				t.Errorf("Expected testDir to appear once, but appeared %d times", count)
			}
		})
	}
}
