package cmd

import (
	"verdict/internal/domain/snapshot"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(snapshotCmd)
}

var snapshotCmd = &cobra.Command{
	Use: "snapshot",
	RunE: func(cmd *cobra.Command, args []string) error {
		return snapshot.Handle()
	},
}
