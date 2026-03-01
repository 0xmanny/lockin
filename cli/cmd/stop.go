package cmd

import (
	"fmt"

	"github.com/0xmanny/lockin/cli/internal/blocker"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Deactivate all blocking",
	Long:  "Stop blocking websites and applications, restoring /etc/hosts to its original state.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !blocker.IsActive() {
			fmt.Println("🔓 lockin is not active")
			return nil
		}

		if err := elevateIfNeeded(); err != nil {
			return fmt.Errorf("failed to get root privileges: %w", err)
		}

		_ = blocker.StopDaemon()

		if err := blocker.UnblockWebsites(); err != nil {
			return fmt.Errorf("failed to unblock websites: %w", err)
		}

		fmt.Println("🔓 All blocks removed. Stay focused!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
