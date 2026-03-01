package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/0xmanny/lockin/cli/internal/blocker"
	"github.com/0xmanny/lockin/cli/internal/config"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

const otherWebsite = "+ Add custom website..."
const otherApp = "+ Add custom app..."

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interactively configure blocked websites and apps",
	Long:  "Open an interactive selector to toggle which websites and apps are blocked",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := runInteractiveConfig(); err != nil {
			return err
		}

		if blocker.IsActive() {
			fmt.Println()
			fmt.Println("  Restarting lockin to apply changes...")

			binary, err := os.Executable()
			if err != nil {
				return fmt.Errorf("failed to find binary: %w", err)
			}

			apply := exec.Command(binary, "_apply")
			apply.Stdin = os.Stdin
			apply.Stdout = os.Stdout
			apply.Stderr = os.Stderr
			if err := apply.Run(); err != nil {
				return fmt.Errorf("failed to apply changes: %w", err)
			}
		}

		return nil
	},
}

// runInteractiveConfig runs the picker and saves the config
// does NOT restart blocking, the caller decides whether to do that
func runInteractiveConfig() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	selectedSites, err := pickWebsites(cfg)
	if err != nil {
		return err
	}

	selectedApps, err := pickApps(cfg)
	if err != nil {
		return err
	}

	cfg.BlockedWebsites = selectedSites
	cfg.BlockedApps = selectedApps

	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	cfgPath, _ := config.FilePath()
	fmt.Println()
	fmt.Printf("  Saved to %s\n", cfgPath)
	fmt.Printf("  %d websites · %d apps\n", len(cfg.BlockedWebsites), len(cfg.BlockedApps))

	return nil
}

func pickWebsites(cfg *config.Config) ([]string, error) {
	currentSet := toSet(cfg.BlockedWebsites)

	allSites := uniqueOrdered(config.KnownWebsites, cfg.BlockedWebsites)

	options := make([]huh.Option[string], 0, len(allSites)+1)
	for _, s := range allSites {
		options = append(options, huh.NewOption(s, s).Selected(currentSet[s]))
	}
	options = append(options, huh.NewOption(otherWebsite, otherWebsite))

	var selected []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Blocked Websites").
				Description("Space to toggle · Enter to confirm").
				Options(options...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	hasOther := false
	result := make([]string, 0, len(selected))
	for _, s := range selected {
		if s == otherWebsite {
			hasOther = true
			continue
		}
		result = append(result, s)
	}

	if hasOther {
		custom, err := promptCustomEntries("Add website", "Domain (e.g. example.com) — leave empty to finish")
		if err != nil {
			return nil, err
		}
		result = append(result, custom...)
	}

	return result, nil
}

func pickApps(cfg *config.Config) ([]string, error) {
	currentSet := toSet(cfg.BlockedApps)

	allApps := uniqueOrdered(config.KnownApps, cfg.BlockedApps)

	options := make([]huh.Option[string], 0, len(allApps)+1)
	for _, a := range allApps {
		options = append(options, huh.NewOption(a, a).Selected(currentSet[a]))
	}
	options = append(options, huh.NewOption(otherApp, otherApp))

	var selected []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Blocked Apps").
				Description("Space to toggle · Enter to confirm").
				Options(options...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	hasOther := false
	result := make([]string, 0, len(selected))
	for _, a := range selected {
		if a == otherApp {
			hasOther = true
			continue
		}
		result = append(result, a)
	}

	if hasOther {
		custom, err := promptCustomEntries("Add app", "App name (e.g. Notion) — leave empty to finish")
		if err != nil {
			return nil, err
		}
		result = append(result, custom...)
	}

	return result, nil
}

func promptCustomEntries(title string, description string) ([]string, error) {
	var entries []string

	for i := 1; ; i++ {
		var value string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title(fmt.Sprintf("%s #%d", title, i)).
					Description(description).
					Value(&value),
			),
		)

		if err := form.Run(); err != nil {
			return nil, err
		}

		value = strings.TrimSpace(value)
		if value == "" {
			break
		}
		entries = append(entries, value)
	}

	return entries, nil
}

func toSet(items []string) map[string]bool {
	s := make(map[string]bool, len(items))
	for _, item := range items {
		s[item] = true
	}
	return s
}

// merges the known list with any custom items from the config,
// preserving the known order and appending custom items at the end
func uniqueOrdered(known []string, current []string) []string {
	seen := make(map[string]bool, len(known))
	result := make([]string, 0, len(known)+len(current))

	for _, k := range known {
		if !seen[k] {
			seen[k] = true
			result = append(result, k)
		}
	}

	for _, c := range current {
		if !seen[c] {
			seen[c] = true
			result = append(result, c)
		}
	}

	return result
}

func init() {
	rootCmd.AddCommand(configCmd)
}
