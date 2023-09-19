package service

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type InitSystem string

const (
	Systemd  InitSystem = "systemd"
	Upstart  InitSystem = "upstart"
	SysVinit InitSystem = "sysvinit"
	OpenRC   InitSystem = "openrc"
	None     InitSystem = "none"
)

// Manager is the service manager.
type Manager struct {
	// systemManager is the system manager
	systemManager InitializationSystemManager
}

// NewManager creates a new service manager.
func NewManager() *Manager {
	manager := &Manager{
		systemManager: nil,
	}
	return manager.bootstrap()
}

// GetSystemManager returns the system manager.
func (manager *Manager) GetSystemManager() InitializationSystemManager {
	return manager.systemManager
}

// HasSystemManager checks if the system manager is available.
func (manager *Manager) HasSystemManager() bool {
	return manager.systemManager != nil
}

// bootstrap bootstraps the service manager.
func (manager *Manager) bootstrap() *Manager {
	switch manager.detectInitializationSystem() {
	case Systemd:
		manager.systemManager = NewSystemdConfigurator()
		break
	//case Upstart:
	//	manager.systemManager = NewUpstartConfigurator()
	//	break
	case SysVinit:
		manager.systemManager = NewSysVInitConfigurator()
		break
	//case OpenRC:
	//	manager.systemManager = NewOpenRCConfigurator()
	//	break
	case None:
		manager.systemManager = nil
		break
	default:
		manager.systemManager = nil
		break
	}

	return manager
}

// detectInitializationSystem detects the initialization system.
func (manager *Manager) detectInitializationSystem() InitSystem {
	if runtime.GOOS == "windows" || manager.isWindowsSubsystemForLinux() {
		return None
	}

	if manager.isCommandAvailable("systemctl") {
		return Systemd
	}

	if manager.isCommandAvailable("initctl") {
		return Upstart
	}

	if manager.isCommandAvailable("service") && manager.isDirectoryExists("/etc/init.d") {
		return SysVinit
	}

	if manager.isCommandAvailable("rc-service") || manager.isCommandAvailable("rc-update") {
		return OpenRC
	}

	return None
}

// isCommandAvailable checks if the command is available.
func (manager *Manager) isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// isDirectoryExists checks if the directory exists.
func (manager *Manager) isDirectoryExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// isWindowsSubsystemForLinux checks if application is running on Windows Subsystem for Linux.
func (manager *Manager) isWindowsSubsystemForLinux() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}

	versionInfo := string(data)

	if strings.Contains(versionInfo, "Microsoft") || strings.Contains(versionInfo, "microsoft") {
		return true
	}
	return false
}
