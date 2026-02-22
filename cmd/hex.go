package cmd

import (
	"fmt"

	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "hex",
		Short: "A color from the spectrum.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			r := rng.Intn(256)
			g := rng.Intn(256)
			b := rng.Intn(256)
			hex := fmt.Sprintf("#%02X%02X%02X", r, g, b)
			swatch := fmt.Sprintf("\033[48;2;%d;%d;%dm    \033[0m", r, g, b)
			fmt.Fprintf(cmd.OutOrStdout(), "%s %s\n", hex, swatch)
		},
	})
}
