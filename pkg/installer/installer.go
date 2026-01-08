package installer

import (
	"fmt"
	"github.com/repoleved08/bii/pkg/archive"
)

// Install extracts and installs binaries from an archive to the destination directory
func Install(archivePath, destDir string, binaries []string) ([]string, error) {
	if len(binaries) == 0 {
		return nil, fmt.Errorf("no binaries to install")
	}
	
	installed, err := archive.Extract(archivePath, destDir, binaries)
	if err != nil {
		return installed, fmt.Errorf("extraction failed: %w", err)
	}
	
	return installed, nil
}
