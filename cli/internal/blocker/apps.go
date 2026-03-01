package blocker

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type AppBlocker struct {
	blockedApps []string
	interval    time.Duration
	stopCh      chan struct{}
	doneCh      chan struct{}
}

func NewAppBlocker(apps []string, interval time.Duration) *AppBlocker {
	return &AppBlocker{
		blockedApps: apps,
		interval:    interval,
		stopCh:      make(chan struct{}),
		doneCh:      make(chan struct{}),
	}
}

func (ab *AppBlocker) Start() {
	go ab.run()
}

func (ab *AppBlocker) Stop() {
	close(ab.stopCh)
	<-ab.doneCh
}

func (ab *AppBlocker) run() {
	defer close(ab.doneCh)

	ab.killBlockedApps()

	ticker := time.NewTicker(ab.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ab.stopCh:
			return
		case <-ticker.C:
			ab.killBlockedApps()
		}
	}
}

func (ab *AppBlocker) killBlockedApps() {
	running, err := getRunningApps()
	if err != nil {
		return
	}

	blocked := make(map[string]bool, len(ab.blockedApps))
	for _, app := range ab.blockedApps {
		blocked[strings.ToLower(app)] = true
	}

	for _, app := range running {
		if blocked[strings.ToLower(app)] {
			quitApp(app)
		}
	}
}

func getRunningApps() ([]string, error) {
	script := `tell application "System Events" to get name of every process whose background only is false`
	out, err := exec.Command("osascript", "-e", script).Output()
	if err != nil {
		return nil, fmt.Errorf("listing apps: %w", err)
	}

	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return nil, nil
	}

	parts := strings.Split(raw, ", ")
	apps := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			apps = append(apps, p)
		}
	}
	return apps, nil
}

func quitApp(name string) {
	escaped := strings.ReplaceAll(name, `\`, `\\`)
	escaped = strings.ReplaceAll(escaped, `"`, `\"`)
	script := fmt.Sprintf(`tell application "%s" to quit`, escaped)
	_ = exec.Command("osascript", "-e", script).Run()
}
