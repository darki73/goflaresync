package helpers

import (
	"os"
	"os/exec"
)

// HasSystemd checks if systemd is available.
func HasSystemd() bool {
	_, err := exec.LookPath("systemctl")
	return err == nil
}

// IsRoot checks if the current user is root.
func IsRoot() bool {
	return os.Geteuid() == 0
}
