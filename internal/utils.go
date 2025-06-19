package scanner

import (
	"os"
	"path/filepath"
)

// FileExists checks if a file exists and is not a directory.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsSupportedFile checks if the extension is supported.
func IsSupportedFile(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".env", ".json", ".yaml", ".yml", ".txt":
		return true
	default:
		return false
	}
}
