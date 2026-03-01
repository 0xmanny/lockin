package blocker

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/0xmanny/lockin/cli/internal/config"
)

type Blocker struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Blocker {
	return &Blocker{cfg: cfg}
}

func (b *Blocker) BlockSites() error {
	return BlockWebsites(b.cfg.BlockedWebsites)
}

func (b *Blocker) UnblockSites() error {
	return UnblockWebsites()
}

// launches 'lockin daemon' as a detached background process
// and writes its PID to ~/.lockin/lockin.pid.
func (b *Blocker) SpawnDaemon() error {
	binary, err := os.Executable()
	if err != nil {
		return fmt.Errorf("resolving executable path: %w", err)
	}

	cmd := exec.Command(binary, "daemon")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.Stdin = nil

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("starting daemon: %w", err)
	}

	pidPath, err := config.PidPath()
	if err != nil {
		return err
	}

	pid := strconv.Itoa(cmd.Process.Pid)
	if err := os.WriteFile(pidPath, []byte(pid), 0644); err != nil {
		return fmt.Errorf("writing pid file: %w", err)
	}

	// detach, don't wait for the child
	_ = cmd.Process.Release()

	return nil
}

// kill from the pid file
func StopDaemon() error {
	pidPath, err := config.PidPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(pidPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		os.Remove(pidPath)
		return nil
	}

	proc, err := os.FindProcess(pid)
	if err == nil {
		_ = proc.Signal(syscall.SIGTERM)
		for i := 0; i < 20; i++ {
			if proc.Signal(syscall.Signal(0)) != nil {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}

	os.Remove(pidPath)
	return nil
}

func IsActive() bool {
	return HostsBlocked() || isDaemonRunning()
}

func isDaemonRunning() bool {
	pidPath, err := config.PidPath()
	if err != nil {
		return false
	}

	data, err := os.ReadFile(pidPath)
	if err != nil {
		return false
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return false
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	return proc.Signal(syscall.Signal(0)) == nil
}

func runCommand(name string, args ...string) error {
	return exec.Command(name, args...).Run()
}
