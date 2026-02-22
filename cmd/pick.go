package cmd

import (
	"fmt"

	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "pick <choice> <choice> [choice...]",
		Short: "Let fate decide.",
		Long:  "Give it options. It picks one. No take-backs.",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), args[rng.Intn(len(args))])
		},
	})
}
