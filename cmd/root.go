package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/repoleved08/bii/pkg/archive"
	"github.com/repoleved08/bii/pkg/installer"
	"github.com/repoleved08/bii/pkg/shell"
	"github.com/spf13/cobra"
)

var (
	destDir    string
	skipPath   bool
	forceYes   bool
	rootCmd    = &cobra.Command{
		Use:   "bii",
		Short: "Binary Installation Interface - Install binaries from archives",
		Long:  `bii helps you install binary tools from ZIP and TAR archives with automatic PATH management.`,
	}
)

func init() {
	installCmd.Flags().StringVarP(&destDir, "dest", "d", "", "Destination directory (default: ~/.local/bin)")
	installCmd.Flags().BoolVarP(&skipPath, "skip-path", "s", false, "Skip PATH configuration")
	installCmd.Flags().BoolVarP(&forceYes, "yes", "y", false, "Skip confirmation prompts")
	
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(inspectCmd)
	rootCmd.AddCommand(versionCmd)
}

var installCmd = &cobra.Command{
	Use:   "install <archive>",
	Short: "Install binaries from an archive",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstall,
}

var inspectCmd = &cobra.Command{
	Use:   "inspect <archive>",
	Short: "Inspect an archive and show detected binaries",
	Args:  cobra.ExactArgs(1),
	RunE:  runInspect,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bii v0.1.0")
		fmt.Println("Author: Norman Bii <geekbii08@gmail.com>")
	},
}

func runInspect(cmd *cobra.Command, args []string) error {
	archivePath := args[0]
	
	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return fmt.Errorf("archive not found: %s", archivePath)
	}
	
	fmt.Printf("📦 Inspecting: %s\n\n", archivePath)
	
	binaries, err := archive.DetectBinaries(archivePath)
	if err != nil {
		return fmt.Errorf("failed to inspect archive: %w", err)
	}
	
	if len(binaries) == 0 {
		fmt.Println("❌ No executable binaries detected in archive")
		return nil
	}
	
	fmt.Printf("✅ Found %d executable(s):\n", len(binaries))
	for _, bin := range binaries {
		fmt.Printf("  • %s\n", bin)
	}
	
	return nil
}

func runInstall(cmd *cobra.Command, args []string) error {
	archivePath := args[0]
	
	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return fmt.Errorf("archive not found: %s", archivePath)
	}
	
	// Set default destination
	if destDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		destDir = filepath.Join(home, ".local", "bin")
	}
	
	// Ensure destination exists
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}
	
	fmt.Printf("📦 Installing from: %s\n", archivePath)
	fmt.Printf("📁 Destination: %s\n\n", destDir)
	
	// Detect binaries
	binaries, err := archive.DetectBinaries(archivePath)
	if err != nil {
		return fmt.Errorf("failed to inspect archive: %w", err)
	}
	
	if len(binaries) == 0 {
		return fmt.Errorf("no executable binaries found in archive")
	}
	
	fmt.Printf("✅ Found %d executable(s):\n", len(binaries))
	for _, bin := range binaries {
		fmt.Printf("  • %s\n", bin)
	}
	fmt.Println()
	
	// Confirm installation
	if !forceYes {
		fmt.Print("Continue with installation? [Y/n]: ")
		var response string
		fmt.Scanln(&response)
		if response != "" && response != "Y" && response != "y" {
			fmt.Println("Installation cancelled")
			return nil
		}
	}
	
	// Install binaries
	installed, err := installer.Install(archivePath, destDir, binaries)
	if err != nil {
		return fmt.Errorf("installation failed: %w", err)
	}
	
	fmt.Printf("\n✅ Successfully installed %d binary(ies) to %s\n", len(installed), destDir)
	for _, bin := range installed {
		fmt.Printf("  • %s\n", filepath.Base(bin))
	}
	
	// Handle PATH configuration
	if !skipPath {
		fmt.Println()
		if err := configurePath(destDir); err != nil {
			fmt.Fprintf(os.Stderr, "⚠️  Warning: PATH configuration failed: %v\n", err)
			fmt.Printf("Please manually add to PATH: export PATH=\"%s:$PATH\"\n", destDir)
		}
	}
	
	return nil
}

func configurePath(dir string) error {
	currentShell, err := shell.DetectShell()
	if err != nil {
		return err
	}
	
	fmt.Printf("🐚 Detected shell: %s\n", currentShell)
	
	inPath, err := shell.IsInPath(dir)
	if err != nil {
		return err
	}
	
	if inPath {
		fmt.Printf("✅ %s is already in PATH\n", dir)
		return nil
	}
	
	fmt.Printf("📝 Adding %s to PATH...\n", dir)
	
	if err := shell.AddToPath(currentShell, dir); err != nil {
		return err
	}
	
	fmt.Printf("✅ PATH updated in shell configuration\n")
	fmt.Println("💡 Restart your shell or run: source ~/.<shell>rc")
	
	return nil
}

func Execute() error {
	return rootCmd.Execute()
}
