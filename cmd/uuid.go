package cmd

import (
	"crypto/rand"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var count int

	uuidCmd := &cobra.Command{
		Use:   "uuid",
		Short: "Summon a unique identifier.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if count < 1 {
				return fmt.Errorf("count must be at least 1")
			}
			for range count {
				var uuid [16]byte
				_, err := rand.Read(uuid[:])
				if err != nil {
					return fmt.Errorf("failed to generate UUID: %w", err)
				}
				uuid[6] = (uuid[6] & 0x0f) | 0x40
				uuid[8] = (uuid[8] & 0x3f) | 0x80
				fmt.Fprintf(cmd.OutOrStdout(), "%08x-%04x-%04x-%04x-%012x\n",
					uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
			}
			return nil
		},
	}

	uuidCmd.Flags().IntVarP(&count, "count", "n", 1, "number of UUIDs to generate")
	rootCmd.AddCommand(uuidCmd)
}
