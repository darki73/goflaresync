package service

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

// SysVInitConfigurator is the init.d configurator.
type SysVInitConfigurator struct {
	// baseDirectory is the base directory for service installation
	baseDirectory string
	// serviceName is the name of the service
	serviceName string
	// processIdentifierHandler is the process identifier handler
	processIdentifierHandler *ProcessIdentifierHandler
}

// NewSysVInitConfigurator creates a new initd configurator.
func NewSysVInitConfigurator() *SysVInitConfigurator {
	return &SysVInitConfigurator{
		baseDirectory:            "/etc/init.d",
		serviceName:              applicationName(),
		processIdentifierHandler: NewProcessIdentifierHandler(),
	}
}

// GetServiceTemplate returns the service template.
func (sysVInitConfigurator *SysVInitConfigurator) GetServiceTemplate() (string, error) {
	executablePath, err := applicationFullPath()
	if err != nil {
		return "", err
	}

	scriptInformation := []string{
		"#!/bin/bash",
		fmt.Sprintf("# %s/%s", sysVInitConfigurator.baseDirectory, sysVInitConfigurator.serviceName),
		"\n",
		"### BEGIN INIT INFO",
		fmt.Sprintf("# Provides:          %s", sysVInitConfigurator.serviceName),
		"# Required-Start:    $local_fs $network",
		"# Required-Stop:     $local_fs",
		"# Default-Start:     2 3 4 5",
		"# Default-Stop:      0 1 6",
		"# Short-Description: GoFlareSync Service",
		"# Description:       Manager GoFlareSync Service",
		"### END INIT INFO",
		"\n",
		fmt.Sprintf("EXEC_PATH=\"%s\"", executablePath),
		fmt.Sprintf("PID_FILE=\"%s\"", sysVInitConfigurator.processIdentifierHandler.GetProcessIdentifierFile()),
	}

	startFunction := []string{
		"start() {",
		"	echo \"Starting GoFlareSync Service\"",
		"	$EXEC_PATH start",
		"}",
	}

	stopFunction := []string{
		"stop() {",
		"	echo \"Stopping GoFlareSync Service\"",
		"	$EXEC_PATH stop",
		"}",
	}

	restartFunction := []string{
		"restart() {",
		"   stop",
		"   start",
		"}",
	}

	statusFunction := []string{
		"status() {",
		"   if [ -e $PID_FILE ]; then",
		"       echo \"GoFlareSync Service is running\"",
		"   else",
		"       echo \"GoFlareSync Service is not running\"",
		"   fi",
		"}",
	}

	serviceDefinition := []string{
		"case \"$1\" in",
		"	start)",
		"		start",
		"		;;",
		"	stop)",
		"		stop",
		"		;;",
		"	restart)",
		"		restart",
		"		;;",
		"	status)",
		"		status",
		"		;;",
		"	*)",
		"		echo \"Usage: $0 {start|stop|restart|status}\"",
		"		exit 1",
		"		;;",
		"esac",
	}

	return combineSlices(
		scriptInformation,
		startFunction,
		stopFunction,
		restartFunction,
		statusFunction,
		serviceDefinition,
	), nil
}

// GetServicePath returns the service path.
func (sysVInitConfigurator *SysVInitConfigurator) GetServicePath() string {
	return path.Join(sysVInitConfigurator.baseDirectory, sysVInitConfigurator.serviceName)
}

// GetProcessIdentifierHandler returns the process identifier handler.
func (sysVInitConfigurator *SysVInitConfigurator) GetProcessIdentifierHandler() *ProcessIdentifierHandler {
	return sysVInitConfigurator.processIdentifierHandler
}

// CreateService creates the service.
func (sysVInitConfigurator *SysVInitConfigurator) CreateService() error {
	sysVInitConfigurator.IsRunningAsRoot()

	if sysVInitConfigurator.IsServiceExists() {
		logInfo("service already exists", nil)
		return nil
	}

	logInfo("creating the service", nil)

	template, err := sysVInitConfigurator.GetServiceTemplate()

	if err != nil {
		return err
	}

	err = os.WriteFile(sysVInitConfigurator.GetServicePath(), []byte(template), 0755)

	if err != nil {
		logDebugf("failed to write service file: %s", nil, err.Error())
		return err
	}

	logInfo("successfully created the service", nil)

	return sysVInitConfigurator.ReloadManager()
}

// DeleteService deletes the service.
func (sysVInitConfigurator *SysVInitConfigurator) DeleteService() error {
	sysVInitConfigurator.IsRunningAsRoot()

	if !sysVInitConfigurator.IsServiceExists() {
		logInfo("service does not exist", nil)
		return nil
	}

	logInfo("deleting the service", nil)

	isRunning, err := sysVInitConfigurator.IsServiceRunning()

	if err != nil {
		return err
	}

	if isRunning {
		if err := sysVInitConfigurator.StopService(); err != nil {
			return err
		}
	}

	if err := sysVInitConfigurator.DisableService(); err != nil {
		return err
	}

	if err := os.Remove(sysVInitConfigurator.GetServicePath()); err != nil {
		return err
	}

	logInfo("successfully deleted the service", nil)

	return sysVInitConfigurator.ReloadManager()
}

// EnableService enables the service.
func (sysVInitConfigurator *SysVInitConfigurator) EnableService() error {
	sysVInitConfigurator.IsRunningAsRoot()

	if !sysVInitConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	isEnabled, err := sysVInitConfigurator.IsServiceEnabled()
	if err != nil {
		return err
	}

	if !isEnabled {
		if err := exec.Command("update-rc.d", sysVInitConfigurator.serviceName, "defaults").Run(); err != nil {
			logDebugf("failed to enable the service: %s", nil, err.Error())
			return err
		}
		logInfo("successfully enabled the service", nil)
	} else {
		logInfo("service is already enabled", nil)
	}

	return nil
}

// DisableService disables the service.
func (sysVInitConfigurator *SysVInitConfigurator) DisableService() error {
	sysVInitConfigurator.IsRunningAsRoot()

	if !sysVInitConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	logInfo("disabling the service", nil)

	isEnabled, err := sysVInitConfigurator.IsServiceEnabled()
	if err != nil {
		return err
	}

	if isEnabled {
		if err := exec.Command("update-rc.d", "-f", sysVInitConfigurator.serviceName, "remove").Run(); err != nil {
			logDebugf("failed to disable the service: %s", nil, err.Error())
			return err
		}
		logInfo("successfully disabled the service", nil)
	} else {
		logInfo("service is already disabled", nil)
	}

	return nil
}

// StartService starts the service.
func (sysVInitConfigurator *SysVInitConfigurator) StartService() error {
	sysVInitConfigurator.IsRunningAsRoot()

	if !sysVInitConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	isRunning, _ := sysVInitConfigurator.IsServiceRunning()

	if !isRunning {
		if err := exec.Command(sysVInitConfigurator.GetServicePath(), "start").Run(); err != nil {
			logDebugf("failed to start the service: %s", nil, err.Error())
			return err
		}
		logInfo("successfully started the service", nil)
	} else {
		logInfo("service is already running", nil)
	}

	return nil
}

// StopService stops the service.
func (sysVInitConfigurator *SysVInitConfigurator) StopService() error {
	sysVInitConfigurator.IsRunningAsRoot()

	if !sysVInitConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	isRunning, _ := sysVInitConfigurator.IsServiceRunning()

	if isRunning {
		if err := exec.Command(sysVInitConfigurator.GetServicePath(), "stop").Run(); err != nil {
			logDebugf("failed to stop the service: %s", nil, err.Error())
			return err
		}
		logInfo("successfully stopped the service", nil)
	} else {
		logInfo("service is already stopped", nil)
	}

	return nil
}

// RestartService restarts the service.
func (sysVInitConfigurator *SysVInitConfigurator) RestartService() error {
	sysVInitConfigurator.IsRunningAsRoot()

	logInfo("restarting the service", nil)

	if err := sysVInitConfigurator.StopService(); err != nil {
		return err
	}

	if err := sysVInitConfigurator.StartService(); err != nil {
		return err
	}

	logInfo("successfully restarted the service", nil)

	return nil
}

// ReloadManager reloads the manager.
func (sysVInitConfigurator *SysVInitConfigurator) ReloadManager() error {
	return nil
}

// IsRunningAsRoot checks if the application is running as root.
func (sysVInitConfigurator *SysVInitConfigurator) IsRunningAsRoot() {
	isRunningAsRoot()
}

// IsServiceExists checks if the service exists.
func (sysVInitConfigurator *SysVInitConfigurator) IsServiceExists() bool {
	return isServiceExists(sysVInitConfigurator.baseDirectory, sysVInitConfigurator.serviceName)
}

// IsServiceEnabled checks if the service is enabled.
func (sysVInitConfigurator *SysVInitConfigurator) IsServiceEnabled() (bool, error) {
	runLevels := []string{"0", "1", "2", "3", "4", "5", "6", "S"}

	for _, runlevel := range runLevels {
		dirPath := filepath.Join("/etc/rc" + runlevel + ".d/")
		enabled, err := filepath.Glob(filepath.Join(dirPath, "S*"+sysVInitConfigurator.serviceName))
		if err != nil {
			return false, err
		}
		if len(enabled) > 0 {
			return true, nil
		}
	}

	return false, nil
}

// IsServiceRunning checks if the service is running.
func (sysVInitConfigurator *SysVInitConfigurator) IsServiceRunning() (bool, error) {
	return sysVInitConfigurator.processIdentifierHandler.IsProcessIdentifierFileExists(), nil
}
