package cmd

import (
	"fmt"

	"github.com/SevensRequiem/omen/internal/data"
	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

func init() {
	var spread int

	tarotCmd := &cobra.Command{
		Use:   "tarot",
		Short: "Draw from the deck.",
		Long:  "Pull a card from the major and minor arcana. The deck remembers.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if spread < 1 || spread > 10 {
				return fmt.Errorf("spread must be between 1 and 10")
			}
			deck := make([]int, len(data.TarotDeck))
			for i := range deck {
				deck[i] = i
			}
			rng.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

			labels := []string{}
			if spread == 3 {
				labels = []string{"Past", "Present", "Future"}
			}

			for i := 0; i < spread; i++ {
				card := data.TarotDeck[deck[i]]
				reversed := rng.Bool()

				if i > 0 {
					fmt.Fprintln(cmd.OutOrStdout())
				}

				if i < len(labels) {
					fmt.Fprintf(cmd.OutOrStdout(), "── %s ──\n", labels[i])
				}

				orientation := "Upright"
				meaning := card.Upright
				if reversed {
					orientation = "Reversed"
					meaning = card.Reversed
				}

				fmt.Fprintf(cmd.OutOrStdout(), "%s (%s) - %s\n", card.Name, orientation, card.Arcana)
				fmt.Fprintf(cmd.OutOrStdout(), "%s\n", meaning)
			}
			return nil
		},
	}

	tarotCmd.Flags().IntVarP(&spread, "spread", "s", 1, "number of cards to draw (3 = past/present/future)")
	rootCmd.AddCommand(tarotCmd)
}
