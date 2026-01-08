package shell

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DetectShell detects the current user's shell
func DetectShell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", fmt.Errorf("SHELL environment variable not set")
	}
	
	return filepath.Base(shell), nil
}

// IsInPath checks if a directory is in the current PATH
func IsInPath(dir string) (bool, error) {
	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")
	
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return false, err
	}
	
	for _, p := range dirs {
		absPath, err := filepath.Abs(p)
		if err != nil {
			continue
		}
		if absPath == absDir {
			return true, nil
		}
	}
	
	return false, nil
}

// AddToPath adds a directory to the PATH in the appropriate shell config file
func AddToPath(shell, dir string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	
	var configFile string
	var exportLine string
	
	switch shell {
	case "bash":
		configFile = filepath.Join(home, ".bashrc")
		exportLine = fmt.Sprintf("\n# Added by bii\nexport PATH=\"%s:$PATH\"\n", dir)
	case "zsh":
		configFile = filepath.Join(home, ".zshrc")
		exportLine = fmt.Sprintf("\n# Added by bii\nexport PATH=\"%s:$PATH\"\n", dir)
	case "fish":
		configFile = filepath.Join(home, ".config", "fish", "config.fish")
		exportLine = fmt.Sprintf("\n# Added by bii\nset -gx PATH %s $PATH\n", dir)
		// Ensure fish config directory exists
		if err := os.MkdirAll(filepath.Dir(configFile), 0755); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}
	
	// Check if already added
	content, err := os.ReadFile(configFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	
	if strings.Contains(string(content), dir) {
		return nil // Already added
	}
	
	// Append to config file
	f, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	
	if _, err := f.WriteString(exportLine); err != nil {
		return err
	}
	
	return nil
}

// GetShellConfigPath returns the config file path for the given shell
func GetShellConfigPath(shell string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	
	switch shell {
	case "bash":
		return filepath.Join(home, ".bashrc"), nil
	case "zsh":
		return filepath.Join(home, ".zshrc"), nil
	case "fish":
		return filepath.Join(home, ".config", "fish", "config.fish"), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}

// TestCommand checks if a command is available in PATH
func TestCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
