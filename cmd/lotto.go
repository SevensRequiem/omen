package cmd

import (
	"fmt"
	"sort"

	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	var count int

	lottoCmd := &cobra.Command{
		Use:   "lotto",
		Short: "Numbers from the ether.",
		Long:  "Generate lottery numbers. Because hope is a resource.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			max := 49
			if count > max {
				count = max
			}
			pool := make([]int, max)
			for i := range pool {
				pool[i] = i + 1
			}
			rng.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })
			picks := pool[:count]
			sort.Ints(picks)

			for i, n := range picks {
				if i > 0 {
					fmt.Fprint(cmd.OutOrStdout(), "  ")
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%02d", n)
			}
			fmt.Fprintln(cmd.OutOrStdout())
		},
	}

	lottoCmd.Flags().IntVarP(&count, "count", "n", 6, "how many numbers to draw")
	rootCmd.AddCommand(lottoCmd)
}
