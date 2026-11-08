package cmd

import (
	"github.com/Shioalati/Verdict/internal/infrastructure/file"

	"github.com/Shioalati/Verdict/internal/domain/initialize"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {
		s := file.NewLocal()

		err := initialize.Handle(s)
		return err
	},
}
