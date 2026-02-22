package cmd

import (
	"fmt"

	"github.com/SevensRequiem/omen/internal/data"
	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "8ball [question]",
		Short: "Shake the oracle.",
		Long:  "Ask a question. Or don't. The answer doesn't care.",
		Run: func(cmd *cobra.Command, args []string) {
			responses := data.Eightball()
			fmt.Fprintln(cmd.OutOrStdout(), responses[rng.Intn(len(responses))])
		},
	})
}
