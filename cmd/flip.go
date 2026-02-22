package cmd

import (
	"fmt"

	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	flipCmd := &cobra.Command{
		Use:     "flip",
		Aliases: []string{"coinflip"},
		Short:   "Toss a coin into the dark.",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if rng.Bool() {
				fmt.Fprintln(cmd.OutOrStdout(), "Heads")
			} else {
				fmt.Fprintln(cmd.OutOrStdout(), "Tails")
			}
		},
	}
	rootCmd.AddCommand(flipCmd)
}
