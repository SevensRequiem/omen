package cmd

import (
	"fmt"

	"github.com/SevensRequiem/omen/internal/data"
	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "fortune",
		Short: "The stars have something to say.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fortunes := data.Fortunes()
			fmt.Fprintln(cmd.OutOrStdout(), fortunes[rng.Intn(len(fortunes))])
		},
	})
}
