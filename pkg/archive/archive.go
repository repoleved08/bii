package archive

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// DetectBinaries inspects an archive and returns paths to executable files
func DetectBinaries(archivePath string) ([]string, error) {
	ext := strings.ToLower(filepath.Ext(archivePath))
	
	switch ext {
	case ".zip":
		return detectBinariesZip(archivePath)
	case ".gz", ".tgz":
		return detectBinariesTarGz(archivePath)
	case ".tar":
		return detectBinariesTar(archivePath)
	default:
		return nil, fmt.Errorf("unsupported archive format: %s", ext)
	}
}

// Extract extracts specific files from an archive to destination
func Extract(archivePath, destDir string, files []string) ([]string, error) {
	ext := strings.ToLower(filepath.Ext(archivePath))
	
	switch ext {
	case ".zip":
		return extractZip(archivePath, destDir, files)
	case ".gz", ".tgz":
		return extractTarGz(archivePath, destDir, files)
	case ".tar":
		return extractTar(archivePath, destDir, files)
	default:
		return nil, fmt.Errorf("unsupported archive format: %s", ext)
	}
}

func detectBinariesZip(path string) ([]string, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	
	var binaries []string
	for _, f := range r.File {
		if isExecutable(f.Name, f.Mode()) && !f.FileInfo().IsDir() {
			binaries = append(binaries, f.Name)
		}
	}
	
	return binaries, nil
}

func detectBinariesTarGz(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	gzr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()
	
	return detectBinariesFromTar(tar.NewReader(gzr))
}

func detectBinariesTar(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	return detectBinariesFromTar(tar.NewReader(f))
}

func detectBinariesFromTar(tr *tar.Reader) ([]string, error) {
	var binaries []string
	
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		
		if isExecutable(header.Name, header.FileInfo().Mode()) && !header.FileInfo().IsDir() {
			binaries = append(binaries, header.Name)
		}
	}
	
	return binaries, nil
}

func isExecutable(name string, mode os.FileMode) bool {
	// Check if in bin/ directory
	if strings.Contains(name, "/bin/") || strings.HasPrefix(name, "bin/") {
		return true
	}
	
	// Check if file has executable bit
	if mode&0111 != 0 {
		// Exclude common non-binary executables
		ext := strings.ToLower(filepath.Ext(name))
		if ext == ".sh" || ext == ".py" || ext == ".rb" || ext == ".pl" {
			return false
		}
		return true
	}
	
	return false
}

func extractZip(archivePath, destDir string, files []string) ([]string, error) {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	
	fileSet := make(map[string]bool)
	for _, f := range files {
		fileSet[f] = true
	}
	
	var extracted []string
	for _, f := range r.File {
		if !fileSet[f.Name] {
			continue
		}
		
		destPath := filepath.Join(destDir, filepath.Base(f.Name))
		
		if err := extractZipFile(f, destPath); err != nil {
			return extracted, err
		}
		
		extracted = append(extracted, destPath)
	}
	
	return extracted, nil
}

func extractZipFile(f *zip.File, destPath string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	
	outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()
	
	_, err = io.Copy(outFile, rc)
	return err
}

func extractTarGz(archivePath, destDir string, files []string) ([]string, error) {
	f, err := os.Open(archivePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	gzr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()
	
	return extractFromTar(tar.NewReader(gzr), destDir, files)
}

func extractTar(archivePath, destDir string, files []string) ([]string, error) {
	f, err := os.Open(archivePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	return extractFromTar(tar.NewReader(f), destDir, files)
}

func extractFromTar(tr *tar.Reader, destDir string, files []string) ([]string, error) {
	fileSet := make(map[string]bool)
	for _, f := range files {
		fileSet[f] = true
	}
	
	var extracted []string
	
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return extracted, err
		}
		
		if !fileSet[header.Name] {
			continue
		}
		
		destPath := filepath.Join(destDir, filepath.Base(header.Name))
		
		outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, header.FileInfo().Mode())
		if err != nil {
			return extracted, err
		}
		
		if _, err := io.Copy(outFile, tr); err != nil {
			outFile.Close()
			return extracted, err
		}
		outFile.Close()
		
		extracted = append(extracted, destPath)
	}
	
	return extracted, nil
}
