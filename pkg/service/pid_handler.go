package service

import (
	"fmt"
	"github.com/darki73/goflaresync/pkg/helpers"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"
)

// ProcessIdentifierHandler is the process identifier handler.
type ProcessIdentifierHandler struct {
	// processIdentifier is the process identifier
	processIdentifier int
	// processIdentifierFileName is the file name containing the process identifier
	processIdentifierFileName string
	// processIdentifierFilePath is the path to the file containing the process identifier
	processIdentifierFilePath string
	// isSupported is a flag indicating whether the process identifier handler is supported
	isSupported bool
}

// NewProcessIdentifierHandler returns a new process identifier handler.
func NewProcessIdentifierHandler() *ProcessIdentifierHandler {
	handler := &ProcessIdentifierHandler{
		processIdentifier:         os.Getpid(),
		processIdentifierFileName: fmt.Sprintf("%s.pid", applicationName()),
		processIdentifierFilePath: "",
		isSupported:               false,
	}

	return handler.bootstrap()
}

// HandleApplicationStart handles the application start.
func (processIdentifierHandler *ProcessIdentifierHandler) HandleApplicationStart() error {
	if !processIdentifierHandler.IsSupported() {
		return nil
	}

	if processIdentifierHandler.IsProcessIdentifierFileExists() {
		processIdentifier, err := processIdentifierHandler.ReadProcessIdentifierFromFile()
		if err != nil {
			return err
		}

		isRunning := processIdentifierHandler.IsProcessRunning(processIdentifier)

		if isRunning && !processIdentifierHandler.IsThisApplication(processIdentifier) {
			logDebug(
				"PID file exists, process is running, but does not belong to this application",
				nil,
			)
			if err := processIdentifierHandler.DeleteProcessIdentifierFile(); err != nil {
				return err
			}
		}

		if !isRunning {
			logDebug(
				"PID file exists, but process is not running",
				nil,
			)
			if err := processIdentifierHandler.DeleteProcessIdentifierFile(); err != nil {
				return err
			}
		}
	}
	return processIdentifierHandler.WriteProcessIdentifierToFile()
}

// HandleApplicationStop handles the application stop.
func (processIdentifierHandler *ProcessIdentifierHandler) HandleApplicationStop() error {
	if !processIdentifierHandler.IsSupported() {
		return nil
	}

	if processIdentifierHandler.IsProcessIdentifierFileExists() {
		processIdentifier, err := processIdentifierHandler.ReadProcessIdentifierFromFile()
		if err != nil {
			return err
		}

		if processIdentifierHandler.IsProcessRunning(processIdentifier) {
			logDebug(
				"process identifier file exists and the process with the given identifier is running",
				nil,
			)

			if processIdentifierHandler.IsThisApplication(processIdentifier) {
				logDebug(
					"process identifier belongs to this application",
					nil,
				)

				process, findErr := os.FindProcess(processIdentifier)
				if findErr != nil {
					return findErr
				}

				if killErr := process.Signal(syscall.SIGTERM); killErr != nil {
					return killErr
				}

				if err := processIdentifierHandler.DeleteProcessIdentifierFile(); err != nil {
					return err
				}
			}
		} else {
			logDebug(
				"process identifier file exists and the process with the given identifier is not running",
				nil,
			)

			if err := processIdentifierHandler.DeleteProcessIdentifierFile(); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetProcessIdentifier returns the process identifier.
func (processIdentifierHandler *ProcessIdentifierHandler) GetProcessIdentifier() int {
	return processIdentifierHandler.processIdentifier
}

// GetProcessIdentifierFileName returns the file name containing the process identifier.
func (processIdentifierHandler *ProcessIdentifierHandler) GetProcessIdentifierFileName() string {
	return processIdentifierHandler.processIdentifierFileName
}

// GetProcessIdentifierFilePath returns the path to the file containing the process identifier.
func (processIdentifierHandler *ProcessIdentifierHandler) GetProcessIdentifierFilePath() string {
	return processIdentifierHandler.processIdentifierFilePath
}

// GetProcessIdentifierFile returns the full path to the file containing the process identifier.
func (processIdentifierHandler *ProcessIdentifierHandler) GetProcessIdentifierFile() string {
	return path.Join(
		processIdentifierHandler.GetProcessIdentifierFilePath(),
		processIdentifierHandler.GetProcessIdentifierFileName(),
	)
}

// IsSupported returns a flag indicating whether the process identifier handler is supported.
func (processIdentifierHandler *ProcessIdentifierHandler) IsSupported() bool {
	return processIdentifierHandler.isSupported
}

// ReadProcessIdentifierFromFile reads the process identifier from the file.
func (processIdentifierHandler *ProcessIdentifierHandler) ReadProcessIdentifierFromFile() (int, error) {
	data, err := os.ReadFile(processIdentifierHandler.GetProcessIdentifierFile())
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(data))
}

// WriteProcessIdentifierToFile writes the process identifier to the file.
func (processIdentifierHandler *ProcessIdentifierHandler) WriteProcessIdentifierToFile() error {
	return os.WriteFile(
		processIdentifierHandler.GetProcessIdentifierFile(),
		[]byte(strconv.Itoa(processIdentifierHandler.GetProcessIdentifier())),
		0644,
	)
}

// DeleteProcessIdentifierFile deletes the file containing the process identifier.
func (processIdentifierHandler *ProcessIdentifierHandler) DeleteProcessIdentifierFile() error {
	return os.Remove(processIdentifierHandler.GetProcessIdentifierFile())
}

// IsProcessIdentifierFileExists returns a flag indicating whether the file containing the process identifier exists.
func (processIdentifierHandler *ProcessIdentifierHandler) IsProcessIdentifierFileExists() bool {
	if _, err := os.Stat(processIdentifierHandler.GetProcessIdentifierFile()); err == nil {
		return true
	}
	return false
}

// IsProcessRunning returns a flag indicating whether the process is running.
func (processIdentifierHandler *ProcessIdentifierHandler) IsProcessRunning(processIdentifier int) bool {
	if processIdentifier == 0 {
		return false
	}

	cmd := exec.Command("kill", "-0", strconv.Itoa(processIdentifier))
	err := cmd.Run()
	return err == nil
}

// IsThisApplication returns a flag indicating whether the process identifier belongs to this application.
func (processIdentifierHandler *ProcessIdentifierHandler) IsThisApplication(processIdentifier int) bool {
	if processIdentifier == 0 {
		return false
	}

	cmd := exec.Command("ps", "-p", strconv.Itoa(processIdentifier), "-o", "args=")
	out, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.Contains(string(out), applicationName())
}

// bootstrap bootstraps the process identifier handler.
func (processIdentifierHandler *ProcessIdentifierHandler) bootstrap() *ProcessIdentifierHandler {
	if helpers.IsDirectoryExists("/var/run") {
		processIdentifierHandler.isSupported = true
		processIdentifierHandler.processIdentifierFilePath = "/var/run"
	} else if helpers.IsDirectoryExists("/run") {
		processIdentifierHandler.isSupported = true
		processIdentifierHandler.processIdentifierFilePath = "/run"
	}

	return processIdentifierHandler
}
