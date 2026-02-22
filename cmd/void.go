package cmd

import (
	"fmt"

	"github.com/SevensRequiem/omen/internal/data"
	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "void",
		Short: "Whispers from the other side.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			messages := data.Void()
			fmt.Fprintln(cmd.OutOrStdout(), messages[rng.Intn(len(messages))])
		},
	})
}
