package service

import (
	"github.com/darki73/goflaresync/pkg/helpers"
	"github.com/darki73/goflaresync/pkg/log"
	"os"
	"path"
	"strings"
)

// InitializationSystemManager is the initialization system manager.
type InitializationSystemManager interface {
	// GetServiceTemplate returns the service template.
	GetServiceTemplate() (string, error)
	// GetServicePath returns the service path.
	GetServicePath() string
	// GetProcessIdentifierHandler returns the process identifier handler.
	GetProcessIdentifierHandler() *ProcessIdentifierHandler
	// CreateService creates the service.
	CreateService() error
	// DeleteService deletes the service.
	DeleteService() error
	// EnableService enables the service.
	EnableService() error
	// DisableService disables the service.
	DisableService() error
	// StartService starts the service.
	StartService() error
	// StopService stops the service.
	StopService() error
	// RestartService restarts the service.
	RestartService() error
	// ReloadManager reloads the manager.
	ReloadManager() error
	// IsRunningAsRoot checks if the application is running as root.
	IsRunningAsRoot()
	// IsServiceExists checks if the service exists.
	IsServiceExists() bool
	// IsServiceEnabled checks if the service is enabled.
	IsServiceEnabled() (bool, error)
	// IsServiceRunning checks if the service is running.
	IsServiceRunning() (bool, error)
}

// mergeLogFields merges the log fields.
func mergeLogFields(fields log.FieldsMap) log.FieldsMap {
	defaultFields := log.FieldsMap{
		"source": "service",
	}

	if fields == nil {
		return defaultFields
	}

	for key, value := range defaultFields {
		fields[key] = value
	}

	return fields
}

// logDebug logs a debug message.
func logDebug(message string, fields log.FieldsMap) {
	log.DebugWithFields(message, mergeLogFields(fields))
}

// logDebugf logs a debug message.
func logDebugf(message string, fields log.FieldsMap, args ...interface{}) {
	log.DebugfWithFields(message, mergeLogFields(fields), args...)
}

// logInfo logs an info message.
func logInfo(message string, fields log.FieldsMap) {
	log.InfoWithFields(message, mergeLogFields(fields))
}

// logInfof logs an info message.
func logInfof(message string, fields log.FieldsMap, args ...interface{}) {
	log.InfofWithFields(message, mergeLogFields(fields), args...)
}

// applicationName returns the name of the application.
func applicationName() string {
	return helpers.GetExecutableName()
}

// applicationFullPath returns the full path of the application.
func applicationFullPath() (string, error) {
	directory, err := helpers.GetExecutableDirectory()
	if err != nil {
		return "", err
	}
	return path.Join(directory, applicationName()), nil
}

// combineSlices combines the slices.
func combineSlices(slices ...[]string) string {
	var combined []string

	for _, slice := range slices {
		combined = append(combined, slice...)
	}

	return strings.Join(combined, "\n")
}

// isRunningAsRoot checks if the application is running as root.
func isRunningAsRoot() {
	if !helpers.IsRoot() {
		log.Fatal("this command must be run as root")
	}
}

// isServiceExists checks if the service exists.
func isServiceExists(baseDirectory string, serviceName string) bool {
	if _, err := os.Stat(path.Join(baseDirectory, serviceName)); err != nil {
		return false
	}
	return true
}
