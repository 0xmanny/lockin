package cmd

import (
	"fmt"

	"github.com/0xmanny/lockin/cli/internal/blocker"
	"github.com/0xmanny/lockin/cli/internal/config"
	"github.com/spf13/cobra"
)

// a hidden command that restarts blocking with the current config
// used internally by `lockin config` to apply changes without two sudo prompts
var applyCmd = &cobra.Command{
	Use:    "_apply",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := elevateIfNeeded(); err != nil {
			return fmt.Errorf("failed to get root privileges: %w", err)
		}

		_ = blocker.StopDaemon()
		_ = blocker.UnblockWebsites()

		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		b := blocker.New(cfg)

		if err := b.BlockSites(); err != nil {
			return fmt.Errorf("failed to block websites: %w", err)
		}

		if err := b.SpawnDaemon(); err != nil {
			_ = b.UnblockSites()
			return fmt.Errorf("failed to start app blocker: %w", err)
		}

		_ = blocker.InstallLaunchDaemon()

		fmt.Println("  🔒 changes applied")
		fmt.Printf("     %d websites · %d apps\n", len(cfg.BlockedWebsites), len(cfg.BlockedApps))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
