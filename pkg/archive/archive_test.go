package archive

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"
)

func TestIsExecutable(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		mode     os.FileMode
		expected bool
	}{
		{"Binary in bin dir", "app/bin/tool", 0644, true},
		{"Binary with exec bit", "tool", 0755, true},
		{"Shell script", "script.sh", 0755, false},
		{"Python script", "script.py", 0755, false},
		{"Regular file", "readme.txt", 0644, false},
		{"Root bin", "bin/mytool", 0755, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isExecutable(tt.path, tt.mode)
			if result != tt.expected {
				t.Errorf("isExecutable(%q, %v) = %v; want %v",
					tt.path, tt.mode, result, tt.expected)
			}
		})
	}
}

func TestDetectBinariesZip(t *testing.T) {
	// Create a temporary zip file for testing
	tmpDir := t.TempDir()
	zipPath := filepath.Join(tmpDir, "test.zip")

	// Create zip with test content
	f, err := os.Create(zipPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	w := zip.NewWriter(f)

	// Add a binary in bin/
	fw, err := w.CreateHeader(&zip.FileHeader{
		Name:               "myapp/bin/tool",
		Method:             zip.Deflate,
		ExternalAttrs:      0755 << 16,
		CreatorVersion:     3 << 8,
		UncompressedSize64: 100,
	})
	if err != nil {
		t.Fatal(err)
	}
	fw.Write([]byte("fake binary content"))

	// Add a non-binary file
	fw2, err := w.Create("myapp/README.md")
	if err != nil {
		t.Fatal(err)
	}
	fw2.Write([]byte("readme content"))

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Test detection
	binaries, err := detectBinariesZip(zipPath)
	if err != nil {
		t.Fatalf("detectBinariesZip failed: %v", err)
	}

	if len(binaries) != 1 {
		t.Errorf("Expected 1 binary, got %d", len(binaries))
	}

	if len(binaries) > 0 && binaries[0] != "myapp/bin/tool" {
		t.Errorf("Expected myapp/bin/tool, got %s", binaries[0])
	}
}

func TestDetectBinariesTar(t *testing.T) {
	tmpDir := t.TempDir()
	tarPath := filepath.Join(tmpDir, "test.tar")

	// Create tar with test content
	f, err := os.Create(tarPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	tw := tar.NewWriter(f)

	// Add a binary in bin/
	content := []byte("fake binary content")
	header := &tar.Header{
		Name: "myapp/bin/tool",
		Mode: 0755,
		Size: int64(len(content)),
	}
	if err := tw.WriteHeader(header); err != nil {
		t.Fatal(err)
	}
	if _, err := tw.Write(content); err != nil {
		t.Fatal(err)
	}

	// Add a non-binary file
	content2 := []byte("readme content")
	header2 := &tar.Header{
		Name: "myapp/README.md",
		Mode: 0644,
		Size: int64(len(content2)),
	}
	if err := tw.WriteHeader(header2); err != nil {
		t.Fatal(err)
	}
	if _, err := tw.Write(content2); err != nil {
		t.Fatal(err)
	}

	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}

	// Test detection
	binaries, err := detectBinariesTar(tarPath)
	if err != nil {
		t.Fatalf("detectBinariesTar failed: %v", err)
	}

	if len(binaries) != 1 {
		t.Errorf("Expected 1 binary, got %d", len(binaries))
	}

	if len(binaries) > 0 && binaries[0] != "myapp/bin/tool" {
		t.Errorf("Expected myapp/bin/tool, got %s", binaries[0])
	}
}

func TestDetectBinariesTarGz(t *testing.T) {
	tmpDir := t.TempDir()
	tarGzPath := filepath.Join(tmpDir, "test.tar.gz")

	// Create tar.gz with test content
	f, err := os.Create(tarGzPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	tw := tar.NewWriter(gw)

	// Add a binary in bin/
	header := &tar.Header{
		Name: "myapp/bin/tool",
		Mode: 0755,
		Size: 19,
	}
	if err := tw.WriteHeader(header); err != nil {
		t.Fatal(err)
	}
	if _, err := tw.Write([]byte("fake binary content")); err != nil {
		t.Fatal(err)
	}

	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gw.Close(); err != nil {
		t.Fatal(err)
	}

	// Test detection
	binaries, err := detectBinariesTarGz(tarGzPath)
	if err != nil {
		t.Fatalf("detectBinariesTarGz failed: %v", err)
	}

	if len(binaries) != 1 {
		t.Errorf("Expected 1 binary, got %d", len(binaries))
	}

	if len(binaries) > 0 && binaries[0] != "myapp/bin/tool" {
		t.Errorf("Expected myapp/bin/tool, got %s", binaries[0])
	}
}

func TestExtract(t *testing.T) {
	tmpDir := t.TempDir()
	zipPath := filepath.Join(tmpDir, "test.zip")
	destDir := filepath.Join(tmpDir, "dest")

	if err := os.MkdirAll(destDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create test zip
	f, err := os.Create(zipPath)
	if err != nil {
		t.Fatal(err)
	}

	w := zip.NewWriter(f)
	fw, err := w.CreateHeader(&zip.FileHeader{
		Name:          "myapp/bin/tool",
		Method:        zip.Deflate,
		ExternalAttrs: 0755 << 16,
	})
	if err != nil {
		t.Fatal(err)
	}
	fw.Write([]byte("test content"))
	w.Close()
	f.Close()

	// Test extraction
	files := []string{"myapp/bin/tool"}
	extracted, err := Extract(zipPath, destDir, files)
	if err != nil {
		t.Fatalf("Extract failed: %v", err)
	}

	if len(extracted) != 1 {
		t.Errorf("Expected 1 extracted file, got %d", len(extracted))
	}

	// Verify file exists
	extractedPath := filepath.Join(destDir, "tool")
	if _, err := os.Stat(extractedPath); os.IsNotExist(err) {
		t.Errorf("Extracted file does not exist: %s", extractedPath)
	}

	// Verify content
	content, err := os.ReadFile(extractedPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "test content" {
		t.Errorf("Expected 'test content', got '%s'", string(content))
	}
}
