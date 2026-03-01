package cmd

import (
	"fmt"

	"github.com/0xmanny/lockin/cli/internal/blocker"
	"github.com/0xmanny/lockin/cli/internal/config"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current blocking status",
	Long:  "Display which websites and applications are currently being blocked.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		active := blocker.IsActive()

		if active {
			fmt.Println("🔒 lockin is ACTIVE")
		} else {
			fmt.Println("🔓 lockin is INACTIVE")
		}

		fmt.Println()

		configPath, _ := config.FilePath()
		fmt.Printf("Config: %s\n\n", configPath)

		fmt.Println("Blocked websites:")
		if len(cfg.BlockedWebsites) == 0 {
			fmt.Println("  (none)")
		} else {
			for _, site := range cfg.BlockedWebsites {
				fmt.Printf("  • %s\n", site)
			}
		}

		fmt.Println()
		fmt.Println("Blocked apps:")
		if len(cfg.BlockedApps) == 0 {
			fmt.Println("  (none)")
		} else {
			for _, app := range cfg.BlockedApps {
				fmt.Printf("  • %s\n", app)
			}
		}

		if active {
			fmt.Println()
			fmt.Println("Run 'lockin stop' to deactivate.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
