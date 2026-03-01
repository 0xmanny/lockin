package cmd

import (
	"fmt"
	"os"

	"github.com/0xmanny/lockin/cli/internal/blocker"
	"github.com/0xmanny/lockin/cli/internal/config"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove lockin and all its data",
	Long:  "Fully remove lockin: restore /etc/hosts, stop the daemon, remove the LaunchDaemon, config, and binary.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := elevateIfNeeded(); err != nil {
			return fmt.Errorf("failed to get root privileges: %w", err)
		}

		_ = blocker.StopDaemon()

		cfg, err := config.Load()
		if err == nil {
			b := blocker.New(cfg)
			_ = b.UnblockSites()
		}

		_ = blocker.UninstallLaunchDaemon()

		dir, err := config.Dir()
		if err == nil {
			_ = os.RemoveAll(dir)
		}

		_ = os.Remove("/usr/local/bin/lockin")

		fmt.Println("lockin has been fully uninstalled.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
