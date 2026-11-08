package cmd

import (
	"fmt"
	"os"

	"github.com/Shioalati/Verdict/internal/infrastructure/file"

	"github.com/Shioalati/Verdict/internal/domain/root"
	"github.com/Shioalati/Verdict/internal/infrastructure/hash"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringVarP(&root.Algo, "algo", "a", "sha256", "algorithm")
	rootCmd.Flags().BoolVarP(&root.List, "list", "l", false, "list per-file hashes for directories")
}

var rootCmd = &cobra.Command{
	Use:  "vec [path]",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			root.Path = args[0]
		}

		l := file.NewLocal()

		h, err := hash.New(root.Algo)
		if err != nil {
			return err
		}

		r, err := root.Handle(l, h)
		if err != nil {
			return err
		}

		fmt.Println(string(r))

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
