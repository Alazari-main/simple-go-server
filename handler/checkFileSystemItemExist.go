package handler

import (
	"os"
)

// checkFileSystemItemExists scans file system for a specified file or folder and returns true if the it exists.
func checkFileSystemItemExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
