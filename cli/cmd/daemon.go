package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/0xmanny/lockin/cli/internal/blocker"
	"github.com/0xmanny/lockin/cli/internal/config"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:    "daemon",
	Short:  "Run the app-blocking daemon (internal)",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		ab := blocker.NewAppBlocker(cfg.BlockedApps, cfg.Settings.PollInterval)
		ab.Start()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		ab.Stop()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
