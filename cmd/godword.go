package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SevensRequiem/omen/internal/data"
	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "godword [N]",
		Short: "Words from beyond.",
		Long:  "Random words from the lexicon. God says what God says.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			count := 1
			if len(args) > 0 {
				n, err := strconv.Atoi(args[0])
				if err != nil || n < 1 {
					return fmt.Errorf("invalid count: %s", args[0])
				}
				count = n
			}
			words := data.Lexicon()
			selected := make([]string, count)
			for i := range count {
				selected[i] = words[rng.Intn(len(words))]
			}
			fmt.Fprintln(cmd.OutOrStdout(), strings.Join(selected, " "))
			return nil
		},
	})
}
