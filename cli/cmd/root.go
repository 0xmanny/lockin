package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const banner = `
  ██╗      ██████╗  ██████╗██╗  ██╗    ██╗███╗   ██╗
  ██║     ██╔═══██╗██╔════╝██║ ██╔╝    ██║████╗  ██║
  ██║     ██║   ██║██║     █████╔╝     ██║██╔██╗ ██║
  ██║     ██║   ██║██║     ██╔═██╗     ██║██║╚██╗██║
  ███████╗╚██████╔╝╚██████╗██║  ██╗    ██║██║ ╚████║
  ╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝    ╚═╝╚═╝  ╚═══╝
`

var rootCmd = &cobra.Command{
	Use:   "lockin",
	Short: "Block distracting websites and apps on macOS",
	Long:  banner + "  Block distracting websites and apps on macOS.\n  https://lockin.sh",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startCmd.RunE(startCmd, args)
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
