package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/SevensRequiem/omen/internal/rng"
	"github.com/spf13/cobra"
)

var diceRegex = regexp.MustCompile(`^(\d*)d(\d+)([+-]\d+)?$`)

func parseDice(input string) (n, sides, mod int, err error) {
	if input == "" {
		return 1, 20, 0, nil
	}
	input = strings.ToLower(strings.TrimSpace(input))
	matches := diceRegex.FindStringSubmatch(input)
	if matches == nil {
		return 0, 0, 0, fmt.Errorf("invalid dice notation: %s (use NdN format, e.g. 2d6, d20, 4d8+3)", input)
	}
	n = 1
	if matches[1] != "" {
		n, err = strconv.Atoi(matches[1])
		if err != nil || n < 1 {
			return 0, 0, 0, fmt.Errorf("invalid number of dice: %s", matches[1])
		}
	}
	sides, err = strconv.Atoi(matches[2])
	if err != nil || sides < 1 {
		return 0, 0, 0, fmt.Errorf("invalid number of sides: %s", matches[2])
	}
	if matches[3] != "" {
		mod, err = strconv.Atoi(matches[3])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid modifier: %s", matches[3])
		}
	}
	return n, sides, mod, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "roll [NdN]",
		Short: "Cast the bones.",
		Long:  "Roll dice in NdN notation. Defaults to 1d20. Supports modifiers like 2d6+3.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			input := ""
			if len(args) > 0 {
				input = args[0]
			}
			n, sides, mod, err := parseDice(input)
			if err != nil {
				return err
			}
			rolls := make([]int, n)
			total := 0
			for i := 0; i < n; i++ {
				rolls[i] = rng.Intn(sides) + 1
				total += rolls[i]
			}
			total += mod

			if n == 1 && mod == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "%d\n", rolls[0])
			} else if n > 1 && mod == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "%v → %d\n", rolls, total)
			} else {
				sign := "+"
				absMod := mod
				if mod < 0 {
					sign = "-"
					absMod = -mod
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%v %s%d → %d\n", rolls, sign, absMod, total)
			}
			return nil
		},
	})
}
