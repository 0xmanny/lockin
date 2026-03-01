package blocker

import (
	"os"
	"os/exec"
)

const (
	plistPath    = "/Library/LaunchDaemons/sh.lockin.cleanup.plist"
	plistContent = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>sh.lockin.cleanup</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/lockin</string>
        <string>stop</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
`
)

// InstallLaunchDaemon writes the cleanup plist to /Library/LaunchDaemons/.
// this ensures /etc/hosts is restored on reboot if the machine dies while
// lockin is active. idempotent, safe to call multiple times.
func InstallLaunchDaemon() error {
	if _, err := os.Stat(plistPath); err == nil {
		return nil
	}

	if err := os.WriteFile(plistPath, []byte(plistContent), 0644); err != nil {
		return err
	}

	_ = exec.Command("launchctl", "load", plistPath).Run()
	return nil
}

// removes the cleanup plist
func UninstallLaunchDaemon() error {
	_ = exec.Command("launchctl", "unload", plistPath).Run()
	return os.Remove(plistPath)
}
