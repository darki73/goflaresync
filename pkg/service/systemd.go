package service

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
)

// SystemdConfigurator is the systemd configurator.
type SystemdConfigurator struct {
	// baseDirectory is the base directory for service installation
	baseDirectory string
	// serviceName is the name of the service
	serviceName string
	// processIdentifierHandler is the process identifier handler
	processIdentifierHandler *ProcessIdentifierHandler
}

// NewSystemdConfigurator creates a new systemd configurator.
func NewSystemdConfigurator() *SystemdConfigurator {
	return &SystemdConfigurator{
		baseDirectory:            "/etc/systemd/system",
		serviceName:              fmt.Sprintf("%s.service", applicationName()),
		processIdentifierHandler: NewProcessIdentifierHandler(),
	}
}

// GetServiceTemplate returns the service template.
func (systemdConfigurator *SystemdConfigurator) GetServiceTemplate() (string, error) {
	unitInformation := []string{
		"[Unit]",
		"Description=GoFlareSync Service",
		"After=network.target",
	}

	executablePath, err := applicationFullPath()
	if err != nil {
		return "", err
	}

	serviceInformation := []string{
		"[Service]",
		"Type=simple",
		"Restart=on-failure",
		fmt.Sprintf("ExecStart=%s start", executablePath),
		fmt.Sprintf("ExecStop=%s stop", executablePath),
		fmt.Sprintf(
			"PIDFile=%s",
			systemdConfigurator.processIdentifierHandler.GetProcessIdentifierFile(),
		),
	}

	installInformation := []string{
		"[Install]",
		"WantedBy=multi-user.target",
	}

	return combineSlices(unitInformation, serviceInformation, installInformation), nil
}

// GetServicePath returns the service path.
func (systemdConfigurator *SystemdConfigurator) GetServicePath() string {
	return path.Join(systemdConfigurator.baseDirectory, systemdConfigurator.serviceName)
}

// GetProcessIdentifierHandler returns the process identifier handler.
func (systemdConfigurator *SystemdConfigurator) GetProcessIdentifierHandler() *ProcessIdentifierHandler {
	return systemdConfigurator.processIdentifierHandler
}

// CreateService creates the service.
func (systemdConfigurator *SystemdConfigurator) CreateService() error {
	systemdConfigurator.IsRunningAsRoot()

	if systemdConfigurator.IsServiceExists() {
		logInfo("service already exists", nil)
		return nil
	}

	logInfo("creating the service", nil)

	template, err := systemdConfigurator.GetServiceTemplate()

	if err != nil {
		return err
	}

	err = os.WriteFile(systemdConfigurator.GetServicePath(), []byte(template), 0644)

	if err != nil {
		logDebugf("failed to write service file: %s", nil, err.Error())
		return err
	}

	logInfo("successfully created the service", nil)

	return systemdConfigurator.ReloadManager()
}

// DeleteService deletes the service.
func (systemdConfigurator *SystemdConfigurator) DeleteService() error {
	systemdConfigurator.IsRunningAsRoot()

	if !systemdConfigurator.IsServiceExists() {
		logInfo("service does not exist", nil)
		return nil
	}

	logInfo("deleting the service", nil)

	isRunning, err := systemdConfigurator.IsServiceRunning()
	if err != nil {
		return err
	}

	if isRunning {
		if err := systemdConfigurator.StopService(); err != nil {
			return err
		}
	}

	if err := os.Remove(systemdConfigurator.GetServicePath()); err != nil {
		return err
	}

	logInfo("successfully deleted the service", nil)

	return systemdConfigurator.ReloadManager()
}

// EnableService enables the service.
func (systemdConfigurator *SystemdConfigurator) EnableService() error {
	systemdConfigurator.IsRunningAsRoot()

	if !systemdConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	logInfof("enabling the service", nil)

	isEnabled, err := systemdConfigurator.IsServiceEnabled()

	if err != nil {
		return err
	}

	if !isEnabled {
		if err := exec.Command("systemctl", "enable", systemdConfigurator.serviceName).Run(); err != nil {
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
func (systemdConfigurator *SystemdConfigurator) DisableService() error {
	systemdConfigurator.IsRunningAsRoot()

	if !systemdConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	logInfo("disabling the service", nil)

	isEnabled, err := systemdConfigurator.IsServiceEnabled()

	if err != nil {
		return err
	}

	if isEnabled {
		if err := exec.Command("systemctl", "disable", systemdConfigurator.serviceName).Run(); err != nil {
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
func (systemdConfigurator *SystemdConfigurator) StartService() error {
	systemdConfigurator.IsRunningAsRoot()

	if !systemdConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	logInfo("starting the service", nil)

	isRunning, err := systemdConfigurator.IsServiceRunning()

	if err != nil {
		return err
	}

	if !isRunning {
		if err := exec.Command("systemctl", "start", systemdConfigurator.serviceName).Run(); err != nil {
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
func (systemdConfigurator *SystemdConfigurator) StopService() error {
	systemdConfigurator.IsRunningAsRoot()

	if !systemdConfigurator.IsServiceExists() {
		return fmt.Errorf("service does not exist")
	}

	logInfo("stopping the service", nil)

	isRunning, err := systemdConfigurator.IsServiceRunning()

	if err != nil {
		return err
	}

	if isRunning {
		if err := exec.Command("systemctl", "stop", systemdConfigurator.serviceName).Run(); err != nil {
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
func (systemdConfigurator *SystemdConfigurator) RestartService() error {
	systemdConfigurator.IsRunningAsRoot()

	logInfo("restarting the service", nil)

	if err := systemdConfigurator.StopService(); err != nil {
		return err
	}

	if err := systemdConfigurator.StartService(); err != nil {
		return err
	}

	logInfo("successfully restarted the service", nil)

	return nil
}

// ReloadManager reloads the manager.
func (systemdConfigurator *SystemdConfigurator) ReloadManager() error {
	logInfo("reloading the manager", nil)

	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		logDebugf("failed to reload systemd: %s", nil, err.Error())
		return err
	}

	logInfo("successfully reloaded the manager", nil)

	return nil
}

// IsRunningAsRoot checks if the application is running as root.
func (systemdConfigurator *SystemdConfigurator) IsRunningAsRoot() {
	isRunningAsRoot()
}

// IsServiceExists checks if the service exists.
func (systemdConfigurator *SystemdConfigurator) IsServiceExists() bool {
	return isServiceExists(systemdConfigurator.baseDirectory, systemdConfigurator.serviceName)
}

// IsServiceEnabled checks if the service is enabled.
func (systemdConfigurator *SystemdConfigurator) IsServiceEnabled() (bool, error) {
	cmd := exec.Command("systemctl", "is-enabled", systemdConfigurator.serviceName)
	logDebugf("executing command: %s", nil, cmd.String())
	var out bytes.Buffer
	cmd.Stdout = &out

	_ = cmd.Run()

	switch out.String() {
	case "enabled\n":
		return true, nil
	case "disabled\n":
		return false, nil
	default:
		return false, errors.New("Unexpected systemctl output: " + out.String())
	}
}

// IsServiceRunning checks if the service is running.
func (systemdConfigurator *SystemdConfigurator) IsServiceRunning() (bool, error) {
	cmd := exec.Command("systemctl", "is-active", systemdConfigurator.serviceName)
	logDebugf("executing command: %s", nil, cmd.String())
	var out bytes.Buffer
	cmd.Stdout = &out

	_ = cmd.Run()

	switch out.String() {
	case "active\n":
		return true, nil
	case "inactive\n":
		return false, nil
	default:
		return false, errors.New("Unexpected systemctl output: " + out.String())
	}
}
