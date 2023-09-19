package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

// GetExecutableDirectory returns the directory containing the executable.
func GetExecutableDirectory() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableDirectory := filepath.Dir(executablePath)
	return executableDirectory, nil
}

// GetExecutableName returns the name of the executable.
func GetExecutableName() string {
	executableName := filepath.Base(os.Args[0])
	if strings.Contains(executableName, ".") {
		executableName = strings.Split(executableName, ".")[0]
	}
	return executableName
}
