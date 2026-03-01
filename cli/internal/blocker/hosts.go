package blocker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xmanny/lockin/cli/internal/config"
)

const (
	hostsFile   = "/etc/hosts"
	markerStart = "# LOCKIN-START — do not edit this block"
	markerEnd   = "# LOCKIN-END"
)

func backupPath() (string, error) {
	dir, err := config.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "hosts.bak"), nil
}

func backupHosts() error {
	data, err := os.ReadFile(hostsFile)
	if err != nil {
		return fmt.Errorf("reading %s: %w", hostsFile, err)
	}

	bp, err := backupPath()
	if err != nil {
		return err
	}

	return os.WriteFile(bp, data, 0644)
}

func BlockWebsites(domains []string) error {
	if err := backupHosts(); err != nil {
		return err
	}

	if err := UnblockWebsites(); err != nil {
		return err
	}

	expanded := config.ExpandWebsites(domains)

	var block strings.Builder
	block.WriteString("\n" + markerStart + "\n")
	for _, d := range expanded {
		block.WriteString(fmt.Sprintf("127.0.0.1  %s\n", d))
	}
	block.WriteString(markerEnd + "\n")

	data, err := os.ReadFile(hostsFile)
	if err != nil {
		return fmt.Errorf("reading %s: %w", hostsFile, err)
	}

	newData := string(data) + block.String()
	if err := os.WriteFile(hostsFile, []byte(newData), 0644); err != nil {
		return fmt.Errorf("writing %s: %w", hostsFile, err)
	}

	return flushDNS()
}

func UnblockWebsites() error {
	data, err := os.ReadFile(hostsFile)
	if err != nil {
		return fmt.Errorf("reading %s: %w", hostsFile, err)
	}

	content := string(data)

	startIdx := strings.Index(content, markerStart)
	if startIdx == -1 {
		return nil
	}

	endIdx := strings.Index(content, markerEnd)
	if endIdx == -1 {
		return nil
	}

	tailIdx := endIdx + len(markerEnd)
	if tailIdx < len(content) && content[tailIdx] == '\n' {
		tailIdx++
	}
	cleaned := content[:startIdx] + content[tailIdx:]
	cleaned = strings.TrimRight(cleaned, "\n") + "\n"

	if err := os.WriteFile(hostsFile, []byte(cleaned), 0644); err != nil {
		return fmt.Errorf("writing %s: %w", hostsFile, err)
	}

	return flushDNS()
}

func HostsBlocked() bool {
	data, err := os.ReadFile(hostsFile)
	if err != nil {
		return false
	}
	return strings.Contains(string(data), markerStart)
}

func flushDNS() error {
	_ = runCommand("dscacheutil", "-flushcache")
	return runCommand("killall", "-HUP", "mDNSResponder")
}
