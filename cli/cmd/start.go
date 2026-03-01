package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/0xmanny/lockin/cli/internal/blocker"
	"github.com/0xmanny/lockin/cli/internal/config"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Activate website and app blocking",
	Long:  "Start blocking distracting websites and applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		firstRun := !config.Exists()
		elevated := os.Getenv("LOCKIN_ELEVATED") == "1"

		if firstRun {
			fmt.Print(banner)
			fmt.Println("  Welcome anon! Let's configure what to block.")
			fmt.Println()

			if err := runInteractiveConfig(); err != nil {
				return err
			}
		}

		if !firstRun && !elevated && blocker.IsActive() {
			fmt.Println("🔒 lockin is already active")
			return nil
		}

		if os.Geteuid() != 0 {
			if !sudoCached() {
				fmt.Println()
				fmt.Println("  lockin needs to execute as root to block websites and apps")
				fmt.Println()
			}

			os.Setenv("LOCKIN_ELEVATED", "1")
			if firstRun {
				os.Setenv("LOCKIN_FIRST_RUN", "1")
			}
			if err := elevateIfNeeded(); err != nil {
				return fmt.Errorf("failed to get root privileges: %w", err)
			}
		}

		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		b := blocker.New(cfg)

		_ = blocker.StopDaemon()

		if err := b.BlockSites(); err != nil {
			return fmt.Errorf("failed to block websites: %w", err)
		}

		if err := b.SpawnDaemon(); err != nil {
			_ = b.UnblockSites()
			return fmt.Errorf("failed to start app blocker: %w", err)
		}

		_ = blocker.InstallLaunchDaemon()

		if !elevated {
			fmt.Print(banner)
		}
		fmt.Println("  🔒 distractions blocked")
		fmt.Printf("     %d websites · %d apps\n", len(cfg.BlockedWebsites), len(cfg.BlockedApps))
		fmt.Println()

		if firstRun || os.Getenv("LOCKIN_FIRST_RUN") == "1" {
			fmt.Println("  🛑 `lockin stop` to unblock")
			fmt.Println("  🔍 `lockin status` to see what's blocked")
			fmt.Println("  🔧 `lockin config` to change the blocklist")
			fmt.Println("  🗑️  `lockin uninstall` to remove lockin") // i have no idea why an extra space is needed here
			fmt.Println()
			fmt.Println("  For more info, visit https://lockin.sh")
			fmt.Println()
		}

		return nil
	},
}

// true if sudo credentials are already cached (no prompt needed)
func sudoCached() bool {
	return exec.Command("sudo", "-n", "true").Run() == nil
}

func init() {
	rootCmd.AddCommand(startCmd)
}
