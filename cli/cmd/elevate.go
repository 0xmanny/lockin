package cmd

import (
	"os"
	"os/exec"
)

// elevateIfNeeded re-executes the current binary via sudo if not already root.
// replaces the current process entirely, so this never returns on success.
func elevateIfNeeded() error {
	if os.Geteuid() == 0 {
		return nil
	}

	binary, err := os.Executable()
	if err != nil {
		return err
	}

	sudoArgs := []string{"--preserve-env=LOCKIN_ELEVATED,LOCKIN_FIRST_RUN", binary}
	sudoArgs = append(sudoArgs, os.Args[1:]...)
	cmd := exec.Command("sudo", sudoArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		return err
	}

	os.Exit(0)
	return nil // unreachable, but satisfies the compiler
}
