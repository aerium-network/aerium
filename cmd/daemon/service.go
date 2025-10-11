package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/aerium-network/aerium/cmd"
	"github.com/spf13/cobra"
)

const (
	_systemdServicePath = "/etc/systemd/system/aerium.service"
	_serviceTemplate    = `[Unit]
Description=Aerium Blockchain Daemon
Documentation=https://aerium.dev
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=%s
WorkingDirectory=%s
ExecStart=%s start -w %s%s
Restart=on-failure
RestartSec=10
StandardOutput=journal
StandardError=journal
SyslogIdentifier=aerium-daemon

# Allow access to working directory and home (SELinux compatible)
ReadWritePaths=%s

# Security hardening (works with SELinux, AppArmor, or no MAC system)
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=read-only

[Install]
WantedBy=multi-user.target
`
)

func buildServiceCmd(parentCmd *cobra.Command) {
	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Manage daemon as a system service in linux",
		Long:  "Manage the Aerium daemon as a system service in Linux using systemd. Requires sudo/root access.",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			// Check if running on Linux
			if runtime.GOOS != "linux" {
				cmd.PrintErrorMsgf("Service management is only supported on Linux systems.")
				cmd.PrintInfoMsgf("Your current operating system: %s", runtime.GOOS)
				os.Exit(1)
			}
		},
	}
	parentCmd.AddCommand(serviceCmd)

	buildInstallServiceCmd(serviceCmd)
	buildUninstallServiceCmd(serviceCmd)
	buildStartServiceCmd(serviceCmd)
	buildStopServiceCmd(serviceCmd)
	buildStatusServiceCmd(serviceCmd)
}

func buildInstallServiceCmd(parentCmd *cobra.Command) {
	serviceCmd := &cobra.Command{
		Use:   "install",
		Short: "Install the Aerium daemon as a system service",
		Long:  "Install the Aerium daemon as a system service. Requires sudo/root access.",
	}
	parentCmd.AddCommand(serviceCmd)

	workingDirOpt := addWorkingDirOption(serviceCmd)

	passwordOpt := serviceCmd.Flags().StringP("password", "p", "",
		"the wallet password")

	passwordFromFileOpt := serviceCmd.Flags().String("password-from-file", "",
		"the file containing the wallet password")

	serviceCmd.Run = func(cobCmd *cobra.Command, _ []string) {
		// Check if running as root
		if !isRoot() {
			cmd.PrintErrorMsgf("This command requires root privileges.")
			cmd.PrintInfoMsgf("Please run with sudo:")
			cmd.PrintInfoMsgf("  sudo aerium-daemon service install")
			os.Exit(1)
		}

		workingDir, _ := filepath.Abs(*workingDirOpt)
		err := os.Chdir(workingDir)
		cmd.FatalErrorCheck(err)

		passwordFlag := ""
		if *passwordOpt != "" {
			passwordFlag = " -p " + *passwordOpt
		}
		if *passwordFromFileOpt != "" {
			passwordFlag = " --password-from-file " + *passwordFromFileOpt
		}

		// Get the executable path dynamically
		exePath, err := os.Executable()
		cmd.FatalErrorCheck(err)
		exePath, err = filepath.Abs(exePath)
		cmd.FatalErrorCheck(err)

		// Set SELinux context if SELinux is enabled
		cmd.PrintInfoMsgf("Checking SELinux status...")
		selinuxCheck := exec.CommandContext(cobCmd.Context(), "getenforce")
		if output, err := selinuxCheck.Output(); err == nil {
			selinuxStatus := string(output)
			if selinuxStatus == "Enforcing\n" || selinuxStatus == "Permissive\n" {
				cmd.PrintInfoMsgf("SELinux is enabled. Setting proper context for executable...")
				// Set the bin_t context which allows systemd to execute the binary
				selinuxCmd := exec.CommandContext(cobCmd.Context(), "chcon", "-t", "bin_t", exePath)
				if err := selinuxCmd.Run(); err != nil {
					cmd.PrintWarnMsgf("Warning: Failed to set SELinux context: %v", err)
					cmd.PrintInfoMsgf("You may need to manually run: sudo chcon -t bin_t %s", exePath)
				} else {
					cmd.PrintInfoMsgf("SELinux context set successfully")
				}
			}
		}

		cmd.PrintInfoMsgf("Installing Aerium daemon as a system service in %s", _systemdServicePath)

		// Get the actual username who invoked sudo
		username := os.Getenv("SUDO_USER")
		if username == "" {
			cmd.PrintWarnMsgf("Warning: Could not detect original user. Service will run as root.")
			cmd.PrintInfoMsgf("It's recommended to run this command with sudo, not as root directly.")
			username = "root"
		}

		unit := fmt.Sprintf(_serviceTemplate, username, workingDir, exePath, workingDir, passwordFlag, workingDir)
		err = os.WriteFile(_systemdServicePath, []byte(unit), 0o600)
		cmd.FatalErrorCheck(err)

		// Reload systemd daemon
		cmd.PrintInfoMsgf("Reloading systemd daemon...")
		exe := exec.CommandContext(cobCmd.Context(), "systemctl", "daemon-reload")
		err = exe.Run()
		cmd.FatalErrorCheck(err)

		// Enable service to start on boot
		cmd.PrintInfoMsgf("Enabling service to start on boot...")
		exe = exec.CommandContext(cobCmd.Context(), "systemctl", "enable", "aerium")
		err = exe.Run()
		cmd.FatalErrorCheck(err)

		cmd.PrintSuccessMsgf("System service installed and enabled successfully!")
		cmd.PrintLine()
		cmd.PrintInfoMsgf("The service will run as user: %s", username)
		cmd.PrintInfoMsgf("Start the service with:")
		cmd.PrintInfoMsgf("  sudo aerium-daemon service start")
		cmd.PrintInfoMsgf("Or use:")
		cmd.PrintInfoMsgf("  sudo systemctl start aerium")
	}
}

func buildUninstallServiceCmd(parentCmd *cobra.Command) {
	serviceCmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall the Aerium daemon system service",
		Long:  "Uninstall the Aerium daemon system service. Requires sudo/root access.",
	}
	parentCmd.AddCommand(serviceCmd)

	serviceCmd.Run = func(cobCmd *cobra.Command, _ []string) {
		// Check if running as root
		if !isRoot() {
			cmd.PrintErrorMsgf("This command requires root privileges.")
			cmd.PrintInfoMsgf("Please run with sudo:")
			cmd.PrintInfoMsgf("  sudo aerium-daemon service uninstall")
			os.Exit(1)
		}

		cmd.PrintInfoMsgf("Uninstalling Aerium daemon system service...")

		// Stop the service if running
		cmd.PrintInfoMsgf("Stopping service if running...")
		exe := exec.CommandContext(cobCmd.Context(), "systemctl", "stop", "aerium")
		_ = exe.Run() // Ignore error if service is not running

		// Disable the service
		cmd.PrintInfoMsgf("Disabling service...")
		exe = exec.CommandContext(cobCmd.Context(), "systemctl", "disable", "aerium")
		_ = exe.Run() // Ignore error if service is not enabled

		// Remove service file
		if _, err := os.Stat(_systemdServicePath); err == nil {
			err = os.Remove(_systemdServicePath)
			cmd.FatalErrorCheck(err)
			cmd.PrintInfoMsgf("Removed service file: %s", _systemdServicePath)
		}

		// Reload systemd daemon
		cmd.PrintInfoMsgf("Reloading systemd daemon...")
		exe = exec.CommandContext(cobCmd.Context(), "systemctl", "daemon-reload")
		err := exe.Run()
		cmd.FatalErrorCheck(err)

		cmd.PrintSuccessMsgf("System service uninstalled successfully!")
	}
}

func buildStartServiceCmd(parentCmd *cobra.Command) {
	serviceCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the Aerium daemon system service",
		Long:  "Start the Aerium daemon system service. Requires sudo/root access.",
	}
	parentCmd.AddCommand(serviceCmd)

	serviceCmd.Run = func(cobCmd *cobra.Command, _ []string) {
		// Check if running as root
		if !isRoot() {
			cmd.PrintErrorMsgf("This command requires root privileges.")
			cmd.PrintInfoMsgf("Please run with sudo:")
			cmd.PrintInfoMsgf("  sudo aerium-daemon service start")
			os.Exit(1)
		}

		cmd.PrintInfoMsgf("Starting Aerium daemon service...")
		exe := exec.CommandContext(cobCmd.Context(), "systemctl", "start", "aerium")
		output, err := exe.CombinedOutput()
		if err != nil {
			cmd.PrintErrorMsgf("Failed to start service: %s", string(output))
			cmd.FatalErrorCheck(err)
		}

		cmd.PrintSuccessMsgf("Service started successfully!")
		cmd.PrintInfoMsgf("Check status with:")
		cmd.PrintInfoMsgf("  sudo aerium-daemon service status")
	}
}

func buildStopServiceCmd(parentCmd *cobra.Command) {
	serviceCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop the Aerium daemon system service",
		Long:  "Stop the Aerium daemon system service. Requires sudo/root access.",
	}
	parentCmd.AddCommand(serviceCmd)

	serviceCmd.Run = func(cobCmd *cobra.Command, _ []string) {
		// Check if running as root
		if !isRoot() {
			cmd.PrintErrorMsgf("This command requires root privileges.")
			cmd.PrintInfoMsgf("Please run with sudo:")
			cmd.PrintInfoMsgf("  sudo aerium-daemon service stop")
			os.Exit(1)
		}

		cmd.PrintInfoMsgf("Stopping Aerium daemon service...")
		exe := exec.CommandContext(cobCmd.Context(), "systemctl", "stop", "aerium")
		output, err := exe.CombinedOutput()
		if err != nil {
			cmd.PrintErrorMsgf("Failed to stop service: %s", string(output))
			cmd.FatalErrorCheck(err)
		}

		// Reset the failed state if the service was in failed state
		cmd.PrintInfoMsgf("Resetting service state...")
		exe = exec.CommandContext(cobCmd.Context(), "systemctl", "reset-failed", "aerium")
		_ = exe.Run() // Ignore error if service was not failed

		cmd.PrintSuccessMsgf("Service stopped successfully!")
	}
}

func buildStatusServiceCmd(parentCmd *cobra.Command) {
	serviceCmd := &cobra.Command{
		Use:   "status",
		Short: "Print the status of the Aerium daemon system service",
		Long:  "Print the status of the Aerium daemon system service. Requires sudo/root access.",
	}
	parentCmd.AddCommand(serviceCmd)

	serviceCmd.Run = func(cobCmd *cobra.Command, _ []string) {
		// Check if running as root
		if !isRoot() {
			cmd.PrintErrorMsgf("This command requires root privileges.")
			cmd.PrintInfoMsgf("Please run with sudo:")
			cmd.PrintInfoMsgf("  sudo aerium-daemon service status")
			os.Exit(1)
		}

		exe := exec.CommandContext(cobCmd.Context(), "systemctl", "status", "aerium")
		exe.Stdout = os.Stdout
		exe.Stderr = os.Stderr
		_ = exe.Run() // Status command may return non-zero even when service is stopped
	}
}

func isRoot() bool {
	return os.Geteuid() == 0
}
