package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func parseEpoch(s string) (time.Time, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timestamp: %s", s)
	}
	if n > 9999999999 {
		return time.UnixMilli(n), nil
	}
	return time.Unix(n, 0), nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "epoch [timestamp]",
		Short: "The time, unmasked.",
		Long:  "No args: current unix timestamp. With arg: convert epoch to human-readable.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				fmt.Fprintln(cmd.OutOrStdout(), time.Now().Unix())
				return nil
			}
			t, err := parseEpoch(args[0])
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), t.Format(time.RFC3339))
			return nil
		},
	})
}
