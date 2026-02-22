package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "omen",
	Short: "A grab-bag cli toolkit for terminal nerds.",
	Long: `omen - dice, coins, fortunes, tarot, moon phases, uuids, and more.

One binary. No network. No dependencies.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = version
}
