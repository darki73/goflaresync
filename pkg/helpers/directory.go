package helpers

import "os"

// IsDirectoryExists returns a boolean value indicating whether the directory exists.
func IsDirectoryExists(directoryPath string) bool {
	info, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
