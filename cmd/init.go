package cmd

import (
	"verdict/internal/domain/initialize"

	"verdict/internal/infrastructure/file"

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
